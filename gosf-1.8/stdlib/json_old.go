package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	// START OMIT
	type Person struct {
		Name     string
		AgeYears int
		SSN      int
	}

	getPerson := func() Person { // OMIT
		return Person{Name: "Chris"} // OMIT
	} // OMIT
	var person Person = getPerson()

	type JSONPerson struct {
		Name     string `json:"full_name"`
		AgeYears int    `json:"age"`
		SSN      int    `json:"-"`
	}

	b, err := json.Marshal(JSONPerson(person))
	// END OMIT

	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(b)
}
