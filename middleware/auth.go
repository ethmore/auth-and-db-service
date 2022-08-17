package middleware

import (
	"e-comm/authService/dotEnv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt"

	"fmt"
	"log"
)

type TokenBody struct {
	Token string
}

func UserAuth(c *gin.Context) string {
	defer func() {
		if recover() != nil {
			log.Println("User not logged in")
			c.JSON(200, gin.H{"message": "loginNeeded"})
		}
	}()

	var tokenBody TokenBody
	// if err := c.BindJSON(&tokenBody); err != nil {
	// 	fmt.Println(err)
	// }

	if err := c.ShouldBindBodyWith(&tokenBody, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}

	clientToken := tokenBody.Token
	fmt.Println(tokenBody.Token)

	secretToken := dotEnv.GoDotEnvVariable("TOKEN")
	hmacSampleSecret := []byte(secretToken)

	token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		fmt.Println(claims["mail"], claims["nbf"])
		// c.Next()
		return claims["mail"].(string)

	} else {
		fmt.Println(err)
		log.Println("User not logged in")
		c.JSON(200, gin.H{"message": "loginNeeded"})
		// c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()

		return ""
	}
}

/*
func AuthRequired(c *gin.Context) {
	var tokenBody TokenBody
	if err := c.BindJSON(&tokenBody); err != nil {
		fmt.Println(err)
	}

	clientToken := tokenBody.Token

	secretToken := dotEnv.GoDotEnvVariable("TOKEN")
	hmacSampleSecret := []byte(secretToken)

	token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["mail"], claims["nbf"])
		c.Next()

	} else {
		fmt.Println(err)
		log.Println("User not logged in")
		// c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
		return
	}
}
*/
