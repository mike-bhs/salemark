package routes

import (
	u "github.com/salemark/utils"
	"net/http"
)

func rootHandler(res http.ResponseWriter, req *http.Request) {
	body := "<h1>Welcome!</h1>"
	u.HtmlResponse(res, 200, body)
}
