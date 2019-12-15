package main
import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

func SecretKey() ([]byte){
    return []byte("jt65he4ae5ae")
}

func DecodeToken(c *gin.Context){
    var usr_token = c.Param("token")

    claims := jwt.MapClaims{}
    token, err := jwt.ParseWithClaims(usr_token, claims, func(token *jwt.Token) (interface{}, error) {
        return SecretKey(), nil
    })

    if err != nil {
        return
    }

    if token.Valid && claims["level"] == "admin" {
        c.JSON(200, gin.H{
    		"valid": "true",
            "id": claims["id"],
            "rut": claims["rut"],
            "password": claims["password"],
            "level": claims["level"],
    	})
	} else if token.Valid && claims["level"] == "user" {
        c.JSON(200, gin.H{
    		"valid": "true",
            "numero": claims["numero"],
            "dueno": claims["dueno"],
            "residente": claims["residente"],
            "telefono": claims["telefono"],
            "correo": claims["correo"],
            "id_cond": claims["id_cond"],
            "telefono_residente": claims["telefono_residente"],
            "correo_residente": claims["correo_residente"],
            "level": claims["level"],
    	})
    }else {
        c.JSON(200, gin.H{
            "valid": "false",
        })
	}
}
