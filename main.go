package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()

	cmd := exec.Command("git", []string{"ls-files", flag.Arg(0)}...)
	result, err := cmd.Output()

	if err != nil {
		fmt.Println(string(result))
		log.Fatal(err)
	}

	trackedFiles := strings.Split(string(result), "\n")

	filepath.Walk("./"+flag.Arg(0), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			// skip .git dir
			if path == ".git" {
				return filepath.SkipDir
			}

			// nop
			return nil
		}

		// skip tracked file
		for _, tf := range trackedFiles {
			if path == tf {
				return nil
			}
		}

		fmt.Println(path)
		return nil
	})
}
