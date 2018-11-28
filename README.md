# GetDocuments

MIS - Maestria en Ingenieria de Software

author: Ing. Eddy Escalante Ustariz

respository: https://github.com/EddyEU/GetDocuments

date: 19/11/2018

/=====================================================================/

## PROBLEM

/=====================================================================/

CLASS EXERCISE 

```
-> Implement GetDocuments operation

-> Use a local folder to store files 

-> Return all files in the local folder 

-> ID should be file's MD5 checksum

-> [Base_Component] (https://github.com/timoteoponce/ds-persistence)
```

/=====================================================================/

## Download GO

/=====================================================================/

https://golang.org/dl/

* Install [Golang](https://golang.org/) for Windows.



/=====================================================================/

### How to compile

/=====================================================================/

        > go build main.go

/=====================================================================/

### How to run

/=====================================================================/

* Execute through go.

        > go run main.go 

* Open an Internet Browser [Localhost port 9000](http://localhost:9000/documents)


## Output

### Implement missing operations

```
 GET	        /documents	List of documents
 GET	        /documents/:id	Get documents with ID
 POST	        /documents	Create a document,having the file as body data
 DELETE 	/documents/:id	Deletes a document
```