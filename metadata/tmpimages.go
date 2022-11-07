package metadata

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/metadata/boltutil"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/tmpimages"
	digest "github.com/opencontainers/go-digest"
	bolt "go.etcd.io/bbolt"
)

type tmpImageStore struct {
	db *DB
}

func NewTmpImageStore(db *DB) tmpimages.Store {
	return &tmpImageStore{db: db}
}

func (s *tmpImageStore) InsertTmpImage(ctx context.Context, tmpimage tmpimages.TmpImage) (tmpimages.TmpImage, error) {
	namespace, err := namespaces.NamespaceRequired(ctx)
	if err != nil {
		return tmpimages.TmpImage{}, err
	}
	if err := update(ctx, s.db, func(tx *bolt.Tx) error {
		if err := validateTmpImage(&tmpimage); err != nil {
			return err
		}
		bkt, err := createTmpImagesBucket(tx, namespace)
		if err != nil {
			return err
		}
		ibkt, err := bkt.CreateBucket([]byte(tmpimage.Name))
		if err != nil {
			if err != bolt.ErrBucketExists {
				return err
			}
			return fmt.Errorf("tmpimage %q: %w", tmpimage.Name, errdefs.ErrAlreadyExists)
		}
		return writeTmpImage(ibkt, &tmpimage)
	}); err != nil {
		return tmpimages.TmpImage{}, err
	}
	return tmpimage, nil
}

func (s *tmpImageStore) GetTmpImage(ctx context.Context, name string) (tmpimages.TmpImage, error) {
	var tmpimage tmpimages.TmpImage
	namespace, err := namespaces.NamespaceRequired(ctx)
	if err != nil {
		return tmpimages.TmpImage{}, err
	}
	if err := view(ctx, s.db, func(tx *bolt.Tx) error {
		bkt := getTmpImagesBucket(tx, namespace)
		if bkt == nil {
			return fmt.Errorf("tmpimage %q: %w", name, errdefs.ErrNotFound)
		}
		ibkt := bkt.Bucket([]byte(name))
		if ibkt == nil {
			return fmt.Errorf("tmpimage %q: %w", name, errdefs.ErrNotFound)
		}
		tmpimage.Name = name
		if err := readTmpImage(&tmpimage, ibkt); err != nil {
			return fmt.Errorf("tmpimage %q: %w", name, err)
		}
		return nil
	}); err != nil {
		return tmpimages.TmpImage{}, err
	}
	return tmpimage, nil
}

func (s *tmpImageStore) Delete(ctx context.Context, name string) error {
	namespace, err := namespaces.NamespaceRequired(ctx)
	if err != nil {
		return err
	}
	return update(ctx, s.db, func(tx *bolt.Tx) error {
		bkt := getTmpImagesBucket(tx, namespace)
		if bkt == nil {
			return fmt.Errorf("tmpimage %q: %w", name, errdefs.ErrNotFound)
		}
		if err = bkt.DeleteBucket([]byte(name)); err != nil {
			if err == bolt.ErrBucketNotFound {
				err = fmt.Errorf("image %q: %w", name, errdefs.ErrNotFound)
			}
			return err
		}
		atomic.AddUint32(&s.db.dirty, 1)
		return nil
	})
}

func validateTmpImage(tmpimage *tmpimages.TmpImage) error {
	if tmpimage.Name == "" {
		return fmt.Errorf("tmpimage name must not be empty: %w", errdefs.ErrInvalidArgument)
	}
	return validateTarget(&tmpimage.Target)
}

func readTmpImage(tmpimage *tmpimages.TmpImage, bkt *bolt.Bucket) error {
	var err error
	tmpimage.Target.Annotations, err = boltutil.ReadAnnotations(bkt)
	if err != nil {
		return err
	}

	tbkt := bkt.Bucket(bucketKeyTarget)
	if tbkt == nil {
		return errors.New("unable to read target bucket")
	}
	return tbkt.ForEach(func(k, v []byte) error {
		if v == nil {
			return nil // skip it? a bkt maybe?
		}

		// TODO(stevvooe): This is why we need to use byte values for
		// keys, rather than full arrays.
		switch string(k) {
		case string(bucketKeyDigest):
			tmpimage.Target.Digest = digest.Digest(v)
		case string(bucketKeyMediaType):
			tmpimage.Target.MediaType = string(v)
		case string(bucketKeySize):
			tmpimage.Target.Size, _ = binary.Varint(v)
		}

		return nil
	})
}

func writeTmpImage(bkt *bolt.Bucket, tmpimage *tmpimages.TmpImage) error {

	if err := boltutil.WriteAnnotations(bkt, tmpimage.Target.Annotations); err != nil {
		return fmt.Errorf("writing Annotations for image %v: %w", tmpimage.Name, err)
	}

	// write the target bucket
	tbkt, err := bkt.CreateBucketIfNotExists(bucketKeyTarget)
	if err != nil {
		return err
	}

	sizeEncoded, err := encodeInt(tmpimage.Target.Size)
	if err != nil {
		return err
	}

	for _, v := range [][2][]byte{
		{bucketKeyDigest, []byte(tmpimage.Target.Digest)},
		{bucketKeyMediaType, []byte(tmpimage.Target.MediaType)},
		{bucketKeySize, sizeEncoded},
	} {
		if err := tbkt.Put(v[0], v[1]); err != nil {
			return err
		}
	}

	return nil
}
