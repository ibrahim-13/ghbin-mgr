package manager

import (
	"gbm/core/release"
	"gbm/util"
)

type PackageDescription interface {
	GetId() string
	GetUser() string
	GetRepository() string
	OnInstall() error
	OnUninstall() error
	OnUpdate() error
}

type Package struct {
	Description PackageDescription
	State       *util.BinaryState
}

func (p *Package) CreateBinaryInfo(info *release.GhReleaseInfo) util.BinaryState {
	return util.BinaryState{
		Id:        p.Description.GetId(),
		Version:   info.TagName,
		UpdatedAt: info.PublishedAt.String(),
	}
}

func (p *Package) HasUpdate(info *release.GhReleaseInfo) bool {
	newBinInfo := p.CreateBinaryInfo(info)
	return p.State.Id != newBinInfo.Id ||
		p.State.Version != newBinInfo.Version ||
		p.State.UpdatedAt != newBinInfo.UpdatedAt
}

func NewPackage(desc PackageDescription, state *util.BinaryState) *Package {
	return &Package{
		Description: desc,
		State:       state,
	}
}
