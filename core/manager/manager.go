package manager

import (
	"errors"
	"gbm/util"
	"net/http"
	"os"
	"path/filepath"
)

type ArchiveType string

const (
	ArchiveGzip ArchiveType = "gzip"
	ArchiveZip  ArchiveType = "zip"
)

func DownloadAndExtract(assetName, downloadUrl, location string, pattern ...string) error {
	var aType ArchiveType
	if util.ContainsAnyMatches(assetName, ".tar.gz") {
		aType = ArchiveGzip
	} else if util.ContainsAnyMatches(assetName, ".zip") {
		aType = ArchiveZip
	} else {
		return errors.New("could not detect archive type")
	}
	parent := filepath.Dir(location)
	archivefile := filepath.Join(parent, assetName)
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
