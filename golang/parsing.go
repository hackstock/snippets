package main

import(
	"net/http"
	"encoding/json"
	"time"
)

//User defines model for storing account details in database
type User struct {
	Username string
	Password string `json:"-"`
	IsAdmin bool
	CreatedAt time.Time
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/echo", echoHandler)

	http.ListenAndServe(":5000", mux)
}

func echoHandler(w http.ResponseWriter, r *http.Request){
	user := User{} //initialize empty user

	//Parse json request body and use it to set fields on user
	//Note that user is passed as a pointer variable so that it's fields can be modified
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		panic(err)
	}

	//Set CreatedAt field on user to current local time
	user.CreatedAt = time.Now().Local()

	//Marshal or convert user object back to json and write to response 
	userJson, err := json.Marshal(user)
	if err != nil{
		panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response 
	w.Write(userJson)
}