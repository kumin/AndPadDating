package entities

import "io"

type File struct {
	Name        string
	ContentType string
	Size        int64
	Buffer      io.Reader
}
