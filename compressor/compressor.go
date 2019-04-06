package compressor

import (
	"bytes"
	"compress/gzip"
	"io"
	"time"
)

type Compressor struct {
	name string
	buf  []uint8
}

func New(name string, size int) *Compressor {
	return &Compressor{
		buf: make([]uint8, size),
	}
}

func (c *Compressor) Buffer() []uint8 {
	return c.buf
}

func (c *Compressor) Compress() (io.Reader, error) {
	src := bytes.NewBuffer(c.buf)
	var buf bytes.Buffer

	zw := gzip.NewWriter(&buf)
	zw.Name = c.name
	zw.ModTime = time.Now()

	if _, err := io.Copy(zw, src); err != nil {
		return nil, err
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return &buf, nil
}
