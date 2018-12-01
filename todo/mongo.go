package main

import (
	"log"

	"github.com/globalsign/mgo"
)

func main() {

	url := "mongodb://dbuser:Password@cluster0-shard-00-00-jersx.mongodb.net:27017,cluster0-shard-00-01-jersx.mongodb.net:27017,cluster0-shard-00-02-jersx.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin"

	database := "tech_inno"
	collectoin := "test"

	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(database).C(collectoin)
	err = c.Insert(map[string]string{"language": "golang"})
	err = c.Insert(map[string]string{"language": "java"})
}
