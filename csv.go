package filesplit

import "encoding/csv"

func (fs *FileSplitter) CsvWriter(record []string) {
	fs.CSVWriter.Write(record)
	fs.CurrentLinesInFile += 1
	if fs.CurrentLinesInFile >= fs.MaxLines {
		fs.CSVWriter.Flush()
		fs.Rollover()
		fs.CSVWriter = csv.NewWriter(fs.CurrentFile)
	}
}

func (fs *FileSplitter) CsvStart() {
	fs.CSVWriter = csv.NewWriter(fs.CurrentFile)
}
