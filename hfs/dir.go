package hfs

import (
	"context"
	"log"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type Dir struct {
	fs *FS
}

var _ fs.Node = (*Dir)(nil)

func (d *Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 1
	a.Mode = os.ModeDir | 0o755
	return nil
}

var _ fs.NodeStringLookuper = (*Dir)(nil)

func (d *Dir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	log.Println("Lookup: ", name)
	d.fs.DynamicFile.Name.Store(name)
	return d.fs.DynamicFile, nil
	// }
	// return nil, syscall.ENOENT
}

var dirDirs = []fuse.Dirent{
	{Inode: 2, Name: "clock", Type: fuse.DT_File},
}

var _ fs.HandleReadDirAller = (*Dir)(nil)

func (d *Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	return dirDirs, nil
}

// type Dir struct {
// 	Type       fuse.DirentType
// 	Attributes fuse.Attr
// 	Entries    map[string]interface{}
// }

// var _ = (fs.Node)((*Dir)(nil))
// var _ = (fs.NodeMkdirer)((*Dir)(nil))
// var _ = (fs.NodeCreater)((*Dir)(nil))
// var _ = (fs.HandleReadDirAller)((*Dir)(nil))
// var _ = (fs.NodeSetattrer)((*Dir)(nil))
// var _ = (EntryGetter)((*Dir)(nil))

// func NewDir() *Dir {
// 	log.Println("NewDir called")
// 	atomic.AddUint64(&inodeCount, 1)
// 	return &Dir{
// 		Type: fuse.DT_Dir,
// 		Attributes: fuse.Attr{
// 			Inode: inodeCount,
// 			Atime: time.Now(),
// 			Mtime: time.Now(),
// 			Ctime: time.Now(),
// 			Mode:  os.ModeDir | 0o777,
// 		},
// 		Entries: map[string]interface{}{},
// 	}
// }

// func (d *Dir) GetDirentType() fuse.DirentType {
// 	return d.Type
// }

// func (d *Dir) Mkdir(ctx context.Context, req *fuse.MkdirRequest) (fs.Node, error) {
// 	log.Println("Mkdir called with name: ", req.Name)
// 	return nil, fmt.Errorf("permission denied")
// }

// func (d *Dir) Attr(ctx context.Context, a *fuse.Attr) error {
// 	*a = d.Attributes
// 	log.Println("Attr permissions: ", a.Mode)
// 	log.Println("Attr: Modified At", a.Mtime.String())
// 	return nil
// }

// func (d *Dir) Lookup(ctx context.Context, name string) (fs.Node, error) {
// 	log.Println("LookUP called: ", name)
// 	node, ok := d.Entries
// 	return FileN{Name: name}, nil

// }

// func (d *Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
// 	log.Println("ReadDirAll called")
// 	// var entries []fuse.Dirent

// 	// for k, v := range d.Entries {
// 	// 	var a fuse.Attr
// 	// 	v.(fs.Node).Attr(ctx, &a)
// 	// 	entries = append(entries, fuse.Dirent{
// 	// 		Inode: a.Inode,
// 	// 		Type:  v.(EntryGetter).GetDirentType(),
// 	// 		Name:  k,
// 	// 	})
// 	// }
// 	return nil, nil
// }

// func (d *Dir) Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (fs.Node, fs.Handle, error) {
// 	// log.Println("Create called with filename: ", req.Name)
// 	// f := NewFile(nil)
// 	// log.Println("Create: Modified at", f.Attributes.Mtime.String())
// 	// d.Entries[req.Name] = f
// 	log.Printf("Denied create of %s\n", req.Name)
// 	return nil, nil, syscall.ENOENT
// }

// func (d *Dir) Setattr(ctx context.Context, req *fuse.SetattrRequest, resp *fuse.SetattrResponse) error {
// 	if req.Valid.Atime() {
// 		d.Attributes.Atime = req.Atime
// 	}
// 	if req.Valid.Mtime() {
// 		d.Attributes.Mtime = req.Mtime
// 	}
// 	if req.Valid.Size() {
// 		d.Attributes.Size = req.Size
// 	}
// 	log.Println("Setattr called: Attributes ", d.Attributes)
// 	return nil
// }

// func (d *Dir) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
// 	//delete(d.Entries, req.Name)
// 	log.Printf("Denied remove of %s\n", req.Name)
// 	return syscall.ENOENT
// }
