package controller

import (
	"carlife-backend/gateway"
	"carlife-chaincode-go/chaincode"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCar(ctx *gin.Context) {
	res, err := gateway.GetCar(ctx.PostForm("carID"))
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
	var user chaincode.User
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
		CarID    string  `json:"carID"`
		Width    float32 `json:"width"`
		Radius   float32 `json:"radius"`
		Workshop string  `json:"workshop"`
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
		CarID    string  `json:"carID"`
		Material string  `json:"material"`
		Weitght  float32 `json:"weitght"`
		Color    string  `json:"color"`
		Workshop string  `json:"workshop"`
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
		CarID    string  `json:"carID"`
		Material string  `json:"material"`
		Weitght  float32 `json:"weitght"`
		Color    string  `json:"color"`
		Workshop string  `json:"workshop"`
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
		CarID    string `json:"carID"`
		Workshop string `json:"workshop"`
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
		CarID   string  `json:"carID"`
		Store   string  `json:"store"`
		Cost    float32 `json:"cost"`
		OwnerID string  `json:"ownerID"`
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
		CarID string  `json:"carID"`
		Name  string  `json:"name"`
		Cost  float32 `json:"cost"`
		Years int     `json:"years"`
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
		CarID  string  `json:"carID"`
		Part   string  `json:"part"`
		Entent string  `json:"entent"`
		Cost   float32 `json:"cost"`
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
		input.Entent,
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
		CarID      string  `json:"carID"`
		NewOwnerID string  `json:"newOwnerID"`
		Cost       float32 `json:"cost"`
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
