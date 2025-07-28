package internal

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func CreateArchive(paths []string, archivePath string) error {
	out, err := os.Create(archivePath)
	if err != nil {
		return err
	}
	defer out.Close()
	
	gw := gzip.NewWriter(out)
	defer gw.Close()
	
	tw := tar.NewWriter(gw)
	defer tw.Close()
	
	for _, path := range paths {
		err := filepath.Walk(path, func(file string, fi os.FileInfo, err error) error {
			if err != nil || fi.IsDir() {
				return err
			}
			hdr, err := tar.FileInfoHeader(fi, file)
			if err != nil {
				return err
			}
			rel, _ := filepath.Rel(path, file)
			hdr.Name = filepath.Join(filepath.Base(path), rel)
			
			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}
			
			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(tw, f)
			return err
		})
		if err != nil {
			return err
		}
	}
	return nil
}
