package main

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// Only interfaces can be embedded within interfaces
type ReadWriter interface {
	Reader
	Writer
}

