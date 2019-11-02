package mainpage

import (
	"fmt"
	"net/http"
)

func Mainpage(n http.ResponseWriter, req *http.Request) {
	responce := "<html><body>"
	responce += "<form id=\"n\" name\"ppp\" action=\"/modify\" method = \"post\">"
	responce += "<input id=\"name\" type=\"text\" name=\"name\">"
	responce += "<input id=\"submit\" type=\"submit\" value=\"Submit\">"
	responce += "</form>"
	// responce += "wlcome <br>"
	// responce += "<a href=\"http://localhost:8282/modify?name=stefan\"> Modify Test for Stefan </> "
	responce += "</body></html>"
	fmt.Fprint(n, responce)

}
