package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Consumer struct {
	ServiceAddress string
}

func (c Consumer) Hello() (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.ServiceAddress+"/", nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func TestServeHttp(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	consumer := Consumer{ServiceAddress: srv.URL}
	res, err := consumer.Hello()
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	if http.StatusOK != res.StatusCode {
		t.Errorf("expected %v, got %v\n", http.StatusOK, res.StatusCode)
	}
	bodyString := string(bodyBytes)
	if bodyString != "hello world" {
		t.Errorf("expected %v, got %v\n", "hello world", bodyString)
	}
}

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/", serveHttpMock)

	srv := httptest.NewServer(handler)

	return srv
}

func serveHttpMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world"))
}
