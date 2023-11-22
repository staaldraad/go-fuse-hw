package hfs

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"bazil.org/fuse/fuseutil"
)

type File struct {
	Fuse    *fs.Server
	Reader  atomic.Pointer[io.ReadCloser]
	Name    atomic.Value
	content atomic.Value
	count   uint64
}

var _ fs.Node = (*File)(nil)

func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {
	log.Println("Attr")
	a.Inode = 2
	a.Mode = 0o444
	reader, size := fetchFile(f.Name.Load().(string))
	f.Reader.Store(reader)
	a.Size = uint64(size)
	return nil
}

var _ fs.NodeOpener = (*File)(nil)

func (f *File) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fs.Handle, error) {
	log.Println("Open")
	if !req.Flags.IsReadOnly() {
		return nil, fuse.Errno(syscall.EACCES)
	}
	resp.Flags |= fuse.OpenKeepCache
	return f, nil
}

var _ fs.Handle = (*File)(nil)

var _ fs.HandleReader = (*File)(nil)

func (f *File) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {
	log.Println("Read")
	var t []byte
	reader := f.Reader.Load()
	(*reader).Read(t)
	log.Println(t)
	fuseutil.HandleRead(req, resp, t)
	return nil
}

// type File struct {
// 	Type       fuse.DirentType
// 	Content    []byte
// 	Attributes fuse.Attr
// }

// var _ = (fs.Node)((*File)(nil))
// var _ = (fs.HandleWriter)((*File)(nil))
// var _ = (fs.HandleReadAller)((*File)(nil))
// var _ = (fs.NodeSetattrer)((*File)(nil))
// var _ = (EntryGetter)((*File)(nil))

// func NewFile(content []byte) *File {
// 	log.Println("NewFile called")
// 	atomic.AddUint64(&inodeCount, 1)
// 	return &File{
// 		Type:    fuse.DT_Block,
// 		Content: content,
// 		Attributes: fuse.Attr{
// 			Inode: inodeCount,
// 			Atime: time.Now(),
// 			Mtime: time.Now(),
// 			Ctime: time.Now(),
// 			Mode:  0o777,
// 		},
// 	}
// }

// func (f *File) GetDirentType() fuse.DirentType {
// 	return f.Type
// }

// func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {
// 	*a = f.Attributes
// 	log.Println("Attr: Modified At", a.Mtime.String())
// 	return nil
// }

// func (f *File) ReadAll(ctx context.Context) ([]byte, error) {
// 	log.Println("ReadAll called")
// 	resp, err := http.Get("https://ifconfig.me")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	return body, nil
// }

// func (f *File) Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) error {
// 	log.Println("Write called: Size ", f.Attributes.Size)
// 	log.Println("Data to write: ", string(req.Data))
// 	f.Content = req.Data
// 	resp.Size = len(req.Data)
// 	f.Attributes.Size = uint64(resp.Size)
// 	return nil
// }

// func (f *File) Setattr(ctx context.Context, req *fuse.SetattrRequest, resp *fuse.SetattrResponse) error {
// 	if req.Valid.Atime() {
// 		f.Attributes.Atime = req.Atime
// 	}
// 	if req.Valid.Mtime() {
// 		f.Attributes.Mtime = req.Mtime
// 	}
// 	if req.Valid.Size() {
// 		f.Attributes.Size = req.Size
// 	}
// 	log.Println("Setattr called: Attributes ", f.Attributes)
// 	return nil
// }

// type FileN struct {
// 	Name string
// 	Size uint64
// }

// func (n FileN) Attr(ctx context.Context, a *fuse.Attr) error {
// 	log.Println("Attr called")
// 	var b []byte
// 	if n.Size == 0 {
// 		b = fetchFile(n.Name)
// 		n.Size = uint64(len(b))
// 	}
// 	a.Inode = 2
// 	a.Mode = 0o444
// 	a.Size = n.Size
// 	return nil
// }

// func (n FileN) ReadAll(ctx context.Context) ([]byte, error) {
// 	log.Println("ReadAll: ", n.Name)
// 	body := fetchFile(n.Name)
// 	return body, nil
// }

func fetchFile(name string) (*io.ReadCloser, int64) {
	log.Println("Fetch remote: ", name)
	resp, err := http.Get(fmt.Sprintf("https://%s", name))
	if err != nil {
		log.Fatalln(err)
	}
	//defer resp.Body.Close()
	return &resp.Body, resp.ContentLength
}
