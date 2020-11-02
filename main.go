package main

import (
	"fmt"

	"github.com/lithammer/shortuuid"
)

func main() {
	u := shortuuid.New()
	fmt.Printf("#ID:%s", u)
}
