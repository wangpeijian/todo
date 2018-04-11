package core

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/url"
)

type router struct {
	routerMap map[string]func(requireHandle RequireHandle) (result Result)
}

/**
 *	注册controller
 */
func (r router) Register(path string, handle func(requireHandle RequireHandle) (result Result)) {
	r.routerMap[path] = handle
}

/**
 *  路由解析请求，并执行controller
 */
func (r router) Execute(res http.ResponseWriter, req *http.Request) {
	f, has := r.routerMap[req.URL.Path]
	if has {

		res.Header().Add("Access-Control-Allow-Origin", "*")
		res.Header().Add("Allow", "OPTIONS, TRACE, GET, HEAD, POST")
		res.Header().Add("Access-Control-Allow-Headers", "content-type")
		if req.Method == "OPTIONS"{
			fmt.Fprint(res, "")
			return
		}

		jsonForm := map[string]interface{}{}
		body, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(body, &jsonForm)

		req.ParseForm()
		fmt.Println(req)

		queryForm, err := url.ParseQuery(req.URL.RawQuery)

		result := f(RequireHandle{
			req,
			body,
			req.Form,
			queryForm,
			jsonForm,
		})

		jsonResult, err := json.Marshal(result)

		if err != nil {
			fmt.Fprint(res, "500")
		} else {
			fmt.Fprint(res, string(jsonResult))
		}

	} else {
		fmt.Fprint(res, "404")
	}
}

var Router = router{
	make(map[string]func(requireHandle RequireHandle) (result Result)),
}
