package manager

import (
	"errors"
	"gbm/core/release"
	"gbm/util"
)

type Manager struct {
	ctx              *util.AppCtx
	Packages         []Package
	releaseInfoCache map[string]*release.GhReleaseInfo
	Release          release.GhRelease
}

func (m *Manager) Register(pkg PackageDescription) error {
	pkgId := pkg.GetId()
	for i := range m.Packages {
		if m.Packages[i].Description.GetId() == pkgId {
			return errors.New("package already registered: " + pkg.GetId())
		}
	}
	m.Packages = append(m.Packages, NewPackage(pkg, m.ctx.State.FindById(pkgId)))
	return nil
}

func (m *Manager) UpdateAllReleaseInfo() error {
	for i := range m.Packages {
		info, err := m.Release.GetRelease(m.Packages[i].Description.GetUser(), m.Packages[i].Description.GetRepository())
		if err != nil {
			return err
		}
		m.releaseInfoCache[m.Packages[i].Description.GetId()] = info
	}
	return nil
}

func (m *Manager) UpdateReleaseInfo(pkgIds []string) error {
	var pkgs []Package
	for i := range m.Packages {
		for j := range pkgIds {
			if m.Packages[i].Description.GetId() == pkgIds[j] {
				pkgs = append(pkgs, m.Packages[i])
			}
		}
	}
	for i := range pkgs {
		info, err := m.Release.GetRelease(pkgs[i].Description.GetUser(), pkgs[i].Description.GetRepository())
		if err != nil {
			return err
		}
		m.releaseInfoCache[pkgs[i].Description.GetId()] = info
	}
	return nil
}

func (m *Manager) GetReleaseInfo(pkg *Package) (*release.GhReleaseInfo, error) {
	pkgId := pkg.Description.GetId()
	v, ok := m.releaseInfoCache[pkgId]
	if ok {
		return v, nil
	}
	info, err := m.Release.GetRelease(pkg.Description.GetUser(), pkg.Description.GetRepository())
	if err != nil {
		return nil, err
	}
	m.releaseInfoCache[pkgId] = info
	return info, nil
}

func (m *Manager) Install(id string) error {
	if m.ctx.State.Exists(id) {
		return m.Update(id)
	}
	for i := range m.Packages {
		if pkgId := m.Packages[i].Description.GetId(); pkgId == id {
			info, err := m.GetReleaseInfo(&m.Packages[i])
			if err != nil {
				return err
			}
			binState := m.Packages[i].CreateBinaryInfo(info)
			err = m.Packages[i].Description.OnInstall(m.ctx.Conf, info)
			if err != nil {
				return err
			}
			return m.ctx.State.Add(binState)
		}
	}
	return errors.New("package not registered: " + id)
}

func (m *Manager) Update(id string) error {
	if !m.ctx.State.Exists(id) {
		return errors.New("package is not installed: " + id)
	}
	for i := range m.Packages {
		if pkgId := m.Packages[i].Description.GetId(); pkgId == id {
			info, err := m.GetReleaseInfo(&m.Packages[i])
			if err != nil {
				return err
			}
			if !m.Packages[i].HasUpdate(info) {
				return errors.New("already up to date: " + id)
			}
			binState := m.Packages[i].CreateBinaryInfo(info)
			err = m.Packages[i].Description.OnUpdate(m.ctx.Conf, info)
			if err != nil {
				return err
			}
			return m.ctx.State.Update(binState)
		}
	}
	return nil
}

func (m *Manager) Uninstall(id string) error {
	if !m.ctx.State.Exists(id) {
		return errors.New("package is not installed: " + id)
	}
	for i := range m.Packages {
		if pkgId := m.Packages[i].Description.GetId(); pkgId == id {
			err := m.Packages[i].Description.OnUninstall(m.ctx.Conf)
			if err != nil {
				return err
			}
			return m.ctx.State.Remove(pkgId)
		}
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
