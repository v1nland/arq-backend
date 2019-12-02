package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// Fetch all departamentos
func (de Departamento) FetchDepartamentos() (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, numero, password, dueno, residente, telefono, correo, cod_condominio FROM departamentos")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Cod_condominio)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Cod_condominio)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetDepartamentos(c *gin.Context){
    // Results container
    de := Departamento{}
    // Fetch from database
	departamentos, err := de.FetchDepartamentos()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}

func (de Departamento) FetchDptoLogin(cond_code string, dpto_num string, dpto_pass string) (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT * FROM departamentos WHERE numero = ? AND password = ? AND id_condominio = ?", dpto_num, dpto_pass, cond_code)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Cod_condominio)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Cod_condominio)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetDptoLogin(c *gin.Context){
    // URL parameters
    var codigo = c.Param("codigo")
    var numero = c.Param("numero")
    var password = c.Param("password")

    fmt.Println(codigo);
    fmt.Println(numero);
    fmt.Println(password);

    // Results container
    de := Departamento{}
    // Fetch from database
	departamentos, err := de.FetchDptoLogin(codigo, numero, password)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}
