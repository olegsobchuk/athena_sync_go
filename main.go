package main

import (
	"fmt"

	"github.com/olegsobchuk/athena_sync_go/athena"
)

func main() {
	apiConn := athena.Connection{}
	err := apiConn.New("hf9c6vkp8ne8s42ad4qc2n5p", "UceE6TZh8HpPXEX", "195900")
	fmt.Println(err)
	fmt.Printf("%v\n", apiConn)
	fmt.Println("It works!")
}
