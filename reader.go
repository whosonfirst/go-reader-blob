package reader

import (
	"context"
	wof_reader "github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-reader/ioutil"	
	"gocloud.dev/blob"
	"io"
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

	fh, err := r.bucket.NewReader(ctx, uri, nil)

	if err != nil {
		return nil, err
	}

	return ioutil.NewReadSeekCloser(fh)
}

func (r *BlobReader) ReaderURI(ctx context.Context, uri string) string {
	return uri
}
