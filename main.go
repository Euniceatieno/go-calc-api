package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//////////////////////////////////////////////////////////////
// set up dependency tracking 
// CMD - go mod init example/go-calc-api - creates a go.mod file
//////////////////////////////////////////////////////////////


////////////////////////////////////////////////////////////////////
//  Dowload gin
//  CMD - go get github.com/gin-gonic/gin -
//  Gin is a high-performance HTTP/Restful web framework written in Golang (Go).
//  Gin has a martini-like API and claims to be up to 40 times faster.
//  Gin allows you to build web applications and microservices in Go.
//  It contains a set of commonly used functionalities 
//  (e.g., routing, middleware support, rendering, etc.)
//////////////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////////////
// Define calc data sructure using a struct
type calc struct{
	Number1  int  `json:"number1"` 
	Number2  int  `json:"number2"`
}




/////////////////////////////////////////////////////////////////
// create the calc server
func main(){
	router := gin.Default() // variable router reps the server
	router.POST("/add")
	router.POST("/subtract")
	router.POST("/multiply")
	router.POST("/divide")
	router.Run("localhost:9000") //running the server at port 9000

}

/////////////////////////////////////////////////////////////////
// RUN APP CMD - go run main.go
/////////////////////////////////////////////////////////////////

