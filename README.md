This package is a simple wrapper for [validator](https://github.com/go-playground/validator).

### Installation
```
go get -u github.com/znbang/validation
```

### Example
```go
package main

import (
	"fmt"

	"github.com/znbang/validation"
)

func main() {
	title := ""
	url := "://github.com/znbang/validation"
	validate := validation.New()
	validate.Required(title, "title", "Please input the title.")
	validate.URL(url, "url", "Invalid URL.")
	for k, v := range validate.Errors {
		fmt.Println("Field:", k, ", Message:", v)
	}
}
```