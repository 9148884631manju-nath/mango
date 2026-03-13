package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"mango/lib"
)

func main(){
	r :=gin.Default()
	//var proj string
	r.GET("/",func(c *gin.Context){
		//proj = c.Query("proj")
		
		var data string
		var result []byte
		data=""
		
		//Read Headers scripts or styles			
			data+=lib.ReadData("src/header_scripts.txt")
		//End Reading Headers 

		// Welcome Greets for Man-Go Framework
		data+= lib.ToHtml("home.json")
		//data += `<h1>Man-GO v 0.1</h1><h3>GO Lang Web Framework V 0.1</h3>`

		//Read Footer scripts or styles
		data+=lib.ReadData("src/footer_scripts.txt")
		//End Reading Footer 

		result = []byte(data)
		c.Data(http.StatusOK, "text/html", result)
	})

	r.Run(":8000")
}