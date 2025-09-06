package controllers

import (
	"net/http"

	"example.com/blog/global"
	"example.com/blog/models"
	"example.com/blog/utils"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hasdPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = hasdPwd
	token, err := utils.GenerateJwt(user.Username, int(user.ID), user.TokenVersion)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := global.Db.Debug().Table("users").Where("username = ?", input.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "username or password uncorrect"})
		return
	}
	if !utils.CheckPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "username or password uncorrect"})
		return
	}
	// 递增令牌版本号 - 使所有旧令牌失效
	if err := global.Db.Model(&user).Update("token_version", user.TokenVersion+1).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "登录失败"})
		return
	}
	//生成新令牌
	token, err := utils.GenerateJwt(user.Username, int(user.ID), user.TokenVersion)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})

}
