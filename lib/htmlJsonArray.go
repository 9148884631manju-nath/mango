package lib

import (

)

func HtmlJsonArray(jsd []map[string]interface{}) string {
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
		}
		//END ELEMENT ATTRIBUTE

		//CLASS ATTRIBUTE
		if classval, ok := rec["class"]; ok {
			s := classval.(string)
			st = `<div class="` + s + `" ` + hatr + `>`
			et = `</div>`
		}
		//END CLASS ATTRIBUTE

		for k, v := range rec {
			jv += ShowHTML(k, v)
		}

		res += st + jv + et
		i += 1
	}
return res
}