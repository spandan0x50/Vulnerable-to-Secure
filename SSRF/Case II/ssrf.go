package main

import(
	"fmt"
	"net/http"
	"io"
	_"net/url"
	"strings"
)

func main(){
	http.HandleFunc("/", ssrf)
	err:=http.ListenAndServe(":8080", nil)
	if err!=nil{
		fmt.Println(err)
	}
}

func ssrf(w http.ResponseWriter, r *http.Request){
	ssrf:=r.URL.Query().Get("ssrf")
	if ssrf==""{
		return
	}
	if strings.Contains(ssrf, "vercel.app"){
		resp, err:=http.Get(ssrf)
		if err!=nil{
			fmt.Println(err)
		}
		body, err:=io.ReadAll(resp.Body)
		if err!=nil{
			fmt.Println(err)
		}
		io.WriteString(w, string(body))
	}
	w.WriteHeader(http.StatusForbidden)
	io.WriteString(w, "Not Allowed!")
}