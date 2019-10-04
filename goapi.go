package main
//each and every file in go must have a package name
import (
	"fmt"
	"net/http"
	"github.com/gorilla/handlers"

)
func main() {
	handler := http.NewServeMux()
	///we create a new router to expose our api
	//to our users
	handler.HandleFunc("/api/hello", SayHello)


	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	//Every time a  request is sent to the endpoint ("/api/hello")
	 //the function SayHello will be invoked
	http.ListenAndServe(":4000", handlers.CORS(headers,methods,origins)(handler))
	//we tell our api to listen to all request to port 8080.
}
func SayHello(w http.ResponseWriter, r *http.Request) {
	//notice how this function takes two parameters
	//the first parameter is a ResponseWriter writer and
	//this is where we write the response we want to send back to the user
	//in this case Hello world
	//the second parameter is a pointer of type  http.Request this holds
	//all information of the request sent by the user
	//this may include query parameters,path parameters and many more
	fmt.Fprintf(w, `API IS BEING CALLED`)
}
