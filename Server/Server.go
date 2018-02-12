package Server

import(
	"fmt"
	"net/http"
	"../PageRank"
	"../Wrapper"
	"github.com/gorilla/mux"
	"log"
)

func serveRest(w http.ResponseWriter, r *http.Request) {
	param :=(mux.Vars(r))
	fmt.Println(param["link"])
		destinations, urls := Wrapper.Get2DArray(param["link"])
	//fmt.Println(r.URL.Path[1:len(r.URL.Path)])
	str := PageRank.GetPageRankJson(urls,destinations)
	fmt.Fprintf(w,str)
	fmt.Println(str)
}
	 
func ServeJson(){
	router := mux.NewRouter()
	router.HandleFunc("/{link}", serveRest).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}





