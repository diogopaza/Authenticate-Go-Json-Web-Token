package main

import(

	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

)

type User struct{

	Username string `json:"username"`
	Password string `json:"password"`

}

type JwtToken struct{
	Token string `json:"token"`
}

type Exception struct{
	Message string `json:"message"`
}

func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request){

	w.Write([]byte("create end point"))
}

func ProtectedEndpoint(w http.ResponseWriter, req *http.Request){

	w.Write([]byte("protected end point"))
}


func main(){

	router := mux.NewRouter()
	fmt.Println("Starting the application....")
	
	router.HandleFunc("/authenticate", CreateTokenEndpoint).Methods("POST")
	router.HandleFunc("/protected", ProtectedEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))




}