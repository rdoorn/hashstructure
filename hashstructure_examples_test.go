package hashstructure

import (
	"fmt"
)

func ExampleHash() {
	type ComplexStruct struct {
		Name     string
		Age      uint
		Metadata map[string]interface{}
	}

	v := ComplexStruct{
		Name: "mitchellh",
		Age:  64,
		Metadata: map[string]interface{}{
			"car":      true,
			"location": "California",
			"siblings": []string{"Bob", "John"},
		},
	}

	hash, err := Hash(v, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x", hash)
	// Output:
	// 62aace8c92f6ea9c4699a62da03c08c1458ebbf0331210a8a0e6b43536c59c68
}
