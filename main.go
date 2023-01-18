package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deepak/module_page/controllers"
	"github.com/deepak/module_page/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ps          services.PageService
	pc          controllers.PageController
	ctx         context.Context
	pagec       *mongo.Collection
	mongoclient *mongo.Client
	err         error
	ulrs        string = ":8080"
	mongo_uri   string = "mongodb://localhost:27017"
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI(mongo_uri)
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	pagec = mongoclient.Database("taskdb").Collection("pages")
	ps = services.NewPageService(pagec, ctx)
	pc = controllers.New(ps)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	//versioning api
	basepath := server.Group("/v1")
	pc.Routes(basepath)

	log.Fatal(server.Run(ulrs))

}
