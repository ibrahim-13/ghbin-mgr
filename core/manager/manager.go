package manager

import (
	"net/http"
	"os"
	"path/filepath"
)

type ArchiveType string

const (
	ArchiveGzip ArchiveType = "gzip"
	ArchiveZip  ArchiveType = "zip"
)

func DownloadAndExtract(downloadUrl, location string, aType ArchiveType, pattern ...string) error {
	parent := filepath.Dir(location)
	basename := filepath.Base(location)
	archivefile := filepath.Join(parent, basename)
	err := downloadFile(downloadUrl, archivefile, http.DefaultClient)
	if err != nil {
		return err
	}

	switch aType {
	case ArchiveGzip:
		err = extractFileTarGz(archivefile, location, pattern...)
	case ArchiveZip:
		err = extractFileZip(archivefile, location, pattern...)
	}
	if err != nil {
		return err
	}
	err = os.Chmod(location, 0755)
	if err != nil {
		return err
	}
	err = os.Remove(archivefile)
	if err != nil {
		return err
	}
	return nil
}

func Download(downloadUrl, location string) error {
	err := downloadFile(downloadUrl, location, http.DefaultClient)
	if err != nil {
		return err
	}
	err = os.Chmod(location, 0755)
	if err != nil {
		return err
	}
	return nil
}
