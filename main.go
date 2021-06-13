package main

import(
	"fmt"
	"github.com/gorilla/mux"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
)
var dbconn *sql.DB
type UserData struct{
	Userid int `json:"userID"`
	Firstname string `json:"firstName"`
	Lastname string `json:"lastName"`
	Email string `json:"email"`
}
func main(){

	dbconn,_=sql.Open("mysql","root:root@tcp(localhost:3306)/assignment")
	
	fmt.Println(dbconn)
	router:=mux.NewRouter()
	router.HandleFunc("/",Ping)

	postRouter:=router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/v1/conn/create",CreateFunction)

	getRouter:=router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/v1/conn/read",ReadFunction)


	log.Fatalln(http.ListenAndServe(":8080",router))

}
func Ping(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application-json")
	json.NewEncoder(w).Encode(`{"success":true}`)
}
func CreateFunction(w http.ResponseWriter,r *http.Request){
	fmt.Println(dbconn)
	requestData:=new(UserData)
	body,_:=ioutil.ReadAll(r.Body)
	err:=json.Unmarshal([]byte(body),&requestData)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(*requestData)
	
	stmt,errstmt:=dbconn.Prepare("INSERT INTO user SET firstname=?,lastname=?,email=?")
	if errstmt!=nil{
		fmt.Println("error Statement",errstmt.Error())
	}
	res,errRes:=stmt.Exec(requestData.Firstname,requestData.Lastname,requestData.Email)
	if errRes!=nil{
		fmt.Println("Error Response",errRes.Error())
	}
	fmt.Println(res)
	
	
}
func ReadFunction(w http.ResponseWriter,r *http.Request){
	requestData:=new(UserData)
	body,_:=ioutil.ReadAll(r.Body)
	err:=json.Unmarshal([]byte(body),&requestData)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(*requestData)
}