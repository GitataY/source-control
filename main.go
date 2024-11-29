package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// check the command provided by the user
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command. Eg: source-control init")
		return
	}

	command := os.Args[1]

	if command == "init" {
		// check if the .repo folder already exists
		if _, err := os.Stat(".repo"); err == nil {
			fmt.Println("Repository already exists")
			return
		}

		// create the .repo folder
		err := os.Mkdir(".repo", 0755)
		if err != nil {
			fmt.Println("Error creating the .repo folder")
			return
		}

		// create the .repo subdirectories
		subdirs := []string{"objects", "refs/heads"}
		for _, subdir := range subdirs {
			err := os.MkdirAll(".repo/"+subdir, 0755)
			if err != nil {
				fmt.Println("Error creating the .repo subdirectories")
				return
			}
		}

		// create files: HEAD, config, description
		files := map[string]string{
			"HEAD":   "ref: refs/heads/main/n",
			"config": "[core]\n\trepositoryformatversion = 0\n\tfilemode = true\n\tbare = false\n",
			"description": "Unnamed repository; edit this file 'description' to name the repository.\n",
		}
		for name, content := range files {
			path := filepath.Join(".repo", name)
			err := os.WriteFile(path, []byte(content), 0644)
			if err != nil {
				fmt.Println("Error creating the .repo files")
				return
			}
		}
		
		fmt.Println("Initialized empty repository in ./.repo/")
	} else {
		fmt.Println("Unknown command")
		return
	}
}
