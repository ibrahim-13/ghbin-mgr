package util

type AppCtx struct {
	Conf  *Configuration
	State *InstallState
}

func NewAppCtx(installDir string) *AppCtx {
	conf := NewConfiguration(installDir)

	state := newInstallState(conf.InstallDir)
	ctx := &AppCtx{
		Conf:  conf,
		State: state,
	}
	return ctx
}
