package main

import (
	"fmt"
	"os"
	rest "paper-cli/rest"

	"github.com/alexflint/go-arg"
)

var args struct {
	DocOrFolder string `arg:"positional"`
	ParentID    string `arg:"-p,--parent-id", "help=Parent folder UUID"`
	Host        string `arg:"env:PAPER_CLI__HOST, required"`
	Token       string `arg:"env:PAPER_CLI__TOKEN, required"`
}

func main() {
	arg.MustParse(&args)

	file, err := os.Open(args.DocOrFolder)

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
		fmt.Printf("%s is directory\n", args.DocOrFolder)
	} else {
		fmt.Printf("%s is file\n")
		rest.Upload(
			&args.Host,
			&args.Token,
			&args.DocOrFolder,
			&args.ParentID)
	}
}
