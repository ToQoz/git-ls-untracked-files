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

func usage() {
	banner := "usage: git ls-untracked-files [<directory>]"

	fmt.Fprintf(os.Stderr, banner)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	targetDir := flag.Arg(0)

	cmd := exec.Command("git", []string{"ls-files", targetDir}...)
	result, err := cmd.Output()

	if err != nil {
		fmt.Println(string(result))
		log.Fatal(err)
	}

	trackedFiles := strings.Split(string(result), "\n")

	filepath.Walk("./"+targetDir, func(path string, info os.FileInfo, err error) error {

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
