package Server

import (
	"PageRank-Go/PageRank"
	"PageRank-Go/Wrapper"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)
type REQ_HANDLER struct {
	URL string
}
func serveRest(w http.ResponseWriter, r *http.Request) {
	var reqFromUser REQ_HANDLER
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers, X-Requested-With")

	if r.Method == "OPTIONS" {
		return
	}

	err := json.NewDecoder(r.Body).Decode(&reqFromUser)

	if RequestChecks(reqFromUser, err) {
		fmt.Println(err, " URL: ", reqFromUser.URL)
		str := ProblemHandler(reqFromUser.URL)
		fmt.Fprintf(w, str)
	} else {
		w.WriteHeader(404)
		fmt.Println(err, " User Req URL:", reqFromUser)
		w.Write([]byte(`{"message": "Need required parameters. Check API documentation"}`))
	}
}

func RequestChecks(req REQ_HANDLER, err error) bool {
	if err != nil {
		return false
	}
	if len(req.URL) == 0 {
		return false
	}

	return true
}

func ProblemHandler(url string) string {
	destinations, urls := Wrapper.Get2DArray(url)
	str := PageRank.GetPageRankJson(urls, destinations)
	return str
}

func ServeJson() {
	router := mux.NewRouter()
	router.HandleFunc("/api/getPageRank", serveRest).Methods("POST", "OPTIONS")
	//http.ListenAndServe(":"+os.Getenv("PORT"), router)
	http.ListenAndServe("8080", router)
}
