package main
import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

func SecretKey() ([]byte){
    return []byte("jt65he4ae5ae")
}

func VerifyToken(c *gin.Context){
    var usr_token = c.Param("token")

    claims := jwt.MapClaims{}

    token, err := jwt.ParseWithClaims(usr_token, claims, func(token *jwt.Token) (interface{}, error) {
        return SecretKey(), nil
    })

    _ = err

    if token.Valid {
        c.JSON(200, gin.H{
    		"valid": "true",
            "id": claims["id"],
            "rut": claims["rut"],
            "password": claims["password"],
            "level": claims["level"],
    	})
	} else {
        c.JSON(200, gin.H{
            "valid": "false",
        })
	}
}
