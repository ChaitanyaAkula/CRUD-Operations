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
	"time"
//	"os"
	"CRUD/packages/cookie"
)
var dbconn *sql.DB
type UserData struct{
	Userid int `json:"userID"`
	Firstname string `json:"firstName"`
	Lastname string `json:"lastName"`
	Email string `json:"email"`

}
type UserData1 struct{

	Fullname string `json:"fullName"`
	Email string `json:"email"`
	Phone string `json:"phoneNumber"`
	Password string `json:"passwordString"`
}
func main(){

	dbconn,_=sql.Open("mysql","root:root@tcp(localhost:3306)/assignment")
	dbconn.SetConnMaxLifetime(time.Minute*3) 

	fmt.Println(dbconn)
	router:=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/signup",Ping)
	router.HandleFunc("/",Homepage)
	router.PathPrefix("/template/").Handler(http.FileServer(http.Dir(".")))
	
	router.HandleFunc("/api/v1/conn/delete",Delete).Methods("POST")
	postRouter:=router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/v1/conn/create",CreateFunction)

	getRouter:=router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/v1/conn/read/",ReadFunction)
	
	router.HandleFunc("/api/v1/conn/update/{username}/",UpdateFunction).Methods("GET","POST")
	router.HandleFunc("/api/v1/conn/signup",Register)
	log.Fatalln(http.ListenAndServe(":8080",router))

}
func Homepage(w http.ResponseWriter,r *http.Request){
	cookie,_:=r.Cookie("id")
	w.Header().Set("Content-Type","application-json")
	json.NewEncoder(w).Encode(`{"PAGE":`+cookie.Value+`}`)
}
func Ping(w http.ResponseWriter,r *http.Request){
	 t:=time.Now()
	cookie.SetCookie(w,"id","10")
	fmt.Fprint(w,t)
	//w.Header().Set("Content-Type","application-json")
	//json.NewEncoder(w).Encode(`{"success":`+t+`}`)
	
//	http.ServeFile(w,r,"/home/gitty/snap/go/golangtraining/CRUD-Operations/template/registration.html")
/*	file,err:=os.Open("/template/home.html")
	if err!=nil{
		http.Error(w,"file not found",404)

	}
	defer file.Close()
	f1,err:=file.Stat()
	if err!=nil{
		http.Error(w,"file not found",404)
		return
	}
	http.ServeContent(w,r,file.Name(),f1.ModTime(),file)*/
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
    email:=r.FormValue("q")
	responseData:=new(UserData)
	queryerr:=dbconn.QueryRow("SELECT iduser,firstname,lastname,email from user where email=?",email).Scan(&responseData.Userid,&responseData.Firstname,&responseData.Lastname,&responseData.Email)
	if queryerr!=nil{
		fmt.Println(queryerr.Error())
	}
	
	w.Header().Set("Content-Type","application-json")
	encoder:=json.NewEncoder(w)
	encoder.SetIndent("","   ")
	err:=encoder.Encode(responseData)
	if err!=nil{
		fmt.Println(err.Error())
	}
}

func UpdateFunction(w http.ResponseWriter,r *http.Request){
	email:=mux.Vars(r)["username"]
	
	stmt,errstmt:=dbconn.Prepare("UPDATE user SET firstname=? where email=?")
	if errstmt!=nil{
		fmt.Println("error Statement",errstmt.Error())
	}
	res,errRes:=stmt.Exec("krish",email)
	if errRes!=nil{
		fmt.Println("Error Response",errRes.Error())
	}
	fmt.Println(res)
	w.Header().Set("Content-Type","application-json")
	json.NewEncoder(w).Encode(`{"Message":"Updated Successfully"}`)
}
func Delete(w http.ResponseWriter,r *http.Request){
	requestData:=new(UserData)
	body,_:=ioutil.ReadAll(r.Body)
	err:=json.Unmarshal([]byte(body),&requestData)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(*requestData)
	
	stmt,errstmt:=dbconn.Prepare("DELETE from user where iduser=?")
	if errstmt!=nil{
		fmt.Println("error Statement",errstmt.Error())
	}
	res,errRes:=stmt.Exec(requestData.Userid)
	if errRes!=nil{
		fmt.Println("Error Response",errRes.Error())
	}
	fmt.Println(res)
}
func Register(w http.ResponseWriter,r *http.Request){
/*	fname:=r.FormValue("fname")
	email:=r.FormValue("email")
	phone:=r.FormValue("phone")
	pwd:=r.FormValue("pwd")
	fmt.Println(fname,email,phone,pwd)
	stmt,errstmt:=dbconn.Prepare("INSERT INTO user SET firstname=?,lastname=?,email=?")
	if errstmt!=nil{
		fmt.Println("error Statement",errstmt.Error())
	}
	res,errRes:=stmt.Exec(fname,"",email)
	if errRes!=nil{
		fmt.Println("Error Response",errRes.Error())
	}
	fmt.Println(res)*/
	fmt.Println("Hello, World")
	requestData:=new(UserData1)
	body,_:=ioutil.ReadAll(r.Body)
	err:=json.Unmarshal([]byte(body),&requestData)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(*requestData)
	//fmt.Fprint(w,"Got Values")
}