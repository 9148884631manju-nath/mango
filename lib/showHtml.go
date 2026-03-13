package lib

func ShowHTML(k string, v interface{}) string {
	var jv string
	var res string
	switch k {
	case "content":
		s := v.(string)
		jv = s
	case "block":
		s := v.(string)
		jv = ToHtml(s)
	case "inblock":
		jv = JsonArrayRenderHtml(v)
	default:
		jv = ""
	}
	res=jv
	return res
}
