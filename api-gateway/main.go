package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RunRequest struct {
        Language string
		Code  string
}

var req RunRequest


func index(w http.ResponseWriter,r *http.Request) {
	http.ServeFile(w,r,"index.html")
}


func handlerun(w http.ResponseWriter,r *http.Request)  {

	   json.NewDecoder(r.Body).Decode(&req)

	   data,_ := json.Marshal(req)

	   resp,err := http.Post("http://python-service:9002/run","application/json",bytes.NewBuffer(data))

	   if err != nil {
		    fmt.Println("Error while posting the code to the docker",err)
	   }

	   output,err_read := io.ReadAll(resp.Body)
	   
	   if err_read != nil {
              fmt.Println("Error in the reading output",err_read)
	   }
	   
	   w.Header().Set("Content-Type", "application/json")
	   w.Write(output)
}

func main() {
	 http.HandleFunc("/home",index)
	 http.HandleFunc("/run",handlerun)
	 fmt.Println("Server for api gateway is listning on 8000")
	 http.ListenAndServe(":8000",nil)
}