package manager

import (
	"gbm/core/release"
	"gbm/util"
)

type Package struct {
	State *util.BinaryState
}

func (p *Package) CreateBinaryInfo(info *release.GhReleaseInfo) util.BinaryState {
	return util.BinaryState{
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
