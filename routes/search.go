package routes

// import (
// 	search "github.com/salemark/search"
// 	"encoding/json"
// 	"net/http"
// 	"strings"
//
// u "github.com/salemark/utils"
// )

type SearchRequest struct {
	Index  string
	Type   string
	Params string
}

// url to find phones must look like
// /search/products?color=black&brand=apple&type=phone
// func SearchHandler(res http.ResponseWriter, req *http.Request) {
// 	sIndex, sType := processUrl(req)
//
// 	searchRequestJson, _ := json.Marshal(processUrl(req))
// 	log.Println("SEARCH PARAMS")
// 	log.Println(string(searchRequestJson[:]))
//
// 	result := search.Find(sType, body)
// 	u.JsonResponse(res, 200, "")
// }

// func processUrl(req *http.Request) SearchRequest {
// 	elements := strings.Split(req.URL.Path, "/")
//
// 	sIndex, sType := "", ""
//
// 	if len(elements) > 3 {
// 		sIndex, sType = elements[2], elements[3]
// 	}
//
// 	if len(elements) > 2 {
// 		sIndex = elements[2]
// 	}
//
// 	jsonParams, _ := json.Marshal(u.RequestParams(req))
//
// 	return SearchRequest{
// 		Index:  sIndex,
// 		Type:   sType,
// 		Params: string(jsonParams[:]),
// 	}
// }
