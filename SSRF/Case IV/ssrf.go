package main

import(
	"fmt"
	"net/http"
	"io"
	"net/url"
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
	u, err :=url.Parse(ssrf)
	if err!=nil{
		fmt.Println(err)
		return
	}
	hostname:=u.Hostname()
	if ssrf==""{
		return
	}
	if hostname=="vercel.app" || strings.HasSuffix(hostname, ".vercel.app"){
		resp, err:=http.Get(ssrf)
		if err!=nil{
			fmt.Println(err)
			return
		}
		finalURL:=resp.Request.URL
		finalH := finalURL.Hostname()
		if finalH=="vercel.app" || strings.HasSuffix(finalH, ".vercel.app"){
			body, err:=io.ReadAll(resp.Body)
			if err!=nil{
				fmt.Println(err)
				return
			}
			io.WriteString(w, string(body))
			return
		}
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, "Final redirect doesn't match the white-list \n")
	}
	w.WriteHeader(http.StatusForbidden)
	io.WriteString(w, "Not Allowed!")
}