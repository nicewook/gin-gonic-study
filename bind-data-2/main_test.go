package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBindDataAPI(t *testing.T) {

	ts := httptest.NewServer(newServer())
	defer ts.Close()

	printResponseBody := func(resp *http.Response) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		t.Logf("StatusCode: %v, responseBody: %v", resp.StatusCode, string(body))
	}

	t.Run("GET Query OK", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/user?id=1&name=hsjeong&email=a@gmail.com", ts.URL))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
		printResponseBody(resp)
	})

	t.Run("POST JSON OK", func(t *testing.T) {
		account := User{
			ID:    1,
			Name:  "Hyunseok, Jeong",
			Email: "a@gmail.com",
		}
		b, _ := json.Marshal(account)
		buff := bytes.NewBuffer(b)
		resp, err := http.Post(fmt.Sprintf("%s/user", ts.URL), "application/json", buff)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
		printResponseBody(resp)
	})

	t.Run("PUT URI OK", func(t *testing.T) {
		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/user/1/hyunseokJeong/a@gmail.com", ts.URL), nil)
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
		printResponseBody(resp)
	})

	t.Run("PUT URI JSON", func(t *testing.T) {
		account := User{
			Name:  "Hyunseok, Jeong",
			Email: "a@gmail.com",
		}
		b, _ := json.Marshal(account)
		buff := bytes.NewBuffer(b)
		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/user/1", ts.URL), buff)
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
		printResponseBody(resp)
	})
}
