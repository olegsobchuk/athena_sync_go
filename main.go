package main

import (
	"fmt"
	"log"

	"github.com/olegsobchuk/athena_sync_go/athena"
)

func main() {
	apiConn := athena.Connection{}
	err := apiConn.New("hf9c6vkp8ne8s42ad4qc2n5p", "UceE6TZh8HpPXEX", "195900")
	if err != nil {
		log.Fatalln("New, err: ", err)
	}
	res, err := apiConn.GET("practiceinfo", map[string]string{})
	if err != nil {
		log.Fatalln("Get, err: ", err)
	}
	fmt.Printf("%v\n", res)
	fmt.Println("It works!")
}
