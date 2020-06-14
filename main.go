package main

import (
	"fmt"
	"log"

	"github.com/olegsobchuk/athena_sync_go/athena"
	"github.com/olegsobchuk/athena_sync_go/athena/database"
)

func main() {
	apiConn := athena.Connection{}
	err := apiConn.New("hf9c6vkp8ne8s42ad4qc2n5p", "UceE6TZh8HpPXEX", "195900")
	if err != nil {
		log.Fatalln("New, err: ", err)
	}
	res, err := apiConn.GET("departments", map[string]string{})
	if err != nil {
		log.Fatalln("Get, err: ", err)
	}
	fmt.Printf("%T\n", res)
	dep := res.(map[string]interface{})["departments"].([]interface{})[0]

	fmt.Println(dep.(map[string]interface{})["departmentid"])
	fmt.Println("Athena API works!")

	db := database.DB
	db.Close()
}
