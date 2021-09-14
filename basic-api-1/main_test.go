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

func TestBasicAPI(t *testing.T) {

	ts := httptest.NewServer(newServer())
	defer ts.Close()

	printResponseBody := func(resp *http.Response) {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		t.Log(string(body))
	}

	t.Run("hello GET", func(t *testing.T) {
		t.Log(ts.URL)
		resp, err := http.Get(fmt.Sprintf("%s", ts.URL))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
		printResponseBody(resp)
	})

	t.Run("hello name GET", func(t *testing.T) {
		t.Log(ts.URL)
		resp, err := http.Get(fmt.Sprintf("%s/hsjeong", ts.URL))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
		printResponseBody(resp)
	})

	t.Run("json POST BadReqeust", func(t *testing.T) {
		t.Log(ts.URL)
		account := struct {
			Id int
		}{
			1,
		}
		b, _ := json.Marshal(account)
		buff := bytes.NewBuffer(b)
		resp, err := http.Post(fmt.Sprintf("%s/add", ts.URL), "application/json", buff)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Logf("Expected status code %v, got %v", http.StatusBadRequest, resp.StatusCode)
		}
		printResponseBody(resp)
	})

	t.Run("json POST", func(t *testing.T) {
		t.Log(ts.URL)
		account := Account{10, "Alex"}
		b, _ := json.Marshal(account)
		buff := bytes.NewBuffer(b)
		resp, err := http.Post(fmt.Sprintf("%s/add", ts.URL), "application/json", buff)
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
