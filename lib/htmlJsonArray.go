package lib

import (
	"github.com/gin-gonic/gin"
)

func HtmlJsonArray(jsd []map[string]interface{}, c *gin.Context) string {
res:=""
for i, rec := range jsd {
		var jv string
		var st string
		var et string

		//ELEMENT ATTRIBUTE
		hatr := ""
		if htmlattr, ok := rec["attr"]; ok {
			a := htmlattr.(map[string]interface{})
			hatr = HtmlAttr(a)
		}else {
			hatr=""
		}
		//END ELEMENT ATTRIBUTE

		//CLASS ATTRIBUTE
		if classval, ok := rec["class"]; ok {
			s := classval.(string)
			st = `<div class="` + s + `" ` + hatr + `>`
			et = `</div>`
		}else {
			st=""
			et=""
		}
		//END CLASS ATTRIBUTE
		
		for k, v := range rec {
			jv += ShowHTML(k, v, rec["keys"], c)
		}
		
		res += st + jv + et
		i += 1
}
//res+=fmt.Sprintf("%v",jsd)
return res
}