package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed inputs/test.txt
var inputs embed.FS

func main() {
	fmt.Println("This is the start of the year 2025.")

}
