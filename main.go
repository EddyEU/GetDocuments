package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Document struct {
	ID   string
	Name string
	Size int64
}

// getDocumentsMD5 : File Path to MD5
func getDocumentsMD5(w http.ResponseWriter, r *http.Request) {

	var docs []Document
	//path := "C:\\go"
	path := "./files/"
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Error in accessing directory:", err)
		log.Fatal(err)
	}

	for _, file := range fileInfos {
		hash := hashFileToMD5CheckSum(path + "/" + file.Name())
		if err == nil {
			docs = append(docs, Document{ID: hash, Name: file.Name(), Size: file.Size()})
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)
}

// hashFileToMD5CheckSum : MD5 library
func hashFileToMD5CheckSum(filePath string) string {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String
}

// getDocumentById : File Path to MD5
func getDocumentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id := vars["id"]
	path := "./files/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		b, err := ioutil.ReadFile(path + f.Name())
		if err != nil {
			fmt.Print(err)
		}
		fileContent := string(b)
		if hashFileToMD5CheckSum(fileContent) == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Document{ID: hashFileToMD5CheckSum(fileContent), Name: f.Name(), Size: f.Size()})
		}
	}
}

func setDocument(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func deleteDocumentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	id := vars["id"]
	path := "./files/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		b, err := ioutil.ReadFile(path + f.Name())
		if err != nil {
			fmt.Print(err)
		}
		fileContent := string(b)
		fmt.Println(path + f.Name())
		if hashFileToMD5CheckSum(fileContent) == id {
			fullpath := path + f.Name()
			err := os.Remove(fullpath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// main : here
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/documents", getDocumentsMD5).Methods("GET")
	router.HandleFunc("/documents/{id}", getDocumentById).Methods("GET")
	router.HandleFunc("/documents", setDocument).Methods("POST")
	router.HandleFunc("/documents/{id}", deleteDocumentById).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", router))

}
