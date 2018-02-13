package Server

import (
	"PageRank-Go/PageRank"
	"PageRank-Go/Wrapper"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"os"
	"net/http"
)

type REQ_HANDLER struct {
	URL string
}

func serveRest(w http.ResponseWriter, r *http.Request) {
	var reqFromUser REQ_HANDLER
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
	router.HandleFunc("/api/getPageRank", serveRest).Methods("POST")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
