//go:build s3blob

package reader

import (
	_ "github.com/aaronland/gocloud-blob/s3"
	_ "gocloud.dev/blob/s3blob"
)
