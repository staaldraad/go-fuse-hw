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

	err = fs.Serve(c, hfs.NewFS())
	if err != nil {
		log.Fatal(err)
	}
}
