package cmd

import (
	"fmt"
	"os"
	"path"
)

const RecordFileName = ".gocmd"

func recordFile(gitDirs *[]gitDir, targetDir string) {
	recordFilePath := path.Join(targetDir, RecordFileName)

	file, err := os.OpenFile(recordFilePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, gitDir := range *gitDirs {
		fmt.Fprintln(file, gitDir.repo())
	}
}
