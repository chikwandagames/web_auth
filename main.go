package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Marshalling
// Tacking go code and turning into JSON

type person struct {
	// Capital "F" for exporting
	First string
}

func main() {
	p1 := person{
		First: "John",
	}

	p2 := person{
		First: "Trev",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp)

	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
