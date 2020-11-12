package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

func main() {
	fw, _ := os.Create("tmp.tar.gz")
	defer fw.Close()

	gw := gzip.NewWriter(fw)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	fi, _ := os.Open("tmp")
	// fmt.Println(fi.Name())
	defer fi.Close()
	info, _ := fi.Stat()

	header, _ := tar.FileInfoHeader(info, info.Name())
	tw.WriteHeader(header)
	io.Copy(tw, fi)
}
