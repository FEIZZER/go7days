package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	client := GetConnect()
	defer client.Disconnect(context.Background())
	db := client.Database("test")
	col := db.Collection("test")
	ctx := context.Background()
	//先在事务外写一条id为“111”的记录
	_, err := col.InsertOne(ctx, bson.M{"_id": "111", "name": "ddd", "age": 50})
	if err != nil {
		fmt.Println(err)
		return
	}

	session, err := db.Client().StartSession()
	if err != nil {
		fmt.Println(err)
		return
	}

	//开始事务
	err = session.StartTransaction()
	if err != nil {
		fmt.Println(err)
		return
	}

	//在事务内写一条id为“222”的记录
	_, err = col.InsertOne(ctx, bson.M{"_id": "222", "name": "ddd", "age": 50})
	if err != nil {
		fmt.Println(err)
		return
	}

	//写重复id
	_, err = col.InsertOne(ctx, bson.M{"_id": "111", "name": "ddd", "age": 50})
	if err != nil {
		fmt.Println(err)
		session.AbortTransaction(ctx)
	} else {
		session.CommitTransaction(ctx)
	}
}
