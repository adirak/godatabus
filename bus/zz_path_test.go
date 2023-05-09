package bus

import (
	"fmt"
	"testing"
)

func TestSplitPath(t *testing.T) {

	root := NewBus()

	path := "x.y.\"b.c\".z[1]"
	paths := root.splitPath(path)

	for i, obj := range paths {
		fmt.Printf("[%v] = %v\n", i, obj)
	}
	fmt.Println()

	if len(paths) != 4 {
		t.Error("paths invalid")
	}
}
