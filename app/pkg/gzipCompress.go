package pkg

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"time"
)

func GzipBytes(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	zw.Name = "compresses body"
	zw.Comment = "this body is compressesd using gzip"
	zw.ModTime = time.Date(1997, time.May, 25, 0, 0, 0, 0, time.UTC)

	_, err := zw.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed to compress data %w", err)
	}

	if err = zw.Close(); err != nil {
		return nil, fmt.Errorf("failed to close the channel %w", err)
	}
	return buf.Bytes(), nil

}
