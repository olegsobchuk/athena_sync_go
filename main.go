package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/olegsobchuk/athena_sync_go/athena"
	"github.com/olegsobchuk/athena_sync_go/athena/database"
)

func main() {
	db := database.DB
	defer db.Close()

	apiConn := athena.Connection{}
	err := apiConn.New("hf9c6vkp8ne8s42ad4qc2n5p", "UceE6TZh8HpPXEX", "195900")
	if err != nil {
		log.Fatalln("New, err: ", err)
	}
	res, err := apiConn.GET("departments", map[string]string{})
	if err != nil {
		log.Fatalln("Get, err: ", err)
	}
	deps := res.(map[string]interface{})["departments"].([]interface{})
	for _, dep := range deps {
		depJSON, _ := json.Marshal(dep)
		err = database.Insert("departments", depJSON, dep.(map[string]interface{})["departmentid"].(string))
		if err != nil {
			log.Fatalln("can't save department to DB, err: ", err)
		}
	}

	fmt.Println("Athena API works!")
}
