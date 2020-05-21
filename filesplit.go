package filesplit

import (
	"encoding/csv"
	"fmt"
	"os"
	log "github.com/sirupsen/logrus"
)

var Permissions = os.FileMode(0644)
var HandlePermissions = os.O_TRUNC | os.O_CREATE

type FileSplitter struct {
	Initialized bool
	MaxLines int
	CurrentLinesInFile int
	CurrentOffset int
	CurrentFile *os.File
	Pattern string
	CSVWriter *csv.Writer
}

func (fs *FileSplitter) Write(input []byte) (n int, err error) {
	if fs.CurrentLinesInFile >= fs.MaxLines {
		fs.Rollover()
		fs.CurrentLinesInFile = 0
	}
	n, err = fs.CurrentFile.Write(input)
	if err != nil {
		fmt.Println(string(input))
		fmt.Println("write: ", err)
		return
	}
	fs.CurrentLinesInFile += 1
	return
}

func (fs *FileSplitter) Close() (err error){
	err = fs.CurrentFile.Close()
	return
}

func (fs *FileSplitter) Start() (err error) {
	fs.CurrentOffset = 1
	filename, err := fs.GetFilename()
	fs.CurrentFile, err = os.OpenFile(filename, HandlePermissions, Permissions)
	return
}

func (fs *FileSplitter) Rollover() (err error) {
	log.Debug("rolling over")
	fs.Close()
	fs.CurrentOffset += 1
	filename, err := fs.GetFilename()
	fs.CurrentFile, err = os.OpenFile(filename, HandlePermissions, Permissions)
	fs.CurrentLinesInFile = 0
	return
}

func (fs *FileSplitter) ShowFilename() {
	fmt.Println(fs.CurrentFile.Name())
}

func (fs *FileSplitter) GetFilename() (filename string, err error) {
	return fmt.Sprintf(fs.Pattern, fs.CurrentOffset), nil
}

