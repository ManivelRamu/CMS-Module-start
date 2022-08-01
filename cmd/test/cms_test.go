package cmstest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func HomeTest(t *testing.T) {

	fmt.Println("Enter the token")
	var token string
	fmt.Scanln(&token)
	url := fmt.Sprintf("http://localhost:9010/movies?token=%s", token)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("content-type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		t.Errorf("got error on api request %s", err)
	}

	defer res.Body.Close()

	//Logging
	t.Logf("response status:%d\n", res.StatusCode)
	t.Logf("response Headers:%s\n", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	t.Logf("response Body:%s\n", string(body))

}
