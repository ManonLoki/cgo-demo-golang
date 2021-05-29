package main


/*

 typedef  struct User {
   char* Name;
   int Age;
  } User;
 */
import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)




//export SayHello
func SayHello(name *C.char){
	fmt.Println("Hello ",C.GoString(name))
}

//export CreateUser
func CreateUser(name *C.char) C.User{
	 user:= C.User{
	 	Name:name,
	 	Age:C.int( rand.Int()),
	 }
	 return user
}

//export DisplayUser
func DisplayUser(user C.User){

	res,err:=json.Marshal(user)
	if err!=nil{
		fmt.Println("Json Encode Error...")
		return
	}

	fmt.Println("Json Encoded Result :")
	fmt.Println(string(res))
}

//export StartWebServer
func StartWebServer(){

		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		fmt.Println("Gin Web Server Start Listen 0.0.0.0:3000")
		r.Run(":3000")
}

// 必须定义 没有办法省略
func main(){}