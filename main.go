package main

import (
	"fmt"
	"os"
	rest "paper-cli/rest"

	"github.com/alexflint/go-arg"
)

var args struct {
	Import *ImportCmd `arg:"subcommand:import"`
	Me     *MeCmd     `arg:"subcommand:me"`
}

type ImportCmd struct {
	DocOrFolder string `arg:"positional"`
	ParentID    string `arg:"-p,--parent-id", "help=Parent folder UUID"`
	Host        string `arg:"env:PAPER_CLI__HOST, required"`
	Token       string `arg:"env:PAPER_CLI__TOKEN, required"`
}

type MeCmd struct {
	Host  string `arg:"env:PAPER_CLI__HOST, required"`
	Token string `arg:"env:PAPER_CLI__TOKEN, required"`
}

func RunImportCmd(cmd *ImportCmd) {
	file, err := os.Open(cmd.DocOrFolder)

	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		fmt.Printf("%s is directory\n", cmd.DocOrFolder)
	} else {
		fmt.Printf("%s is file\n", cmd.DocOrFolder)
		rest.Upload(
			&cmd.Host,
			&cmd.Token,
			&cmd.DocOrFolder,
			&cmd.ParentID)
	}
}

func RunMeCmd(cmd *MeCmd) {
	user := rest.Me(&cmd.Host, &cmd.Token)

	fmt.Printf("ID: %s\n", user.ID)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Inbox folder ID: %s\n", user.InboxFolderID)
	fmt.Printf("Home folder ID: %s\n", user.HomeFolderID)
}

func main() {
	arg.MustParse(&args)

	switch {
	case args.Import != nil:
		RunImportCmd(args.Import)
	case args.Me != nil:
		RunMeCmd(args.Me)
	}

}
