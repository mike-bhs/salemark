package routes

import (
	search "github.com/salemark/search"
	u "github.com/salemark/utils"
	"net/http"
)

func SearchHandler(res http.ResponseWriter, req *http.Request, s search.Search) {
	sType := u.GetParamValue(req, "type")
	body := u.GetParamValue(req, "body")

	if sType == "" || body == "" {
		u.JsonResponse(res, 400, "Bad Data")
		return
	}

	result := search.Find(s, sType, body)
	u.JsonResponse(res, 200, result)
}
