package routes

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not save user.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfull.",
	})

}
