package main

import "fmt"

func main() {

	{
		// START LIT OMIT
		person := struct {
			Firstname string
			Lastname  string
		}{
			Firstname: "Christopher",
			Lastname:  "Hlubek",
		}

		fmt.Println(person)
		// END LIT OMIT

		_ = person
	}

	{
		// START TYPE OMIT
		type Person struct {
			Firstname string
			Lastname  string
		}

		person := Person{
			Firstname: "Christopher",
			Lastname:  "Hlubek",
		}

		var otherPerson Person
		otherPerson.Firstname = "John"

		// END TYPE OMIT
		_ = person
	}
}
