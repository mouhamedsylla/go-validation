package main

import (
	"fmt"
	"github.com/mouhamedsylla/go-validation"
)

type User struct {
	Username string `validate:"username required"`
	Email    string `validate:"email required"`
	Age      int    `validate:"min(18) max(100)"`
	Password string `validate:"required"`
}

func main() {
	validator := validation.NewValidator()
	validator.Init(User{
		Username: "Mouhamed",
		Email:    "syllamouhamed99@gmail.com",
		Age:      2,
		Password: "12339432",
	})

	for _, target := range validator.Targets {
		fmt.Println("Name: ", target.Name)
		fmt.Println("Value: ", target.Value)
		fmt.Println("Tag: ", target.Tag)
		fmt.Println("=====================================")
	}

	fmt.Println(validator.Validate())
}
