package main

// START OMIT
type Person struct {
	Firstname string
	Lastname  string
}

func (p Person) Name() string {
	return p.Firstname + " " + p.Lastname
}

// END OMIT
