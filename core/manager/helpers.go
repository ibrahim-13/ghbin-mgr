package manager

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

func containsAllMatches(str string, match ...string) bool {
	s := strings.ToLower(str)
	matchCount := 0
	for i := range match {
		if strings.Contains(s, strings.ToLower(match[i])) {
			matchCount += 1
		}
	}
	return len(match) == matchCount
}

func downloadFile(url, filePath string, client *http.Client) error {
	stat, err := os.Stat(filePath)
	if err != nil && !os.IsNotExist(err) {
		return err
	} else if stat != nil {
		if stat.IsDir() {
			return errors.New("download path is a directory: " + filePath)
		}
		os.Remove(filePath)
	}
	output, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer output.Close()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	_, err = io.Copy(output, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func extractFileTarGz(targz, target string, match ...string) error {
	gzFile, err := os.Open(targz)
	if err != nil {
		return err
	}
	defer gzFile.Close()

	gzReader, err := gzip.NewReader(gzFile)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if containsAllMatches(header.Name, match...) {
			destFile, err := os.Create(target)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, tarReader)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return os.ErrNotExist
}
