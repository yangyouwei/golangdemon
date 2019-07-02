package main

import (
"context"
"fmt"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
"go.mongodb.org/mongo-driver/mongo/readpref"
"time"
)

var mongodburl string

type User_str struct {
	Id  primitive.ObjectID  `bson:"_id"`
	Book_id int32	`bson:"book_id"`
	Text string `bson:"text"`
}

var (
	err             error
	client          *mongo.Client
	collection      *mongo.Collection
	insertOneRes    *mongo.InsertOneResult
	insertManyRes   *mongo.InsertManyResult
	delRes          *mongo.DeleteResult
	updateRes       *mongo.UpdateResult
	cursor          *mongo.Cursor
	user_str		User_str
	size            int64
)

func main()  {
	//mongodburl = "mongodb://192.168.1.28:27017/novel?authMechanism=SCRAM-SHA-1"
	mongodburl = "mongodb://192.168.1.28:27017/"
	client, err := mongo.Connect(getContext(), options.Client().ApplyURI(mongodburl))
	if err != nil {
		panic(err)
	}
	err = client.Ping(getContext(), readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}
	collection = client.Database("novel").Collection("countent")
	collection.FindOne(getContext(),bson.D{}).Decode(&user_str)
	fmt.Println(user_str.Text)
	fmt.Println(user_str.Id)

}

func getContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}

