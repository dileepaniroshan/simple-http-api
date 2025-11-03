package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHelloWorldHandler(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   interface{}
	}{
		{"valid name A", "name=Alice", http.StatusOK, Response{Message: "Hello Alice"}},
		{"valid name lowercase a", "name=alice", http.StatusOK, Response{Message: "Hello alice"}},
		{"valid name M", "name=Mike", http.StatusOK, Response{Message: "Hello Mike"}},
		{"invalid name N", "name=Nancy", http.StatusBadRequest, ErrorResponse{Error: "Invalid Input"}},
		{"invalid name z", "name=zane", http.StatusBadRequest, ErrorResponse{Error: "Invalid Input"}},
		{"invalid name Z", "name=Zane", http.StatusBadRequest, ErrorResponse{Error: "Invalid Input"}},
		{"missing name", "", http.StatusBadRequest, ErrorResponse{Error: "Invalid Input"}},
		{"empty name", "name=", http.StatusBadRequest, ErrorResponse{Error: "Invalid Input"}},
		{"whitespace name", "name=   ", http.StatusBadRequest, ErrorResponse{Error: "Invalid Input"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			
			if tt.queryParam == "" {
				req = httptest.NewRequest("GET", "/hello-world", nil)
			} else if tt.name == "whitespace name" {
				req = httptest.NewRequest("GET", "/hello-world", nil)
				q := url.Values{}
				q.Set("name", "   ")
				req.URL.RawQuery = q.Encode()
			} else {
				req = httptest.NewRequest("GET", "/hello-world?"+tt.queryParam, nil)
			}
			
			rr := httptest.NewRecorder()
			helloWorldHandler(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("got status %d, want %d", rr.Code, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusOK {
				var resp Response
				json.NewDecoder(rr.Body).Decode(&resp)
				expected := tt.expectedBody.(Response)
				if resp.Message != expected.Message {
					t.Errorf("got %s, want %s", resp.Message, expected.Message)
				}
			} else {
				var errResp ErrorResponse
				json.NewDecoder(rr.Body).Decode(&errResp)
				expected := tt.expectedBody.(ErrorResponse)
				if errResp.Error != expected.Error {
					t.Errorf("got %s, want %s", errResp.Error, expected.Error)
				}
			}
		})
	}
}
