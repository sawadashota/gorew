package cmd

import (
	"bufio"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

type repository struct {
	name string
}

type installResult struct {
	success []repository
	errors  []repository
}

func install(targetDir string) {
	repositories := getRepos(targetDir)
	result := execInstall(repositories)
	result.stdoutErrors()
}

func execInstall(repositories []repository) *installResult {
	var result installResult

	for _, repo := range repositories {
		color.Green("go get -u %v", repo.name)
		err := exec.Command("go", "get", "-u", repo.name).Run()

		if err != nil {
			color.Red("Error occurred")
			result.errors = append(result.errors, repo)
			continue
		}

		result.success = append(result.success, repo)
	}

	return &result
}

func (i *installResult) stdoutErrors() {
	if len(i.errors) < 1 {
		return
	}

	color.Red("*** Error List ***")

	for _, repo := range i.errors {
		color.Red(repo.name)
	}
}

func getRepos(targetDir string) []repository {
	var repositories []repository

	if _, err := os.Stat(recordFilePath(targetDir)); err != nil {
		panic(err)
	}

	file, err := os.Open(recordFilePath(targetDir))

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		repositories = append(repositories, repository{name: scanner.Text()})
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return repositories
}
