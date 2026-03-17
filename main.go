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
		proj := c.Query("proj")
		if proj=="" {
			proj="home"
		}else {
			proj=proj
		}
		var data string
		var result []byte
		data=""
		
		//Read Headers scripts or styles			
			data+=lib.ReadData(proj+"/header_scripts.txt")
		//End Reading Headers 

		// Welcome Greets for Man-Go Framework
		
		data+= lib.ToHtml(proj+".json", c)
		//data += `<h1>Man-GO v 0.1</h1><h3>GO Lang Web Framework V 0.1</h3>`

		//Read Footer scripts or styles
		data+=lib.ReadData(proj+"/footer_scripts.txt")
		//End Reading Footer 

		result = []byte(data)
		c.Data(http.StatusOK, "text/html", result)
	})

	r.Run(":8000")
}