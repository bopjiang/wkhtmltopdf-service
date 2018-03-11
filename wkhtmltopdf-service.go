package main

import (
	"net/http"
	"os/exec"
	"log"
	"bytes"
)

func htmlToPdf(w http.ResponseWriter, r *http.Request){
	var args []string
	if r.Method == "GET" {
		if err := r.ParseForm(); err != nil{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		url := r.Form.Get("url")
		if url == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("err: url is empty"))
		}
		args = []string{url, "/dev/stdout"}

	}else if r.Method == "POST"{

	}

	cmd := exec.Command("wkhtmltopdf", args...)  
	var out bytes.Buffer  
	cmd.Stdout = &out 
	err := cmd.Run()  
	if err != nil {  
	    log.Fatalf("run wkhtmltopdf %v err, %s, %s", args, err, out.String()) 
	    return
	}   

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/pdf")
	_ ,err = out.WriteTo(w)
	if err != nil{
		log.Printf("failed to send response, %s", err)
	}
}


func main(){
	http.HandleFunc("/htmltopdf", htmlToPdf)
	if err := http.ListenAndServe("0.0.0.0:80", nil); err !=nil{
		log.Printf("failed to start http, %s", err)
	}
}