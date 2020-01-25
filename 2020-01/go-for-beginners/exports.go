package main

// START INTERNAL OMIT
var notExported int

func internal() {
	// ...
}

type alsoInternal string

// END INTERNAL OMIT

// START EXPORTED OMIT

// ExportedConst is the answer
const ExportedConst = 42

// Exported is an exported function
func Exported() {

}

// AlsoExported is a custom string type
type AlsoExported string

// END EXPORTED OMIT
