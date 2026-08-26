package rtl_test

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/rtl"
)

// ExampleProcess demonstrates how to shape and reorder Arabic text so a left to
// right writer, such as a PDF text operator, draws it correctly.
func ExampleProcess() {
	// Process expects a single already wrapped line. A paragraph must be broken
	// into lines first, and each line processed on its own.
	visual := rtl.Process("مرحبا بالعالم")

	fmt.Println(visual)
}

// ExampleContainsArabic demonstrates how to check whether a text holds any
// Arabic character before applying right-to-left handling.
func ExampleContainsArabic() {
	fmt.Println(rtl.ContainsArabic("مرحبا بالعالم"))
	fmt.Println(rtl.ContainsArabic("hello world"))
}
