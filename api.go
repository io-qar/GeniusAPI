package main

import (
	"crypto/rand"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
)

func request() (string, int) {
	client := &http.Client{}

	var max = 1000
	rndm, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	rndmInt := int(rndm.Uint64())
	var song_id string = strconv.Itoa(rndmInt)

	req, err := http.NewRequest(http.MethodGet, "https://api.genius.com/songs/" + song_id, nil)
	CheckError(err)

	req.Header.Add("Authorization", "Bearer TKqINJxo3mdqtcIlf7ChIjCDLewQLpX5hCg831FReTOvpKKXio098yqrkr19TX6o")
	resp, err := client.Do(req)
	CheckError(err)

	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err)

	resp.Body.Close()
	return string(body[:]), rndmInt
}
