package cmd

import (
	"fmt"
	"os"
	"path"
)

const RecordFileName = ".gorew"

func recordFile(gitDirs *[]commandDir, targetDir string) {
	file, err := os.OpenFile(recordFilePath(targetDir), os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, gitDir := range *gitDirs {
		fmt.Fprintln(file, gitDir.repo())
	}
}

func recordFilePath(targetDir string) string {
	return path.Join(targetDir, RecordFileName)
}
