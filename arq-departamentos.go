package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
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
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente, &dpto.Prorrateo)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente, dpto.Prorrateo)
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

func (de Departamento) FetchDptoLogin(cond_code string, dpto_num string, dpto_pass string) (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from departamentos where (id_condominio = (select id from condominios where codigo = ?)) and numero= ? and password = ?", cond_code, dpto_num, dpto_pass)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente, &dpto.Prorrateo)

        token := jwt.New(jwt.SigningMethodHS256)
        claims := token.Claims.(jwt.MapClaims)
        claims["id"] = dpto.Id
        claims["numero"] = dpto.Numero
        claims["password"] = dpto.Password
        claims["dueno"] = dpto.Dueno
        claims["residente"] = dpto.Residente
        claims["telefono"] = dpto.Telefono
        claims["correo"] = dpto.Correo
        claims["id_cond"] = dpto.Id_condominio
        claims["telefono_residente"] = dpto.Telefono_residente
        claims["correo_residente"] = dpto.Correo_residente
	claims["prorrateo"] = dpto.Prorrateo
        claims["level"] = "user"
        dpto_token, err := token.SignedString( SecretKey() )
        dpto.Token = dpto_token
        dpto.Level = "user"

        _ = err

        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente, dpto.Prorrateo, dpto.Token)
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
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente, &dpto.Prorrateo)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente, dpto.Prorrateo)
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
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente, &dpto.Prorrateo)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente, dpto.Prorrateo)
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

//get bodega y estacionamiento----------------------------------------------------------------------

func (de DepartamentoBodEst) FetchDepartamentosEstBod(id_dpto string) (departamentos []DepartamentoBodEst, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT departamentos.id, departamentos.numero, departamentos.password, departamentos.dueno, departamentos.residente, departamentos.telefono, departamentos.correo, departamentos.id_condominio, departamentos.telefono_residente, departamentos.correo_residente, departamentos.prorrateo, estacionamientos.numero as n_estacionamientos, bodegas.numero as n_bodegas FROM departamentos join estacionamientos on departamentos.id = estacionamientos.id_departamentos join bodegas on estacionamientos.id_departamentos = bodegas.id_departamentos where departamentos.id = ?", id_dpto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto DepartamentoBodEst
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente, &dpto.Prorrateo, &dpto.N_bodega, &dpto.N_estacionamiento)
	dpto.Level = "user"
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente, dpto.Prorrateo, dpto.N_bodega, dpto.N_estacionamiento)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetDepartamentosEstBod(c *gin.Context){
    // URL parameters
    var iddpto = c.Param("iddpto")

    fmt.Println(iddpto);
    // Results container
    de := DepartamentoBodEst{}
    // Fetch from database
	departamentos, err := de.FetchDepartamentosEstBod(iddpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}

