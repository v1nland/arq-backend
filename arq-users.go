package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func (usr Usuario) FetchUsuarios() (users []Usuario, err error) {
    // Opens DB
    db = GetConnection()
    defer db.Close()

    // SQL query
	rows, err := db.Query("SELECT id, rut, password FROM usuarios")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var usr Usuario
        rows.Scan(&usr.Id, &usr.Rut, &usr.Password)
        fmt.Println(usr.Id, usr.Rut, usr.Password)
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

func (usr Usuario) FetchUsuariosPorID(usuario_rut string) (users []Usuario, err error) {
    // Opens DB
    db = GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, rut, password FROM usuarios WHERE rut = ?", usuario_rut)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var usr Usuario
        rows.Scan(&usr.Id, &usr.Rut, &usr.Password)
        fmt.Println(usr.Id, usr.Rut, usr.Password)
        users = append(users, usr)
	}
	defer rows.Close()

	return
}

func GetUsuariosPorID(c *gin.Context){
    // Get URL parameters
    var rut = c.Param("rut")

    // Results container
    usr := Usuario{}
    // Fetch from database
	users, err := usr.FetchUsuariosPorID(rut)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": users,
		"count": len(users),
	})
}

func (usr Usuario) FetchLogin(usuario_rut string, usuario_pass string) (users []Usuario, err error) {
    // Opens DB
    db = GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, rut, password FROM usuarios WHERE rut = ? AND password = ?", usuario_rut, usuario_pass)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var usr Usuario
        rows.Scan(&usr.Id, &usr.Rut, &usr.Password)
        fmt.Println(usr.Id, usr.Rut, usr.Password)
        users = append(users, usr)
	}
	defer rows.Close()

	return
}

func GetLogin(c *gin.Context){
    // Get URL parameters
    var rut = c.Param("rut")
    var password = c.Param("password")

    // Results container
    usr := Usuario{}
    // Fetch from database
	users, err := usr.FetchLogin(rut, password)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": users,
		"count": len(users),
	})
}
