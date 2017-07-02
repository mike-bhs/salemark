package routes

import (
	search "github.com/salemark/search"
	u "github.com/salemark/utils"
	"net/http"
)

// url to find phones must look like
// /search/products?color=black&brand=apple&type=phone
func SearchHandler(res http.ResponseWriter, req *http.Request) {
	sType := u.GetParamValue(req, "type")
	body := u.GetParamValue(req, "body")

	if sType == "" || body == "" {
		u.JsonResponse(res, 200, "Nothing found")
		return
	}

	result := search.Find(sType, body)
	u.JsonResponse(res, 200, result)
}
