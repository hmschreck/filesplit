# Filesplit
A simple way to do files of a given line count
## Installation
`go get -u github.com/hmschreck/filesplit`

## How to use
```
import "github.com/hmschreck/filesplit"

...
fs := filesplit.FileSplitter{
  Pattern: "something-%d.csv",
  MaxLines: 1000,
}

fs.Start()
```

You can now write to the proper rollover file
by using the Write method of `fs`

### CSVs
There is a specific implementation for CSV files.

```
fs := filesplit.FileSplitter{
  Pattern: "something-%d.csv",
  MaxLines: 1000,
}

fs.Start()
fs.CsvStart()
```
fs.CsvWriter([]string) is now an automatic
rollover when there are too many lines.
