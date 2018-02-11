package Server

import(
	"fmt"
	"net/http"
)

//User defines model for storing account details in database
var jsonToServe string
var urlToParse string
type Link struct {
	Link string
}
func serveRest(w http.ResponseWriter, r *http.Request) {
	response := jsonToServe
	urlToParse = r.URL.Path[1:len(r.URL.Path)])
	fmt.Fprintf(w,response)
}
	 
func ServeJson(jts string){
	jsonToServe = jts

	http.HandleFunc("/", serveRest)

	http.ListenAndServe("localhost:8080", nil)
}
func GetUrl() {
	return urlToParse
}




