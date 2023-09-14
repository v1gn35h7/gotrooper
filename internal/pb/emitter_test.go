package pb

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/v1gn35h7/gotrooper/internal/goshell"
)

func TestSaveOutputToPb(t *testing.T) {
	output := "Test output with long string...."
	homeDir, _ := os.UserHomeDir()
	outputFilePath := filepath.Join(homeDir, "gotrooper.log")
	file, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer file.Close()

	if err != nil {
		t.Error(err)
	}

	outputFile := &goshell.OutputFile{
		File: file,
	}

	for i := 0; i < 100; i++ {
		err := SaveOutputToPb(output, outputFile)

		if err != nil {
			t.Error(err)
			t.Fail()
		}
	}

}
