package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func enableCORS(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}



func getContents() string {
	   contents := ""
	   return contents
}

func handle_create_folder(w http.ResponseWriter,r *http.Request) {

	    enableCORS(w)
     
	    name,err := io.ReadAll(r.Body)

		if err != nil {
			fmt.Println("Error while reading the folder name",err)
		}   

		foldername := string(name)

		path := "./"+ foldername

		errors := os.MkdirAll(path,os.ModePerm)

		w.Header().Set("Content-Type", "application/json")

		msg := "Sucessfully folder is created"
		

		if errors != nil {
			fmt.Println("Error while creating a folder",err)
		} else {
		     json.NewEncoder(w).Encode(msg)
		}

		
}

func handle_create_file(w http.ResponseWriter, r *http.Request) {

	enableCORS(w)

	  name,err:= io.ReadAll(r.Body)

	  if err != nil {
		   fmt.Println("Error while reading the file name",err)
	  }

	  filename:= string(name)

	  path := "./"+filename

	  content := getContents()

	  msg := "Sucessfully file is created"

	  file,errors  := os.Create(path)

	  if errors != nil {
		    fmt.Println("Error while creating the file")
	  } else {
		json.NewEncoder(w).Encode(msg)
	  }

	  file.WriteString(content)


	
}

func main() {
	http.HandleFunc("/folder",handle_create_folder)
	http.HandleFunc("/file",handle_create_file)
	fmt.Println("Server for file service is listening in 8001")
	http.ListenAndServe(":8001",nil)
}