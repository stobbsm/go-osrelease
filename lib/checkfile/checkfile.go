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

// SizeLess checks the given files reported size
// against the given size.
// If the files size is less then the given size, return true
// return false if the file cannot be read, or size is larger
func SizeLess(filename string, size int64) bool {
	fstat, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return fstat.Size() < size
}

// SizeGreater checks the given files reported size
// against the given size.
// If the files size is greater then the given size,
// return true.
func SizeGreater(filename string, size int64) bool {
	fstat, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return fstat.Size() > size
}

// SizeEqual checks the given files reported size
// against the given size.
// If the sizes are equal, return true. Everything
// else returns false
func SizeEqual(filename string, size int64) bool {
	fstat, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return fstat.Size() == size
}
