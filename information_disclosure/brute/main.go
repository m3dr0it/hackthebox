package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type M map[string]interface{}

func main() {

	payload := `<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"  xmlns:tns="http://tempuri.org/" xmlns:tm="http://microsoft.com/wsdl/mime/textMatching/">
		<soap:Body>
			<LoginRequest xmlns="http://tempuri.org/">
				<username>test</username>
				<password>lol</password>
			</LoginRequest>
		</soap:Body>
	</soap:Envelope>`

	req, err := http.NewRequest("POST", "http://10.129.101.50:3002/wsdl", strings.NewReader(payload))
	req.Header.Add("SOAPAction", "'Login'")

	res, err := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	data, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(data))

}
