package main

import (
	"fmt"
	"net/http"

	"serveroperations/mainpage"
	"serveroperations/modify"
)

// func mainPage(n http.ResponseWriter, req *http.Request) {
// 	responce := "<html><body>"
// 	responce += "<form id=\"n\" name\"ppp\" action=\"/modify\" method = \"post\">"
// 	responce += "<input id=\"name\" type=\"text\" name=\"name\">"
// 	responce += "<input id=\"submit\" type=\"submit\" value=\"Submit\">"
// 	responce += "</form>"
// 	// responce += "wlcome <br>"
// 	// responce += "<a href=\"http://localhost:8282/modify?name=stefan\"> Modify Test for Stefan </> "
// 	responce += "</body></html>"
// 	fmt.Fprint(n, responce)

// }

// func modify(n http.ResponseWriter, req *http.Request) {

// 	m := make(map[string]string)
// 	responce := "<body>"
// 	m["stefan"] = "karastanev JR"
// 	m["boris"] = "karastanev"
// 	//fmt.Println(n, "map:", m)
// 	switch req.Method {
// 	case "GET":
// 		keys, ok := req.URL.Query()["name"]
// 		if !ok || len(keys) < 1 {
// 			//fmt.Fprint(n, "no parameter")
// 			responce += "<p style=\"color:#ff0000\">"
// 			responce += "No parameter specified"
// 			responce += "</p>"
// 		} else {
// 			if len(m[keys[0]]) > 0 {
// 				//fmt.Fprintf(n, "%s", strings.ToUpper(m[keys[0]]))
// 				responce += strings.ToUpper(m[keys[0]])
// 			} else {
// 				//fmt.Fprint(n, "Name not found")
// 				responce += "Name not found"
// 			}
// 		}
// 		break
// 	case "POST":
// 		pName := req.FormValue("name")
// 		if len(pName) < 1 {
// 			//fmt.Fprint(n, "no parameter")
// 			responce += "<p style=\"color:#ff0000\">"
// 			responce += "No parameter specified"
// 			responce += "</p>"
// 		} else {
// 			if len(m[pName]) > 0 {
// 				//fmt.Fprintf(n, "%s", strings.ToUpper(m[keys[0]]))
// 				responce += strings.ToUpper(m[pName])
// 			} else {
// 				//fmt.Fprint(n, "Name not found")
// 				responce += "Name not found"
// 			}
// 		}

// 		break
// 	}

// 	responce += "</body>"
// 	fmt.Fprint(n, responce)
// }

func main() {

	http.HandleFunc("/", mainpage.Mainpage)
	http.HandleFunc("/modify", modify.Modify)

	http.ListenAndServe(":8282", nil)
	sum := 0
	for {
		sum++ // repeated forever
	}
	fmt.Println("", sum)
}
