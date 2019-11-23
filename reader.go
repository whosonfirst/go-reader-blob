package reader

import (
	"context"
	wof_reader "github.com/whosonfirst/go-reader"
	"gocloud.dev/blob"
	"io"
)

func init() {

	r := NewBlobReader()

	for _, scheme := range blob.DefaultURLMux().BucketSchemes() {
		wof_reader.Register(scheme, r)
	}
}

type BlobReader struct {
	wof_reader.Reader
	bucket *blob.Bucket
}

func NewBlobReader() wof_reader.Reader {

	r := BlobReader{}
	return &r
}

func (r *BlobReader) Open(ctx context.Context, uri string) error {

	bucket, err := blob.OpenBucket(ctx, uri)

	if err != nil {
		return err
	}

	r.bucket = bucket
	return nil
}

func (r *BlobReader) Read(ctx context.Context, uri string) (io.ReadCloser, error) {

	return r.bucket.NewReader(ctx, uri, nil)
}
