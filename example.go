package main

import (
	"fmt"
	version "github.com/yakumioto/canonical-version-go-example"
	version2 "github.com/yakumioto/canonical-version-go-example/v2"
)

func main() {
	fmt.Println(version.Version())
	fmt.Println(version2.Version())
}