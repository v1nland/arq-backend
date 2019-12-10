package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// Fetch all departamentos ---------------------------------------------------------------------------

func (de Departamento) FetchDepartamentos() (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT * FROM departamentos")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente)
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



//SELECT PARA LOGIN ------------------------------------------------------------------------------

func (de Departamento) FetchDptoLogin(codigo_condominio string, numero_dpto string, password_dpto string) (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from departamentos where (id_condominio = (select id from condominios where codigo = ?)) and numero= ? and password = ?", codigo_condominio, numero_dpto, password_dpto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente)
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


//SELECT DEPARTAMENTOS POR ID_CONDOMINIO-------------------------------------------------------------

func (de Departamento) FetchDptoCondominio(id_condominio string) (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from departamentos where id_condominio = ?", id_condominio)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetDptoCondominio(c *gin.Context){
    // URL parameters
    var idcondominio = c.Param("idcondominio")

    fmt.Println(idcondominio);

    // Results container
    de := Departamento{}
    // Fetch from database
	departamentos, err := de.FetchDptoCondominio(idcondominio)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}

//SELECT DEPARTAMENTOS POR ID--------------------------------------------------------------------

func (de Departamento) FetchDptoID(id_departamento string) (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from departamentos where id= ?", id_departamento)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetDptoID(c *gin.Context){
    // URL parameters
    var iddpto = c.Param("iddpto")

    fmt.Println(iddpto);

    // Results container
    de := Departamento{}
    // Fetch from database
	departamentos, err := de.FetchDptoID(iddpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}


//SELECT TODOS LOS DATOS CON ID---------------------------------------------------------------------

func (de Departamento) FetchTodoDptoID(id_departamento string) (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from departamentos JOIN tickets on departamentos.id = tickets.id_departamentos JOIN bodegas on tickets.id_departamentos = bodegas.id_departamentos JOIN estacionamientos on bodegas.id_departamentos = estacionamientos.id_departamentos JOIN gastos_comunes on estacionamientos.id_departamentos = gastos_comunes.id_departamentos JOIN mediciones_agua on gastos_comunes.id_departamentos = mediciones_agua.id_departamentos JOIN multas on mediciones_agua.id_departamentos = multas.id_departamentos JOIN pagos_gastos_comunes on multas.id_departamentos = pagos_gastos_comunes.id_departamentos where departamentos.id= ?", id_departamento)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
		//var tick Ticket
		//var bode Bodega
		//var esta Estacionamiento
		//var gast Gasto_comun
		//var medi Medicion_agua
		//var mult Multa
		//var pago Pago_gastos_comunes
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetTodoDptoID(c *gin.Context){
    // URL parameters
    var iddpto = c.Param("iddpto")

    fmt.Println(iddpto);

    // Results container
    de := Departamento{}
    // Fetch from database
	departamentos, err := de.FetchTodoDptoID(iddpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}
