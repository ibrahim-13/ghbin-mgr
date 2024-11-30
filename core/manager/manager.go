package manager

import (
	"errors"
	"gbm/core/release"
	"gbm/util"
	"net/http"
	"os"
	"path/filepath"
)

type Manager struct {
	ctx              *util.AppCtx
	Packages         []Package
	releaseInfoCache map[string]*release.GhReleaseInfo
	Release          release.GhRelease
}

func (m *Manager) DownloadAndExtract(conf *util.Configuration, info *release.GhReleaseInfo) error {
	var downloadUrl string
	for i := range info.Assets {
		if containsAllMatches(info.Assets[i].Name, "darwin", "arm64", "tar.gz") {
			downloadUrl = info.Assets[i].BrowserDownloadURL
			break
		}
	}
	if downloadUrl == "" {
		return errors.New("no asset found for os: macos")
	}
	archivefile := filepath.Join(conf.TmpDir, "lazygit.tar.gz")
	err := downloadFile(downloadUrl, archivefile, http.DefaultClient)
	if err != nil {
		return err
	}
	binfile := filepath.Join(conf.InstallDir, "lazygit")
	err = extractFileTarGz(archivefile, binfile, "lazygit")
	if err != nil {
		return err
	}
	err = os.Chmod(binfile, 0755)
	if err != nil {
		return err
	}
	err = os.Remove(archivefile)
	if err != nil {
		return err
	}
	return nil
}

func NewManager(ctx *util.AppCtx) *Manager {
	return &Manager{
		ctx:              ctx,
		Release:          release.NewRelease(),
		releaseInfoCache: make(map[string]*release.GhReleaseInfo),
	}
}
