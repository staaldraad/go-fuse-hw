package hfs

import (
	"github.com/anacrolix/fuse"
	"github.com/anacrolix/fuse/fs"
)

var inodeCount uint64

type EntryGetter interface {
	GetDirentType() fuse.DirentType
}

type FS struct{}

func NewFS() FS {
	return FS{}
}

func (f FS) Root() (fs.Node, error) {
	return NewDir(), nil
}
