package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Epic Fail, NEED file path!")
		return
	}
	getDocuments(os.Args[1])
}

// getDocuments : File Path to MD5
func getDocuments(dir string) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error in accessing directory:", err)
	}

	for _, file := range fileInfos {
		hash, err := hashFileToMD5CheckSum(dir + "/" + file.Name())
		if err == nil {
			fmt.Printf("MD5 is : %s  -  Name is : %s  -  Size File: %d Kbps \n", hash, file.Name(), file.Size())
		}
	}
}

// hashFileToMD5CheckSum : MD5 library
func hashFileToMD5CheckSum(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}
