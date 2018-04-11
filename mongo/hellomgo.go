package main

import (
    "encoding/json"
    "fmt"
//    "log"
//    "net/http"
    mgo "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)
func main() {

    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    defer session.Close()
    DB := session.DB("test")
    COL := DB.C("testColl")
  var bsonEPs []bson.M
    err = COL.Find(bson.M{}).All(&bsonEPs)
    if err != nil {
        fmt.Println(err)
    }

	jsonResp, merr := json.Marshal(bsonEPs)
    if merr != nil {
        fmt.Println(err)
    }
    fmt.Println(jsonResp)
      fmt.Printf("\nJSON DATA\n")
    fmt.Println(string(jsonResp))

    }