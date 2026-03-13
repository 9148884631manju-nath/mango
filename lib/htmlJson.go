package lib

func HtmlJson(data interface{}) string {
	res := ""
	res = RenderHTML(data.(string))
	return res
}
