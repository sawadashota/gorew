package cmd

import (
	"github.com/sawadashota/reposioty"
	"io/ioutil"
	"log"
	"regexp"
)

type gitDir struct {
	path string
}

func (g *gitDir) repo() string {
	r := regexp.MustCompile(`src\/(.+)$`)
	matches := r.FindStringSubmatch(g.path)

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

func goCommands(srcPath string, binPath string) *[]gitDir {
	var gitDirs []gitDir

	repoPaths, err := reposioty.Get(srcPath)
	binNames := bins(binPath)

	if err != nil {
		panic(err)
	}

	for _, repoPath := range repoPaths {
		if inArray(basename(repoPath), &binNames) {
			gitDirs = append(gitDirs, gitDir{path: repoPath})
		}
	}

	return &gitDirs
}

func bins(binPath string) []string {
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

	return bins
}

func inArray(needle string, haystack *[]string) bool {
	for _, val := range *haystack {
		if val == needle {
			return true
		}
	}
	return false
}
