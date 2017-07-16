package services

import (
	"net/http"
	"regexp"
	"strings"

	routes "github.com/salemark/routes"
	utils "github.com/salemark/utils"
)

type Request struct {
	Source *http.Request
}

func (req Request) MatchRoute(route routes.Route) bool {
	return req.Method() == route.Method && req.Path() == route.Path
}

func (req Request) MatchPatter(pattern string) bool {
	matched, err := regexp.Match(pattern, []byte(req.Path()))

	if utils.HandleError(err) {
		return false
	}

	return matched
}

func (req Request) GetParamValue(name string) (string, error) {
	err := req.ParseForm()

	if err != nil {
		return "", err
	}

	value := req.Source.Form[name]
	return strings.Join(value[:], ", "), nil
}

func (req Request) Path() string {
	return req.Source.URL.Path
}

func (req Request) Method() string {
	return req.Source.Method
}

func (req Request) Url() string {
	return req.Source.URL.String()
}

func (req Request) SimpleParams() (map[string]string, error) {
	err := req.ParseForm()

	if err != nil {
		return nil, err
	}

	m := make(map[string]string)

	for key, values := range req.Source.Form {
		v := strings.Join(values[:], ", ")
		m[key] = v
	}

	return m, nil
}

func (req Request) ParseForm() error {
	err := req.Source.ParseForm()

	if utils.HandleError(err) {
		return err
	}

	return nil
}
