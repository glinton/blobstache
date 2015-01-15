package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"encoding/json"
)


func listUsers() {
	resp, err := roundTrip("GET", "/users")
	handleError(err)
	handleJson(resp)
}

func createUser() {
	resp, err := roundTrip("POST", "/users")
	handleError(err)
	handleJson(resp)
}

func deleteUser(id string) {
	resp, err := roundTrip("DELETE", "/users/"+id)
	handleError(err)
	handleBadStatus(resp)
	fmt.Println("Success")
}


func listBuckets() {
	resp, err := roundTrip("GET", "/buckets")
	handleError(err)
	handleJson(resp)
}

func createBucket() {
	resp, err := roundTrip("POST", "/buckets")
	handleError(err)
	handleJson(resp)
}

func showBucket(id string) {
	resp, err := roundTrip("GET", "/buckets/"+id)
	handleError(err)
	handleJson(resp)
}

func deleteBucket(id string) {
	resp, err := roundTrip("DELETE", "/buckets/"+id)	
	handleError(err)
	handleBadStatus(resp)
	fmt.Println("Success")
}


func listObjects() {
	resp, err := roundTrip("GET", "/objects")
	handleError(err)
	handleJson(resp)
}

func createObject() {
	// cant user round trip here
	req := buildRequest("POST", "/objects")	
	req.Body = os.Stdin
	resp, err := (&http.Client{}).Do(req)
	handleError(err)
	handleJson(resp)
}

func showObjectSize(id string) {
	resp, err := roundTrip("HEAD", "/objects/"+id)
	handleError(err)
	handleBadStatus(resp)
	fmt.Println(resp.Header.Get("Object-Size"))
}
func getObject(id string) {
	// must read full body and out put either to file or stdout
	resp, err := roundTrip("GET", "/objects/"+id)
	handleError(err)
	handleBadStatus(resp)
	io.Copy(os.Stdout, resp.Body)
}

func deleteObject(id string) {
	resp, err := roundTrip("DELETE", "/objects/"+id)	
	handleError(err)
	handleBadStatus(resp)
}

func roundTrip(method, path string) (*http.Response, error) {
	req := buildRequest(method, path)	
	return (&http.Client{}).Do(req)
}

func buildRequest(method, path string) *http.Request {
	if id == "" || key == "" {
		fmt.Println("I cant do that unless i have the userkey or userid")
		os.Exit(1)
	}

	req, _ := http.NewRequest(method, hostname+path, nil)
	req.Header.Add("Userid", id)
	req.Header.Add("Key", key)
	req.Header.Add("Bucketid", bucketid)
	req.Header.Add("Objectid", objectid)
	return req
}

func handleError(err error) {
	if err != nil {
		fmt.Println("An Error Happened:"+ err.Error())
		os.Exit(1)
	}
}

func handleJson(resp *http.Response) {
	handleBadStatus(resp)
	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	var i interface{}
	err = json.Unmarshal(body, &i)
	body, err = json.MarshalIndent(i, "", "  ")
	
	fmt.Println(string(body))
}

func handleBadStatus(resp *http.Response) {
	if resp.StatusCode/100 != 2 {
		fmt.Printf("Failure from server: %+v\n", resp)
		os.Exit(1)
	}
	
}