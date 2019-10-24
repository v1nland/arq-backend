package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// Fetch all condominios
func (co Condominio) FetchCondominios() (condominios []Condominio, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, codigo, nombre, ubicacion FROM condominios")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var cond Condominio
        rows.Scan(&cond.Id, &cond.Codigo, &cond.Nombre, &cond.Ubicacion)
        fmt.Println(cond.Id, cond.Codigo, cond.Nombre, cond.Ubicacion)
        condominios = append(condominios, cond)
	}
	defer rows.Close()

	return
}

func GetCondominios(c *gin.Context){
    // Results container
    co := Condominio{}
    // Fetch from database
	condominios, err := co.FetchCondominios()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": condominios,
		"count": len(condominios),
	})
}

// Fetch all condominios where ID = condominio_id
func (co Condominio) FetchCondominiosPorID(condominio_id string) (condominios []Condominio, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, codigo, nombre, ubicacion FROM condominios WHERE id = ?", condominio_id)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var cond Condominio
        rows.Scan(&cond.Id, &cond.Codigo, &cond.Nombre, &cond.Ubicacion)
        fmt.Println(cond.Id, cond.Codigo, cond.Nombre, cond.Ubicacion)
        condominios = append(condominios, cond)
	}
	defer rows.Close()

	return
}

func GetCondominiosPorID(c *gin.Context){
    // Get URL parameters
    var id = c.Param("id")

    // Results container
    co := Condominio{}

    // Fetch from database
	condominios, err := co.FetchCondominiosPorID(id)
	if err != nil {
        panic(err.Error())
	}

    // Render via GET method
	c.JSON(200, gin.H{
		"rows": condominios,
		"count": len(condominios),
	})
}
