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

	t.Run("GET Query BadReqeust", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/user?id=a&password=qwer", ts.URL))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("Expected status code %v, got %v", http.StatusBadRequest, resp.StatusCode)
		}
		printResponseBody(resp)
	})

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

	t.Run("POST JSON BadReqeust", func(t *testing.T) {
		account := struct {
			Password string
		}{
			"qwer",
		}
		b, _ := json.Marshal(account)
		buff := bytes.NewBuffer(b)
		resp, err := http.Post(fmt.Sprintf("%s/user", ts.URL), "application/json", buff)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %v, got %v. even POST wrong JSON. it does not matter", http.StatusOK, resp.StatusCode)
		}
		t.Logf("even POST wrong JSON. it does not matter. StatusCode is %v", resp.StatusCode)
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

	t.Run("PUT URI BadReqeust", func(t *testing.T) {
		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/user/1/badname", ts.URL), nil)
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Fatalf("Expected status code %v, got %v", http.StatusNotFound, resp.StatusCode)
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

	t.Run("PUT URI JSON BadReqeust", func(t *testing.T) {
		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/user/1", ts.URL), nil)
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("Expected status code %v, got %v", http.StatusBadRequest, resp.StatusCode)
		}
		printResponseBody(resp)
	})

	t.Run("PUT URI JSON OK", func(t *testing.T) {
		account := User{
			ID:    3,
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
