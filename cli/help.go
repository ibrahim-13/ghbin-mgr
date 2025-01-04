package cli

import "fmt"

func PrintHelp() {
	fmt.Print("ghbin-mgr manage binaries of github releases\n\n")
	fmt.Println("comands:")
	fmt.Println("    info      get release information")
	fmt.Println("    check     check for update")
	fmt.Println("    install   install binary from latest release")
	fmt.Println("    installx  install binary from latest release archive")
	fmt.Println("    script    run multiple commands with a script")
}
