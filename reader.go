package reader

import (
	"context"
	"io"

	"github.com/whosonfirst/go-ioutil"
	wof_reader "github.com/whosonfirst/go-reader/v2"
	"gocloud.dev/blob"
)

func init() {

	ctx := context.Background()

	for _, scheme := range blob.DefaultURLMux().BucketSchemes() {

		err := wof_reader.RegisterReader(ctx, scheme, NewBlobReader)

		if err != nil {
			panic(err)
		}
	}
}

type BlobReader struct {
	wof_reader.Reader
	bucket *blob.Bucket
}

func NewBlobReader(ctx context.Context, uri string) (wof_reader.Reader, error) {

	bucket, err := blob.OpenBucket(ctx, uri)

	if err != nil {
		return nil, err
	}

	r := &BlobReader{
		bucket: bucket,
	}

	return r, nil
}

func (r *BlobReader) Read(ctx context.Context, uri string) (io.ReadSeekCloser, error) {

	blob_r, err := r.bucket.NewReader(ctx, uri, nil)

	if err != nil {
		return nil, err
	}

	return ioutil.NewReadSeekCloser(blob_r)
}

func (r *BlobReader) Exists(ctx context.Context, uri string) (bool, error) {
	return r.bucket.Exists(ctx, uri)
}

func (r *BlobReader) ReaderURI(ctx context.Context, uri string) string {
	return uri
}
