package main

import (
 	"encoding/json"
	"github.com/mongodb/mongo-go-driver/mongo"
	"fmt"
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
)

func main() {
	client, err := mongo.NewClient("mongodb://localhost")
	if err != nil {
		fmt.Println(err)
	}
	db := client.Database("test")

	collection := db.Collection("testColl")

	result, err := collection.InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("item", "canvas"),
			bson.EC.Int32("qty", 100),
		))

	fmt.Printf("%#v\n", result)
	fmt.Printf("%#v\n", result.InsertedID)

	count, err := collection.Count(context.Background(), bson.NewDocument(
		bson.EC.Int32("qty", 100),
	))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T %v\n", count, count)
	cursor, err := collection.Find(
			context.Background(),
			bson.NewDocument(),
		)
	fmt.Printf("%#v\n", cursor)
	doc := bson.NewDocument()
	for cursor.Next(context.Background()) {
			doc.Reset()
			err := cursor.Decode(doc)
			if err !=nil  { 
				fmt.Println(err.Error())
			}
			fmt.Print("doc=%v\n",doc)

    		jsonResp, merr := json.Marshal(doc)
    		if merr != nil {
        	fmt.Println(err)
   				 } else {

   				fmt.Println("json=",jsonResp)
   			}

			keys, err := doc.Keys(false)
		
			fmt.Print("%v",keys)
			for index, value := range keys {
					fmt.Println("index=",index)
					fmt.Println("value=",value)
					got,err := doc.Lookup(value.Name)
					if err== nil {
						fmt.Println(got)
					}
			}
	

		}
	}

