package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)
func AuthorizeJWT(c *gin.Context) interface{} {
	tokenHeader := c.GetHeader("Authorization")
	if tokenHeader == "" {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusForbidden, gin.H{
			"error":   true,
			"message": "missing token",
		})
		return ""
	}
	//The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Invalid/Malformed auth token",
		})
		c.Header("Content-Type", "application/json")
		return ""
	}
	//Grab the token part
	tokenPart := splitted[1]
	token, err := jwt.Parse(tokenPart, func(token *jwt.Token) (interface{}, error) {
		return []byte("key"), nil
	})
claims,_:=token.Claims.(jwt.MapClaims)
	if err != nil { //Malformed token, returns with http code 403 as usual
		c.JSON(403, gin.H{
			"error":   true,
			"message": "Malformed authentication token",
		})
		c.Header("Content-Type", "application/json")
		return ""
	}

	if !token.Valid { //Token is invalid, maybe not signed on this server
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Malformed authentication token",
		})
		c.Header("Content-Type", "application/json")
		return ""
	}
	//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
	//fmt.Sprintf("User %", tk.input) //Useful for monitoring
	c.Set("user", claims["user_id"])
	c.Next()
	return claims["user_id"]
}