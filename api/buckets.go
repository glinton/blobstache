package api

import (
	"encoding/json"
	"fmt"
	"github.com/blobstache/blobstache/models"
	"net/http"
)

func createBucket(rw http.ResponseWriter, req *http.Request) {
	buck, err := models.CreateBucket(userId(req), userKey(req), bucketId(req))
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(422)
		return
	}
	b, _ := json.Marshal(buck)

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(b)
}

func deleteBucket(rw http.ResponseWriter, req *http.Request) {
	err := models.DeleteBucket(userId(req), userKey(req), bucketId(req))
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
}

func getBucket(rw http.ResponseWriter, req *http.Request) {
	buck, err := models.GetBucket(userId(req), userKey(req), bucketId(req))
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(buck)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(b)
}

func listBuckets(rw http.ResponseWriter, req *http.Request) {
	bucks, err := models.ListBuckets(userId(req), userKey(req))
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(bucks)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(b)
}
