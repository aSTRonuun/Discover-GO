package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// Struct to JSON
	p1 := product{1, "Notebook", 1999.90, []string{"Eletronic", "Office"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// Json to Struct
	var p2 product
	jsonString := `{"id":2,"name":"Smartphone","price":999.90,"tags":["Eletronic","Mobile", "Entreteniment"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
