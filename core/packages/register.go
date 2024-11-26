package packages

import "gbm/core/manager"

func RegisterPackages(m *manager.Manager) {
	m.Register(NewLazyGit())
}
