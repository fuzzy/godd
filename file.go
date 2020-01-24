package main

import "os"

type GoddFileMover struct {
	name   string
	handle *os.File
}

func (self GoddFileMover) Sync() {
	_ = self.handle.Sync()
}

func (self GoddFileMover) Size() int64 {
	r, _ := self.handle.Stat()
	return r.Size()
}

func (self GoddFileMover) Read(b []byte) (n int, err error) {
	return self.handle.Read(b)
}

func (self GoddFileMover) Write(b []byte) (n int, err error) {
	return self.handle.Write(b)
}

func NewFileReader(fn string) GoddFileMover {
	fp, _ := os.Open(fn)
	return GoddFileMover{fn, fp}
}

func NewFileWriter(fn string) GoddFileMover {
	fp, _ := os.OpenFile(fn, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600)
	return GoddFileMover{fn, fp}
}
