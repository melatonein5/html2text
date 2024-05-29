package entity_test

import (
	"fmt"

	"github.com/melatonein5/html2text/src/entity"
)

// ExampleGetEntity will return a rune for a given entity
func ExampleGetEntity() {
	e := entity.GetEntity("A")
	fmt.Println(e)
	e = entity.GetEntity("zwnj")
	fmt.Println(e)
	//Output:
	// 0
	// 8204
}
