package clipboard_test

import (
	"fmt"
	clipboard2 "github.com/zzztttkkk/pwgen/clipboard"
)

func Example() {
	clipboard2.WriteAll("日本語")
	text, _ := clipboard2.ReadAll()
	fmt.Println(text)

	// Output:
	// 日本語
}
