package lsmt

import (
	"bytes"
	"io"
)

const (
	maxFileLen       = 1024
	indexSparseRatio = 3
)

type File struct {
	Index  *Node
	Data   io.ReadSeeker
	Size   int
	Buffer bytes.Buffer
}
