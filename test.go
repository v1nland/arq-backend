package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func HomePage(c *gin.Context){
    // Set an standard HomePage message
    c.JSON(200, gin.H{
        "message": "Hello World",
    })
}

func TestGoEngine(){
    fmt.Println("Package working right!")
}

func TestUrlParameters(c *gin.Context){
    var id = c.Param("id")
    var hola = c.Param("hola")

    c.JSON(200, gin.H{
		"ID": id,
        "Hola": hola,
	})
}

func TestDatabaseConnection(c *gin.Context){
    // Try DB connection and show its results
    connState := ""
    db, err := sql.Open("mysql", "b5ef0f5f54c166:fc7bd410@tcp(us-cdbr-iron-east-02.cleardb.net:3306)/heroku_eb1b634cb7a7bd8")

    // si es que existe error
    if err != nil {
        connState = "Not connected"
    }else{
        connState = "Succefully connected"
    }

    c.JSON(200, gin.H{
        "message": connState,
    })

    defer db.Close()
}
