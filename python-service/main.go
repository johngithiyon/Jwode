package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

 type data struct {
	Language string
	Code string
}

func run(w http.ResponseWriter, r *http.Request) {

	var d data

	err := json.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }


   cmd := exec.Command("python3","-")

   stdin,_:= cmd.StdinPipe()

   func () {
        defer stdin.Close()
		io.WriteString(stdin,d.Code)
   }()


   output,err := cmd.CombinedOutput()

   if err != nil {
	    fmt.Println("Error in the command running")
   }
   
   w.Header().Set("Content-Type", "text/plain")
   w.Write(output)
}

func main() {
	http.HandleFunc("/run",run)
	fmt.Println("Docker python is listening in 9002")
	http.ListenAndServe(":9002",nil)
}