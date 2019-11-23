package cmd

import (
	"io/ioutil"
	"log"
	"regexp"

	"github.com/sawadashota/reposioty"
)

type commandDir struct {
	path string
}

func (c *commandDir) repo() string {
	r := regexp.MustCompile(`src\/(.+)$`)
	matches := r.FindStringSubmatch(c.path)

	if len(matches) <= 1 {
		log.Fatal("Cannot read repository name from path")
	}

	return matches[1]
}

func basename(path string) string {
	r := regexp.MustCompile(`\/([^\/]+)$`)
	matches := r.FindStringSubmatch(path)

	if len(matches) <= 1 {
		log.Fatal("Cannot read basename from path")
	}

	return matches[1]
}

func goCommands(srcPath string, binPath string) *[]commandDir {
	var commandDirs []commandDir

	repoPaths, err := reposioty.Get(srcPath)
	binNames := bins(binPath)

	if err != nil {
		panic(err)
	}

	for _, repoPath := range repoPaths {
		for _, mainDir := range *mainDirs(repoPath) {
			if inArray(basename(mainDir), binNames) {
				commandDirs = append(commandDirs, commandDir{path: mainDir})
			}
		}
	}

	return &commandDirs
}

func bins(binPath string) *[]string {
	files, err := ioutil.ReadDir(binPath)

	if err != nil {
		log.Fatal(err)
	}

	var bins []string
	for _, file := range files {
		if !file.IsDir() && file.Mode() >= 0655 {
			bins = append(bins, file.Name())
		}
	}

	return &bins
}

func inArray(needle string, haystack *[]string) bool {
	for _, val := range *haystack {
		if val == needle {
			return true
		}
	}
	return false
}
