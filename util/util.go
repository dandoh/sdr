package util

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"bytes"
)

func PrintBody(req *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(req.Body)
	bodyString := string(bodyBytes)
	fmt.Println("\n\n", bodyString, "\n\n")
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}
