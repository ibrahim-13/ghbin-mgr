package packages

import (
	"errors"
	"gbm/core/release"
	"gbm/util"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type LazyGit struct {
	client *http.Client
}

func NewLazyGit(client *http.Client) *LazyGit {
	return &LazyGit{
		client: client,
	}
}

func (p *LazyGit) GetId() string {
	return "jesseduffield/lazygit"
}

func (p *LazyGit) GetUser() string {
	return "jesseduffield"
}

func (p *LazyGit) GetRepository() string {
	return "lazygit"
}

func (p *LazyGit) OnInstall(conf *util.Configuration, info *release.GhReleaseInfo) error {
	switch conf.OsType {
	case util.OsMac:
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
		err := downloadFile(downloadUrl, archivefile, p.client)
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
	default:
		return errors.New("platform not supported: " + runtime.GOOS)
	}
	return nil
}

func (p *LazyGit) OnUninstall(conf *util.Configuration) error {
	return nil
}

func (p *LazyGit) OnUpdate(conf *util.Configuration, info *release.GhReleaseInfo) error {
	return nil
}
