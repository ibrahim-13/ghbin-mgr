package packages

import (
	"gbm/core/manager"
	"net/http"
	"time"
)

func RegisterPackages(m *manager.Manager) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	m.Register(NewLazyGit(client))
}
