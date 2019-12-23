package main
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetConnection() *sql.DB {
    // Para evitar realizar una nueva conexión en cada llamada a
    // la función GetConnection.
    if db != nil {
        return db
    }

    // Declaramos la variable err para poder usar el operador
    // de asignación “=” en lugar que el de asignación corta,
    // para evitar que cree una nueva variable db en este scope y
    // en su lugar que inicialice la variable db que declaramos a
    // nivel de paquete.
    var err error
    // Conexión a la base de datos
    db, err = sql.Open("mysql", "b5ef0f5f54c166:fc7bd410@tcp(us-cdbr-iron-east-02.cleardb.net:3306)/heroku_eb1b634cb7a7bd8")
    if err != nil {
        panic(err)
    }

    db.SetMaxIdleConns(0)

    return db
}
