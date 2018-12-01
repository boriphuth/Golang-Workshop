package main

import (
	"crypto/tls"
	"log"
	"net"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

type todo1 struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Topic string        `json:"topic"`
	Done  bool          `json:"done"`
}

func main1() {

	url := "mongodb://dbuser1234:Password1234@cluster0-shard-00-00-jersx.mongodb.net:27017,cluster0-shard-00-01-jersx.mongodb.net:27017,cluster0-shard-00-02-jersx.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin"

	// url := "localhost:2277"
	database := "tech_inno"
	collectoin := "test"

	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	dialInfo, err := mgo.ParseURL(url)
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		return tls.Dial("tcp", addr.String(), tlsConfig)
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// session, err := mgo.Dial(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer session.Close()

	c := session.DB(database).C(collectoin)
	err = c.Insert(todo1{ID: bson.NewObjectId(), Topic: "golang"})
	// err = c.Insert(map[string]string{"language": "golang"})
	err = c.Insert(todo1{ID: bson.NewObjectId(), Topic: "java"})
}
