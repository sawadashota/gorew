package cmd

import (
	"path"
)

const RecordFileName = ".gorew"

func recordFilePath(targetDir string) string {
	return path.Join(targetDir, RecordFileName)
}
