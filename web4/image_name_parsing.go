package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type filename struct {
	Num      int    `bson: "num"`
	Category string `bson: "category"`
	Name     string `bson: "name"`
	Url      string `bson: "url"`
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func mongoConn() (client *mongo.Client) {
	credential := options.Credential{
		Username: "test",
		Password: "test",
	}
	clientOptions := options.Client().ApplyURI("mongodb://15.165.15.37:27017").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	CheckErr(err)

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Made")

	return client
}

func mongoDisConn(client *mongo.Client) {

	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}

func extractFileInfo(info fs.FileInfo, targetCollection *mongo.Collection) {
	str := info.Name()
	i, j, k := strings.LastIndex(str, "_"), strings.LastIndex(str, " "), strings.LastIndex(str, ".")
	num := str[:i]
	newint, _ := strconv.Atoi(num)
	var name string
	if j == -1 {
		name = str[i+1 : k]
	} else {
		name = str[j+1 : k]
	}

	url := "15.165.15.37:5000/image/" + str

	data := filename{Num: newint, Category: "예쁜아이돌", Name: name, Url: url}

	_, err := targetCollection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

}

func main() {
	mongo := mongoConn()
	targetCollection := mongo.Database("test").Collection("FileInfo")
	dirname := "./image/1_예쁜 여자 아이돌"

	entries, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	infos := make([]fs.FileInfo, 0, len(entries))
	//Directory에 있는 모든 File 정보 받아오기
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			log.Fatal(err)
		}
		infos = append(infos, info)
	}

	//Parsing 해서 Db에 데이터 입력
	for _, info := range infos {
		extractFileInfo(info, targetCollection)
	}

	defer mongoDisConn(mongo)
}
