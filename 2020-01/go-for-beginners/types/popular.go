package popular

// START io.Reader OMIT

// io.Reader
type Reader interface {
	Read(p []byte) (n int, err error)
}

// END io.Reader OMIT

// START io.Writer OMIT

// io.Writer
type Writer interface {
	Write(p []byte) (n int, err error)
}

// END io.Writer OMIT

// START sort.Interface OMIT

// sort.Interface
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// END sort.Interface OMIT
