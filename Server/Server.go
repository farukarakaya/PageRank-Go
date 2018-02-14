package Server

import (
	"PageRank-Go/PageRank"
	"PageRank-Go/Wrapper"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	s "strings"
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
		go ProblemHandler(reqFromUser.URL)
		w.WriteHeader(200)
		return
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

func ProblemHandler(url string) {
	destinations, urls := Wrapper.Get2DArray(url)

	if s.HasPrefix(url, "https://") {
		url = s.Split(url, "https://")[1]
	} else if s.HasPrefix(url, "http://") {
		url = s.Split(url, "http://")[1]
	}
	if s.HasSuffix(url, "/") {
		url = s.Split(url, "/")[0]
	}

	file, err := os.Open("file.go") // For read access.
	if err != nil {
		str := PageRank.GetPageRankJson(urls, destinations)

		// write the whole body at once
		err := ioutil.WriteFile(s.Join([]string{url, ".txt"}, ""), []byte(str), 0644)
		if err != nil {
			panic(err)
		}
	} else {
		return
	}

}
func getStatus(w http.ResponseWriter, r *http.Request) {
	var reqFromUser REQ_HANDLER
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers, X-Requested-With")

	if r.Method == "OPTIONS" {
		return
	}

	err := json.NewDecoder(r.Body).Decode(&reqFromUser)

	if RequestChecks(reqFromUser, err) {
		fmt.Println(err, " URL: ", reqFromUser.URL)
		requrl := reqFromUser.URL
		if s.HasPrefix(requrl, "https://") {
			requrl = s.Split(requrl, "https://")[1]
		} else {
			requrl = s.Split(requrl, "http://")[1]
		}
		if s.HasSuffix(requrl, "/") {
			requrl = s.Split(requrl, "/")[0]
		}
		file, err := ioutil.ReadFile(s.Join([]string{requrl, ".txt"}, ""))
		if err != nil {
			w.WriteHeader(200)
			w.Write([]byte(`{"message": "No file"}`))
			return
		} else {
			textNum := []byte(`{ "message":"Found", "content":[`) //file"] }"}
			textNum = append(textNum, file...)
			textNum = append(textNum, `]}`...)
			w.WriteHeader(200)
			w.Write([]byte(textNum))
			return
		}
	} else {
		w.WriteHeader(404)
		fmt.Println(err, " User Req URL:", reqFromUser)
		w.Write([]byte(`{"message": "Need required parameters. Check API documentation"}`))
	}
}

func ServeJson() {
	router := mux.NewRouter()
	router.HandleFunc("/api/getPageRank", serveRest).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getStatus", getStatus).Methods("POST", "OPTIONS")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
	//http.ListenAndServe(":8080", router)
}
