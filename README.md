# Validation

A Go package for validating various input fields using a set of predefined rules. This library provides a flexible and extensible way to validate form inputs and other data structures.

## Features

- **Required**: Ensures a field is not empty.
- **Email**: Validates email addresses.
- **Username**: Validates usernames (alphanumeric and underscores, length 3-20).
- **MaxLength**: Validates the maximum length of a string.
- **MinLength**: Validates the minimum length of a string.
- **MinNumeric**: Validates the minimum value of a numeric field.
- **MaxNumeric**: Validates the maximum value of a numeric field.

## Installation

To install the package, use:

```sh
go get github.com/yourusername/validation
```

## Usage
### Define Your Struct
Define a struct with validation tags:

```
type User struct {
    Username string `validate:"required username"`
    Email    string `validate:"required email"`
    Age      int    `validate:"min 18 max 65"`
}
```
### Initialize the Validator
Create a new **Validator** instance and initialize it with your struct:

```
package main

import (
    "fmt"
    "github.com/mouhamedsylla/validation"
)

func main() {
    user := User{
        Username: "johndoe",
        Email:    "john.doe@example.com",
        Age:      30,
    }

    validator := validation.NewValidator()
    validator.Init(user)

    if err := validator.Validate(); err != nil {
        fmt.Println("Validation failed:", err)
    } else {
        fmt.Println("Validation succeeded")
    }
}
```

### Contributing
Contributions are welcome! Please submit a pull request or open an issue to discuss your ideas.


