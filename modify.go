package modify

import (
	"fmt"
	"net/http"
	"strings"
)

func Modify(n http.ResponseWriter, req *http.Request) { //modify function

	m := make(map[string]string)
	responce := "<body>"
	m["stefan"] = "karastanev JR"
	m["boris"] = "karastanev"
	//fmt.Println(n, "map:", m)
	switch req.Method {
	case "GET":
		keys, ok := req.URL.Query()["name"]
		if !ok || len(keys) < 1 {
			//fmt.Fprint(n, "no parameter")
			responce += "<p style=\"color:#ff0000\">"
			responce += "No parameter specified"
			responce += "</p>"
		} else {
			if len(m[keys[0]]) > 0 {
				//fmt.Fprintf(n, "%s", strings.ToUpper(m[keys[0]]))
				responce += strings.ToUpper(m[keys[0]])
			} else {
				//fmt.Fprint(n, "Name not found")
				responce += "Name not found"
			}
		}
		break
	case "POST":
		pName := req.FormValue("name")
		if len(pName) < 1 {
			//fmt.Fprint(n, "no parameter")
			responce += "<p style=\"color:#ff0000\">"
			responce += "No parameter specified"
			responce += "</p>"
		} else {
			if len(m[pName]) > 0 {
				//fmt.Fprintf(n, "%s", strings.ToUpper(m[keys[0]]))
				responce += strings.ToUpper(m[pName])
			} else {
				//fmt.Fprint(n, "Name not found")
				responce += "Name not found"
			}
		}

		break
	}

	responce += "</body>"
	fmt.Fprint(n, responce)
}
