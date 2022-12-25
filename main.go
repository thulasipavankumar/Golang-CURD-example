package main
import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	_ "flag"
)

type Movie struct{
	
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}
type Director struct{
 Firstname  string `json:"firstname"`
 Lirstname  string `json:"lastname"`
}
var movies []Movie
func main(){
	
	movies = append(movies,Movie{"1","1234","Movie one",&Director{"firstname","lastname"}})
	r:=mux.NewRouter()
	
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	fmt.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8000",r))
	
}
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
	log.Printf("Called get Movies")
	
}
func getMovie(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("Called get Movie  on id %v",id)
	for _,item := range movies{
		if item.ID == id {
			res, _ := json.Marshal(item)
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(res)
			return 	
		}
	}
	res, _ := json.Marshal("{error : given id not found in the list}")
	w.WriteHeader(http.StatusNotFound)
	w.Write(res)
	
}
func deleteMovie( w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	log.Printf("Called delete Movie  on id  %v",vars["id"])
	for index,item := range movies{
		if item.ID == vars["id"] {
			movies=append(movies[:index],movies[index+1:]...)
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	res, _ := json.Marshal("{error : given id not found in the list}")
	w.WriteHeader(http.StatusNotFound)
	w.Write(res)
	
}
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	log.Printf("Called create Movie ")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
    movies = append(movies,	movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("Called update Movie %v",id)
	for index,item := range movies{
		if item.ID == id {
			movies=append(movies[:index],movies[index+1:]...)		
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = id
    		movies = append(movies,	movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	res, _ := json.Marshal("{error:given id not found in the list}")
	w.WriteHeader(http.StatusNotFound)
	w.Write(res)
	
}