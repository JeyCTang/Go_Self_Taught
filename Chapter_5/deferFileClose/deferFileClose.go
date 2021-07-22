package main

import "os"

func fileSize(filename string) int64 {
	// open the file according to the file name
	f, err := os.Open(filename)
	// if error occurred and return 0
	if err != nil {
		return 0
	}
	// get the info of file
	info, err := f.Stat()
	// if error occurred, close the file and return 0
	if err != nil {
		f.Close()
		return 0
	}
	// get the size of file
	size := info.Size()
	// close the file
	f.Close()
	// return the size of file
	return size
}

// equivalent code using defer
func fileSize2(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	// calling return 0 with defer, at this time close won't be called
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		// trigger defer, call Close()
		return 0
	}
	size := info.Size()
	// trigger of defer, call Close()
	return size
}
