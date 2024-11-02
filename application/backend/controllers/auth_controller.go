package controller

import (
	"carlife-backend/gateway"
	"carlife-backend/model"
	"carlife-backend/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	userID := ctx.PostForm("userID")
	userType := ctx.PostForm("userType")
	password := ctx.PostForm("password")
	hashPwd, err := utils.HashPassword(password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "register failed,can't hash password:" + err.Error()})
		return
	}
	res, err := gateway.RegisterUser(userID, userType, hashPwd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "register failed:" + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"txid": res,
		"msg":  "register success",
	})
}

func Login(ctx *gin.Context) {
	//get user
	user := model.User{}
	userID := ctx.PostForm("userID")
	password := ctx.PostForm("password")
	userJson, err := gateway.GetUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unauthorized:" + err.Error()})
		return
	}
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unauthorized:" + err.Error()})
		return
	}

	//check password
	if !utils.CheckPassword(password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: password error"})
		return
	}

	//create JWT
	token, err := utils.GenerateJWT(user.UserID, user.UserType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "create jwt failed:" + err.Error()})
		return
	}

	//return
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"msg":   "login success",
	})
}
