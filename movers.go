package main

type GoddMover interface {
	Size() int64
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Sync()
}

type GoddReader interface {
	Read(b []byte) (n int, err error)
}

type GoddWriter interface {
	Write(b []byte) (n int, err error)
	Sync()
}
