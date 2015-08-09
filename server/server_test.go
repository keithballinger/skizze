package server

import (
	"io/ioutil"
	"strings"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)


type domainsResult struct {
	Result []string      `json:"result"`
	Error  error       `json:"error"`
}


func request(s *Server, t *testing.T, method string, domain string, body string) *httptest.ResponseRecorder {
	reqBody := strings.NewReader(body)
	req, err := http.NewRequest("GET", "http://counters.io/testdomain", reqBody)
	if err != nil {
		t.Fatalf("%s", err)
	}
	respw := httptest.NewRecorder()
	s.ServeHTTP(respw, req)
	return respw
}

func unmarschal(resp *httptest.ResponseRecorder) domainsResult {
	body, _ := ioutil.ReadAll(resp.Body)
	var r domainsResult
	json.Unmarshal(body, &r)
	return r
}


func TestDomainsInitiallyEmpty(t *testing.T) {
	s := New()
	resp := request(s, t, "GET", "testdomain", "")
	if resp.Code != 200 {
		t.Fatalf("Invalid Response Code %d - %s", resp.Code, resp.Body.String())	
		return
	}
	result := unmarschal(resp)
	if len(result.Result) != 0 {
		t.Fatalf("Initial resultCount != 0. Got %s", result)	
	}
}

func TestCreateDomain(t *testing.T) {
	s := New()
	resp := request(s, t, "POST", "testdomain", `{
		"domain": "testdomain",
		"domainType": "",
		"capacity": 1000,
		"values": []
	}`)

	if resp.Code != 200 {
		t.Fatalf("Invalid Response Code %d - %s", resp.Code, resp.Body.String())
		return
	}
	
	result := unmarschal(resp)
	if len(result.Result) != 1 {
		t.Fatalf("after add resultCount != 1. Got %s", result)	
	}
}