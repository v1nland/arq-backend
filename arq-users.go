package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

// Fetch all users
func (usr Usuario) FetchUsuarios() (users []Usuario, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, rut, password, nombre FROM usuarios")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var usr Usuario
        rows.Scan(&usr.Id, &usr.Rut, &usr.Password, &usr.Nombre)
        users = append(users, usr)
	}
	defer rows.Close()

	return
}

func GetUsuarios(c *gin.Context){
    // Results container
    usr := Usuario{}
    // Fetch from database
	users, err := usr.FetchUsuarios()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": users,
		"count": len(users),
	})
}

func (usr Usuario) FetchUserLogin(usuario_rut string, usuario_pass string) (users []Usuario, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, rut, password, nombre FROM usuarios WHERE rut = ? AND password = ?", usuario_rut, usuario_pass)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var usr Usuario

        rows.Scan(&usr.Id, &usr.Rut, &usr.Password, &usr.Nombre)

        token := jwt.New(jwt.SigningMethodHS256)
        claims := token.Claims.(jwt.MapClaims)
        claims["id"] = usr.Id
        claims["rut"] = usr.Rut
        claims["password"] = usr.Password
	claims["nombre"] = usr.Nombre
        claims["level"] = "admin"
        usr_token, err := token.SignedString( SecretKey() )
        usr.Token = usr_token
        usr.Level = "admin"

        _ = err

        fmt.Println(usr.Id, usr.Rut, usr.Password, usr.Level, usr.Token)
        users = append(users, usr)
	}
	defer rows.Close()

	return
}

func GetUserLogin(c *gin.Context){
    // Get URL parameters
    var rut = c.Param("rut")
    var password = c.Param("password")

    usr := Usuario{}

    // Fetch from database
	users, err := usr.FetchUserLogin(rut, password)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": users,
		"count": len(users),
	})
}
