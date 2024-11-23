package checkfile

import (
	"os"
)

// Readable checks if a file is readable by attempting
// to open the file for read access. Failure returns false,
// otherwise the file is readable
func Readable(filename string) bool {
	var f *os.File
	var err error
	if f, err = os.Open(filename); err != nil {
		return false
	}
	f.Close()
	return true
}

// Exists is a simple helper function that
// returns true if the filename exists,
// false otherwise
func Exists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}
	return true
}
