package main

import (
	"context"
	"fmt"
	"gin-mongodb-api/controllers"
	"gin-mongodb-api/repositories"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userRepository repositories.UserRepository
	userController controllers.UserController
	ctx            context.Context
	userCollection *mongo.Collection
	mongoClient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb+srv://shop_admin_2020:$hop__admin__2020@cluster0.5qhec.mongodb.net/?retryWrites=true&w=majority")
	mongoClient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err := mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection established")

	userCollection = mongoClient.Database("userdb").Collection("users")
	userRepository = repositories.NewUserRepository(userCollection, ctx)
	userController = controllers.New(userRepository)
	server = gin.Default()

}

func main() {

	defer mongoClient.Disconnect(ctx)

	basePath := server.Group("/v1")
	userController.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":9090"))
}
