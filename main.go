package main
import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
)

type Movie struct{
	Director *dicrector `json:"director"`
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
}
type Director struct{
 Firstname  string `json:"firstname"`
 Lirstname  string `json:"lastname"`
}
var movies []Movie
func main(){
movies = make([]Movie)
}