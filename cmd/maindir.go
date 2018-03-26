package cmd

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func mainDirs(gitRoot string) []string {
	files, err := ioutil.ReadDir(gitRoot)
	if err != nil {
		panic(err)
	}

	var paths []string

	if hasMainGo(gitRoot) {
		paths = append(paths, gitRoot)
	}

	for _, file := range files {
		abs := path.Join(gitRoot, file.Name())

		if !file.IsDir() || file.Name() == "vendor" {
			continue
		}

		if hasMainGo(abs) {
			paths = append(paths, abs)
			continue
		}

		paths = append(paths, mainDirs(abs)...)
	}

	return paths
}

func hasMainGo(dirPath string) bool {
	files, err := ioutil.ReadDir(dirPath)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if isMainGo(path.Join(dirPath, file.Name())) {
			return true
		}
	}

	return false
}

func isMainGo(filepath string) bool {
	if !isMainGoFileName(filepath) {
		return false
	}

	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 32*1024)

	for {
		lineByte, _, err := reader.ReadLine()
		line := string(lineByte)

		if strings.HasPrefix(line, "package") {
			return isPackageMain(line)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return false
}

func isMainGoFileName(fileName string) bool {
	return strings.HasSuffix(fileName, "main.go")
}

func isPackageMain(line string) bool {
	return line == "package main"
}
