package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

var nameApp string
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "creates the initial structure of the project",
	Run: func(cmd *cobra.Command, args []string) {
		nameApp = os.Args[2]
		createMainDirectory(nameApp)
		createFileMain()
		createSubDirectory("internal")
		createSubDirectory("pkg")

		command := exec.Command("go", "mod", "init", nameApp)
		command.Dir = nameApp
		err := command.Run()
		if err != nil {
			panic(err)
		}
	},
}

func createFileMain() {
	fileMain := nameApp + "/cmd/" + nameApp + "/main.go"
	err := os.MkdirAll(filepath.Dir(fileMain), os.ModePerm)
	if err != nil {
		panic(err)
	}

	mainFile, err := os.Create(fileMain)
	if err != nil {
		panic(err)
	}

	_, err = mainFile.WriteString(`package main

func main() {

}
`)
	if err != nil {
		panic(err)
	}
	defer mainFile.Close()
}

func createMainDirectory(directory string) {
	createDirectory(directory, true)
}

func createSubDirectory(directory string) {
	createDirectory(directory, false)
}

func createDirectory(directory string, main bool) {
	if main {
		if err := os.MkdirAll(nameApp, os.ModePerm); err != nil {
			panic(err)
		}
		return
	}

	err := os.MkdirAll(nameApp+"/"+directory, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(createCmd)
}
