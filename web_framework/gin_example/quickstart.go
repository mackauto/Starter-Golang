package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// create gin router with default middleware: logger and recovery
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//router.GET("/get", getting)
	//router.POST("/post", posting)
	//router.PUT("/put", putting)
	//router.DELETE("/delete", deleting)
	//router.PATCH("/patch", patching)
	//router.HEAD("/head", heading)
	//router.OPTIONS("/option", optioning)

	// this handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})

	// will mach /user/john/ and also /user/john/send
	// if no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK, message)
	})

	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname")
		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	port := ":5000"
	_ = router.Run(port)
	//_ = router.Run()
}
