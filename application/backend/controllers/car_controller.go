package controller

import (
	"carlife-backend/gateway"
	"carlife-backend/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCar(ctx *gin.Context) {
	fmt.Printf("ctx.Query(\"carID\"): %v\n", ctx.Query("carID"))
	res, err := gateway.GetCar(ctx.Query("carID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "query success",
		"data": res,
	})
}

func GetCarList(ctx *gin.Context) {
	var user = model.User{}
	userID, _ := ctx.Get("userID")
	res, err := gateway.GetUser(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	carlist := user.CarList
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "query success",
		"data": carlist,
	})
}

func SetCarTires(ctx *gin.Context) {
	var input struct {
		CarID    string  `form:"carID"`
		Width    float32 `form:"width"`
		Radius   float32 `form:"radius"`
		Workshop string  `form:"workshop"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := ctx.Get("userID")
	res, err := gateway.SetCarTires(
		userID.(string),
		input.CarID,
		input.Width,
		input.Radius,
		input.Workshop,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "set failed:" + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "set success",
		"txid":  res,
		"carID": input.CarID,
	})
}

func SetCarBody(ctx *gin.Context) {
	var input struct {
		CarID    string  `form:"carID"`
		Material string  `form:"material"`
		Weitght  float32 `form:"weitght"`
		Color    string  `form:"color"`
		Workshop string  `form:"workshop"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := ctx.Get("userID")
	res, err := gateway.SetCarBody(
		userID.(string),
		input.CarID,
		input.Material,
		input.Weitght,
		input.Color,
		input.Workshop,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "set failed:" + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "set success",
		"txid":  res,
		"carID": input.CarID,
	})
}

func SetCarInterior(ctx *gin.Context) {
	var input struct {
		CarID    string  `form:"carID"`
		Material string  `form:"material"`
		Weitght  float32 `form:"weitght"`
		Color    string  `form:"color"`
		Workshop string  `form:"workshop"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := ctx.Get("userID")
	res, err := gateway.SetCarInterior(
		userID.(string),
		input.CarID,
		input.Material,
		input.Weitght,
		input.Color,
		input.Workshop,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "set failed:" + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "set success",
		"txid":  res,
		"carID": input.CarID,
	})
}

func SetCarManu(ctx *gin.Context) {
	var input struct {
		CarID    string `form:"carID"`
		Workshop string `form:"workshop"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := ctx.Get("userID")
	res, err := gateway.SetCarManu(
		userID.(string),
		input.CarID,
		input.Workshop,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "set failed:" + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "set success",
		"txid":  res,
		"carID": input.CarID,
	})
}

func SetCarStore(ctx *gin.Context) {
	var input struct {
		CarID   string  `form:"carID"`
		Store   string  `form:"store"`
		Cost    float32 `form:"cost"`
		OwnerID string  `form:"ownerID"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := ctx.Get("userID")
	res, err := gateway.SetCarStore(
		userID.(string),
		input.CarID,
		input.Store,
		input.Cost,
		input.OwnerID,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "set failed:" + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "set success",
		"txid":  res,
		"carID": input.CarID,
	})
}

func SetCarInsure(ctx *gin.Context) {
	var input struct {
		CarID string  `form:"carID"`
		Name  string  `form:"name"`
		Cost  float32 `form:"cost"`
		Years int     `form:"years"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := ctx.Get("userID")
	res, err := gateway.SetCarInsure(
		userID.(string),
		input.CarID,
		input.Name,
		input.Cost,
		input.Years,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "set failed:" + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "set success",
		"txid":  res,
		"carID": input.CarID,
	})
}

func SetCarMaint(ctx *gin.Context) {
	var input struct {
		CarID  string  `form:"carID"`
		Part   string  `form:"part"`
		Extent string  `form:"extent"`
		Cost   float32 `form:"cost"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := ctx.Get("userID")
	res, err := gateway.SetCarMaint(
		userID.(string),
		input.CarID,
		input.Part,
		input.Extent,
		input.Cost,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "set failed:" + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "set success",
		"txid":  res,
		"carID": input.CarID,
	})
}

func TransferCar(ctx *gin.Context) {
	var input struct {
		CarID      string  `form:"carID"`
		NewOwnerID string  `form:"newOwnerID"`
		Cost       float32 `form:"cost"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := ctx.Get("userID")
	res, err := gateway.TransferCar(
		userID.(string),
		input.CarID,
		input.NewOwnerID,
		input.Cost,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "transfer failed:" + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "transfer success",
		"txid":  res,
		"carID": input.CarID,
	})
}
