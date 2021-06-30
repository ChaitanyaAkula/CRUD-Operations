package cookie

import(
	"net/http"
	"time"
)
var min,_=time.ParseDuration("10000m")
var expire=time.Now().Add(min)

func SetCookie(w http.ResponseWriter,name,value string){
	http.SetCookie(w,&http.Cookie{
		Name:name,
		Value:value,
		Expires:expire,
		HttpOnly: true,
	})
}