# ordered-object
JSON objects that respect insertion order.

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	orderedobject "github.com/SolarSystems-Software/ordered-object"
)

func main() {
	obj := new(orderedobject.Object[int])
	obj.Set("b", 0)
	obj.Set("a", 0)

	encoded, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	// outputs: {"b":0,"a":0}
	fmt.Println(string(encoded))
}
```