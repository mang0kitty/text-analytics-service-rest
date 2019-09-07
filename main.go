package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {

	var subscriptionKeyVar string = "0bb36e457c9c45d2909afc26aab522fa"
	// if "" == os.Getenv(subscriptionKeyVar) {
	// 	log.Fatal("Please set/export the environmental variable " + subscriptionKeyVar + ".")
	// }
	//var subscriptionKey string = os.Getenv(subscriptionKeyVar)
	var endpointVar string = "https://cognitiveservicesaideentest.cognitiveservices.azure.com/"
	// if "" == os.Getenv(endpointVar) {
	// 	log.Fatal("Please set/export the environment variable" + endpointVar + ".")
	// }
	//var endpoint string = os.Getenv(endpointVar)

	const uriPath = "/text/analytics/v2.1/languages"
	var uri = endpointVar + uriPath

	data := []map[string]string{
		{"id": "1", "text": "This is a document written in English."},
		{"id": "2", "text": "Este es un document escrito en Español."},
		{"id": "3", "text": "这是一个用中文写的文件"},
	}

	//Marshal returns the JSON encoding of v.
	documents, err := json.Marshal(&data)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}
	r := strings.NewReader("{\"documents\":" + string(documents) + "}")
	fmt.Printf("the string new reader is %s", r)

	client := &http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest("POST", uri, r)
	if err != nil {
		fmt.Printf("Error creating request %v/n", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKeyVar)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error on request %v/n", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body %v/n", err)
		return
	}
	var f interface{}

	//Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v
	json.Unmarshal(body, &f)

	//MarshalIndent is like Marshal but applies Indent to format the output.
	//Each JSON element in the output will begin on a new line beginning with prefix followed by one or more copies of indent according to the indentation nesting.
	jsonFormatted, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		fmt.Printf("Error producing JSON %v/n", err)
		return
	}
	fmt.Println(string(jsonFormatted))
}
