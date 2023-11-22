package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/staaldraad/fusefs-hello-world/hfs"
)

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {

	var mountpoint string

	flag.StringVar(&mountpoint, "mount", "", "--mount /path/to/mount")
	flag.Usage = usage
	flag.Parse()
	if mountpoint == "" {
		usage()
	}

	fmt.Println(mountpoint)

	c, err := fuse.Mount(
		mountpoint,
		fuse.FSName("hfs"),
		fuse.Subtype("hfs"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	//f := hfs.NewFS()
	srv := fs.New(c, nil)
	filesys := &hfs.FS{
		// We pre-create the clock node so that it's always the same
		// object returned from all the Lookups. You could carefully
		// track its lifetime between Lookup&Forget, and have the
		// ticking & invalidation happen only when active, but let's
		// keep this example simple.
		DynamicFile: &hfs.File{
			Fuse: srv,
		},
	}

	err = fs.Serve(c, filesys)
	if err != nil {
		log.Fatal(err)
	}
}
