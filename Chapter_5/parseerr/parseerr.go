package main

import "fmt"

// ParseError state a parse error
type ParseError struct {
	Filename string // filename
	Line     int    // which line
}

// interface of error, return the description of error
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// create some Parse errors
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}
func main() {
	var e error
	// create a error instance, including the filename and line number
	e = newParseError("main.go", 1)
	// check the description of e via the interface of error
	fmt.Println(e.Error())
	// get the detailed info via the specific error type
	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
	default:
		fmt.Println("other error")
	}
}
