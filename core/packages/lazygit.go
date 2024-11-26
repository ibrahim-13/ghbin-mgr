package packages

type LazyGit struct {
}

func NewLazyGit() *LazyGit {
	return &LazyGit{}
}

func (p *LazyGit) GetId() string {
	return "jesseduffield/lazygit"
}

func (p *LazyGit) GetUser() string {
	return "jesseduffield"
}

func (p *LazyGit) GetRepository() string {
	return "jesseduffield"
}

func (p *LazyGit) OnInstall() error {
	return nil
}

func (p *LazyGit) OnUninstall() error {
	return nil
}

func (p *LazyGit) OnUpdate() error {
	return nil
}
