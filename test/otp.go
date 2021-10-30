package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func hello(rollno string, coins int) {
	url := "https://iitkcoin-op711.herokuapp.com/wallet/add"
	var jsonStr = []byte(`{"rollno":"` + rollno + `", "coins":` + fmt.Sprintf("%d", coins) + `}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("cookie", "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xsbm8iOiIyMDA0MzMiLCJleHAiOjE2NDA5MzYyMzd9.5le5Lsh_PF2kFBi1tLX-vmII9svgrO7EKbWSTCXnC9k; Expires=Fri, 31 Dec 2021 07:37:17 GMT; Path=/; HttpOnly; Domain=localhost	")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// print response and status
	bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    bodyString := string(bodyBytes)
	fmt.Println(resp.Status, rollno, bodyString, coins)
}

func post_otp(rollno string) {
	url := "http://localhost:8080/auth/otp"
	var jsonStr = []byte(`{"rollno":"` + rollno + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// print response and status
	fmt.Println(resp.Status, rollno)
}

func main2(i int, j int) {
	for i := i; i < j; i++ {
		go hello(fmt.Sprintf("%d", (200123 + i%6)), rand.Intn(5))
	}
}

func main()  {
	if false {
		for i := 0; i < 10; i++ {
			// get random rollno
			rollno := rand.Intn(100000)
			go post_otp(fmt.Sprintf("%d+%d", i, rollno))
			time.Sleep(2 * time.Second)
		}
	}
	fmt.Println("done")
	for i := 0; i < 10; i++ {
		// get two random
		i := rand.Intn(10)
		j := 10+rand.Intn(10)
		go main2(i, j)
	}
	time.Sleep(2 * time.Minute)
}
