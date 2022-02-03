package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clinet, err := mongo.NewClient(options.Client().ApplyURT(""))
	fmt.Println(clinet, err)
}
