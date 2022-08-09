package controllers

import (
	"gin-mongodb-api/models"
	"gin-mongodb-api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repositories.UserRepository
}

func New(repo repositories.UserRepository) UserController {
	return UserController{
		userRepository: repo,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := uc.userRepository.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{"message": "success"})

}

func (uc *UserController) GetUser(ctx *gin.Context) {

	id := ctx.Param("id")
	user, err := uc.userRepository.GetUser(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(200, gin.H{"message": user})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := uc.userRepository.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{"message": user})
}

func (uc *UserController) GetAll(ctx *gin.Context) {

	users, err := uc.userRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(200, gin.H{"message": users})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {

	id := ctx.Param("id")
	err := uc.userRepository.DeleteUser(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.PUT("/update", uc.UpdateUser)
	userRoute.DELETE("/delete/:id", uc.DeleteUser)
	userRoute.GET("/get/:id", uc.GetUser)
	userRoute.GET("/fetchAll", uc.GetAll)

}
