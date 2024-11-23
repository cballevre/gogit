/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new, empty repository.",
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		workTree := getWorkTree(path)

		_, err := os.Stat(workTree)
		if err == nil {
			fmt.Println("Directory already exists")
			return
		}

		if os.IsNotExist(err) {
			fmt.Println("Creating repository at:", workTree)
			err := os.MkdirAll(workTree, 0755)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
		} else {
			fmt.Println("Error checking directory:", err)
			return
		}

		os.MkdirAll(workTree+"/branches", 0755)
		os.MkdirAll(workTree+"/objects", 0755)
		os.MkdirAll(workTree+"/refs/tags", 0755)
		os.MkdirAll(workTree+"/refs/heads", 0755)

		descriptionFile, _ := os.Create(workTree + "/description")
		defer descriptionFile.Close()
		descriptionFile.WriteString("Unnamed repository; edit this file 'description' to name the repository.\n")

		headFile, _ := os.Create(workTree + "/HEAD")
		defer headFile.Close()
		headFile.WriteString("ref: refs/heads/master\n")

		setupConfig(workTree)
	},
}

func setupConfig(workTree string) {
	inidata := ini.Empty()
	sec, err := inidata.NewSection("core")
	if err != nil {
		panic(err)
	}

	_, err = sec.NewKey("repositoryformatversion", "0")
	if err != nil {
		panic(err)
	}

	_, err = sec.NewKey("filemode", "true")
	if err != nil {
		panic(err)
	}

	_, err = sec.NewKey("bare", "false")
	if err != nil {
		panic(err)
	}

	_, err = sec.NewKey("logallrefupdates", "true")
	if err != nil {
		panic(err)
	}

	err = inidata.SaveTo(workTree + "/config")
	if err != nil {
		panic(err)
	}
}

func getWorkTree(path string) string {
	return path + "/.git"
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
