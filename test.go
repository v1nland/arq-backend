package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/dgrijalva/jwt-go"
)


func jwttest(c *gin.Context) {
    // Create token
    token := jwt.New(jwt.SigningMethodHS256)
    // Set claims
    // This is the information which frontend can use
    // The backend can also decode the token and get admin etc.
    claims := token.Claims.(jwt.MapClaims)
    claims["user"] = "Jonsito"
    // Generate encoded token and send it as response.
    // The signing string should be secret (a generated UUID          works too)
    t, err := token.SignedString( []byte("jt65he4ae5ae") )

    c.JSON(200, gin.H{
        "token": t,
        "error": err,
    })
}

func HomePage(c *gin.Context){
    // Set an standard HomePage message
    c.JSON(200, gin.H{
        "message": "Hello World",
        "test2": []byte("jt65he4ae5ae"),
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
