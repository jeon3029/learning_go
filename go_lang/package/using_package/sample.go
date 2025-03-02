package main

import (
	"fmt"
	"github.com/huandu/xstrings"
)

func main() {
	fmt.Println(xstrings.Reverse("Hello"))      // "olleH"
	fmt.Println(xstrings.Shuffle("abcdefg"))    // Random shuffle like "gdfcabe"
	fmt.Println(xstrings.Center("Go", 10, "*")) // "***Go*****"
}
