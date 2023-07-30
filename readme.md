# be kind
This package is a simple example how to use grettings

## installation

```bash
go get -u github.com/malisoft/grettings
```

## usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/malisoft/grettings"
)

func main() {
	log.SetPrefix("grettings: ")
	log.SetFlags(0) 
	merlo, error := grettings.Greet("Merlote")
	if (error != nil) {
		log.Fatal(error)
	}
	fmt.Println(merlo)

	names := []string{"merlote", "merlo", "Jhon"}
	greets, err := grettings.Greets(names)
	if (error != nil) {
		log.Fatal(err)
	}
	fmt.Println(greets)
}
```

