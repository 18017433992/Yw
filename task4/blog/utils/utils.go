package utils

import (
	"errors"
	"fmt"
	"time"

	"example.com/blog/config"
	"example.com/blog/global"
	"example.com/blog/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}
func GenerateJwt(username string, id int, tokenversion uint) (string, error) {
	exptime := config.AppConfig.JWT.ExpireHours
	secret := config.AppConfig.JWT.SecretKey
	fmt.Println("secret", secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":           id,
		"username":     username,
		"exp":          time.Now().Add(time.Hour * time.Duration(exptime)).Unix(), // //time.Now().Add(time.Hour * 24).Unix(),
		"tokenversion": tokenversion,
	})
	signedToken, err := token.SignedString([]byte(secret))
	fmt.Println()
	return "Bearer " + signedToken, err
}
func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func ParseJWT(tokenString string) (*models.User, error) {

	secret := config.AppConfig.JWT.SecretKey
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return nil, errors.New("username claims is not Valid ")
		}
		// 提取用户ID和令牌版本
		userID := uint(claims["id"].(float64))
		tokenVersion := uint(claims["tokenversion"].(float64))
		// 从数据库获取用户
		var user models.User
		if err := global.Db.First(&user, userID).Error; err != nil {
			return nil, fmt.Errorf("user not found: %v", err)
		}

		// 验证令牌版本
		if user.TokenVersion != tokenVersion {
			return nil, errors.New("token has been invalidated - please login again")
		}
		user.Username = username
		user.ID = userID
		return &user, nil
	}
	return nil, err
}
