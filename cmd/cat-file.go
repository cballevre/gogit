package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// catFileCmd represents the cat-file command
var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Provide content or type and size information for repository objects.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cat-file called")
		fmt.Println("args:", args)
		test := findRepo(".")
		fmt.Println("repo path:", test)
		objectRead("test", args[1])
	},
}

// Find the current repository
func findRepo(path string) string {
	if path == "" {
		path = "."
	}

	// Check if the current directory is a repository
	workingTree := getWorkTree(path)
	fmt.Println("workingTree:", workingTree)

	return "test"
}

func objectRead(repo string, sha string) {
	// Read the object from the repository and print it to the console.
}

func init() {
	rootCmd.AddCommand(catFileCmd)
}
