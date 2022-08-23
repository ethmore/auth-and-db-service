package middleware

import (
	"e-comm/authService/dotEnv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt"

	"fmt"
	"log"
	"net/http"
)

type TokenBody struct {
	Token string
}

func UserAuth(c *gin.Context) (x, y string) {
	defer func() {
		if recover() != nil {
			log.Println("User not logged in")
			c.JSON(http.StatusOK, gin.H{"message": "loginNeeded"})
		}
	}()

	var tokenBody TokenBody
	if err := c.ShouldBindBodyWith(&tokenBody, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}

	clientToken := tokenBody.Token
	secretToken := dotEnv.GoDotEnvVariable("TOKEN")
	hmacSampleSecret := []byte(secretToken)

	token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["mail"].(string), claims["type"].(string)

	} else {
		fmt.Println(err)
		log.Println("User not logged in")
		c.JSON(http.StatusOK, gin.H{"message": "loginNeeded"})
		c.Abort()

		return "", ""
	}
}
