package main

import (
	"log"
	"net/http"
	"crypto/tls"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

func main() {

	for j:=0; j<300; j++ {
		result := make (chan int)
		for i:=0; i<50; i++ {
			go runTest(result)
		}
		for i:=0; i<50; i++ {
			<-result
		}
	}

}

func runTest(result chan<- int) {
	respCode, _ := getPortal()
	//log.Println("Portal object - ", portalObj)
	if respCode != http.StatusOK {
		result <- respCode
	}
	result <- respCode
	//go getPortalJwt(portalObj)

}

func getPortal() (respCode int, respObj map[string]interface{}) {

	log.Println("Get portal")
	// Load our TLS key pair to use for authentication
	cert, err := tls.LoadX509KeyPair("/vagrant_data/apps/misc/SSETest/certs/portal/v2/ac.cert", "/vagrant_data/apps/misc/SSETest/certs/portal/v2/ac.key")
	if err != nil {
		log.Fatalln("Unable to load cert", err)
	}

	//disabling TLS check.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, Certificates: []tls.Certificate{cert}},
	}

	inData := bytes.NewReader([]byte(""))
	req, err := http.NewRequest("GET", "https://staging-sse.cisco.com/providers/sse/services/scim/v2/Portals/b7eb9785-f885-4bfa-ad1c-af7186e866eb", inData)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Client error - ", err)
	}
	log.Println("Sending request")
	defer resp.Body.Close()
	tr.CloseIdleConnections()

	outData, _ := ioutil.ReadAll(resp.Body)

	response := make(map[string]interface{})
	json.Unmarshal(outData, &response)

	log.Println("Response code - ", resp.StatusCode)
	//log.Println("Response object", response)

	return resp.StatusCode, response
}

func getPortalJwt(portalObj map[string]interface{}) (respCode int, respObj map[string]interface{}) {
	// Load our TLS key pair to use for authentication
	cert, err := tls.LoadX509KeyPair("/Users/krishnaprasadbandakothur/DEVELOP/apps/misc/SSETest/certs/portal/v2/ac.cert", "/Users/krishnaprasadbandakothur/DEVELOP/apps/misc/SSETest/certs/portal/v2/ac.key")
	if err != nil {
		log.Fatalln("Unable to load cert", err)
	}

	//disabling TLS check.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, Certificates: []tls.Certificate{cert}},
	}

	log.Println("Portal Id - ", portalObj["Resources"])
	inData := bytes.NewReader([]byte(`{"principal": { "portalId": ` + portalObj["Resources"].(map[string]interface{})["id"].(string) + `},"expiry": 129600}`))
	req, err := http.NewRequest("POST", "https://staging-sse.cisco.com/providers/sse/services/token/api/v2/jwt/access", inData)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	tr.CloseIdleConnections()

	outData, _ := ioutil.ReadAll(resp.Body)

	response := make(map[string]interface{})
	json.Unmarshal(outData, &response)

	log.Println("Response code - ", resp.StatusCode)
	log.Println("Response object", response)

	return resp.StatusCode, response
}