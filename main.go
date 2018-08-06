package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/GTD/cmd"
	"github.com/GTD/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()
	must(err)
	dbPath := filepath.Join(home, "GTD.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
