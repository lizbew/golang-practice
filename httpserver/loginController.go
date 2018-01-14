package main
 
import (
    "net/http"
	"html/template"
	"log"
	"encoding/json"
	//"time"
)

type Result struct{
    Ret int
    Reason string
    Data interface{}
}
 
type loginController struct {
}
 
func (this *loginController)IndexAction(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("templates/html/login/index.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}

func (this *loginController)PostAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
    err := r.ParseForm()
    if err != nil {
        OutputJson(w, 0, "error params", nil)
        return
	}
	
	user_name := r.FormValue("name")
	user_password := r.FormValue("password")

	if user_name == "" || user_password == "" {
		OutputJson(w, 0, "error params", nil)
        return
	}

	//expiration := time.Unix(1, 0)
	cookie := http.Cookie{Name: "user", Value: user_name, Path: "/" }
	http.SetCookie(w, &cookie)
	OutputJson(w, 1, "ok", nil)
    return
}


func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
    out := &Result{ret, reason, i}
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}
