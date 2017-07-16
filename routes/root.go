package routes

import (
	s "github.com/salemark/services"
)

func rootHandler(res s.Response, req s.Request) {
	body := "<h1>Welcome!</h1>"
	res.HtmlResponse(200, body)
}
