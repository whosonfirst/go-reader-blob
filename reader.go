package reader

import (
	"context"
	wof_reader "github.com/whosonfirst/go-reader"
	"gocloud.dev/blob"
	"io"
	"net/url"
)

func init() {
	r := NewBlobReader()
	wof_reader.Register("blob", r)
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

	u, err := url.Parse(uri)

	if err != nil {
		return err
	}

	u.Scheme = u.Host
	u.Host = ""

	blob_uri := u.String()

	bucket, err := blob.OpenBucket(ctx, blob_uri)

	if err != nil {
		return err
	}

	r.bucket = bucket
	return nil
}

func (r *BlobReader) Read(ctx context.Context, uri string) (io.ReadCloser, error) {

	return r.bucket.NewReader(ctx, uri, nil)
}
