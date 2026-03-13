package lib

import (
	"encoding/json"
)

func RenderHTML(data string) string {
	res := ""
	var jsd []map[string]interface{}
	json.Unmarshal([]byte(data), &jsd)
	res=HtmlJsonArray(jsd)
	return res
}
