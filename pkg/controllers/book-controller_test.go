package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/iamadarsh0309/Kindle-Go/pkg/routes"
)

func newTestRouter() *mux.Router {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	return r
}

func doJSON(t *testing.T, router *mux.Router, method, path string, body interface{}) (*httptest.ResponseRecorder, map[string]interface{}) {
	t.Helper()
	var reader *bytes.Reader
	if body != nil {
		raw, _ := json.Marshal(body)
		reader = bytes.NewReader(raw)
	} else {
		reader = bytes.NewReader(nil)
	}
	req := httptest.NewRequest(method, path, reader)
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	var out map[string]interface{}
	if res.Body.Len() > 0 && res.Header().Get("Content-Type") == "application/json" {
		json.Unmarshal(res.Body.Bytes(), &out)
	}
	return res, out
}

func TestCreateBookEndpoint(t *testing.T) {
	router := newTestRouter()
	res, out := doJSON(t, router, "POST", "/book/", map[string]string{"isbn": "CTRL-001", "title": "Create Test"})
	if res.Code != 201 {
		t.Fatalf("expected 201, got %d", res.Code)
	}
	if out["isbn"] != "CTRL-001" {
		t.Errorf("unexpected body: %v", out)
	}
}

func TestFullEndpointFlow(t *testing.T) {
	router := newTestRouter()

	// Create
	res, created := doJSON(t, router, "POST", "/book/", map[string]string{"isbn": "CTR-FLOW", "title": "Flow"})
	if res.Code != 201 {
		t.Fatalf("POST /book/: expected 201, got %d", res.Code)
	}
	id := created["id"]
	if id == nil {
		t.Fatalf("no id in create response: %v", created)
	}
	bookID, ok := id.(float64)
	if !ok {
		t.Fatalf("id is not a number: %v", id)
	}
	bookPath := "/book/" + strconv.Itoa(int(bookID))

	// Get all — created book present
	res, _ = doJSON(t, router, "GET", "/book/", nil)
	if res.Code != 200 {
		t.Fatalf("GET /book/: expected 200, got %d", res.Code)
	}
	var all []map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &all)
	if len(all) == 0 {
		t.Error("GET /book/ returned empty list")
	}

	// Get by id
	res, got := doJSON(t, router, "GET", bookPath, nil)
	if res.Code != 200 || got["title"] != "Flow" {
		t.Errorf("GET %s: got %d %v", bookPath, res.Code, got)
	}

	// Update
	res, updated := doJSON(t, router, "PUT", bookPath, map[string]string{"isbn": "CTR-FLOW", "title": "Updated"})
	if res.Code != 200 || updated["title"] != "Updated" {
		t.Errorf("PUT %s: got %d %v", bookPath, res.Code, updated)
	}

	// Delete
	res, _ = doJSON(t, router, "DELETE", bookPath, nil)
	if res.Code != 200 {
		t.Fatalf("DELETE %s: expected 200, got %d", bookPath, res.Code)
	}

	// Gone
	res, _ = doJSON(t, router, "GET", bookPath, nil)
	if res.Code != 404 {
		t.Errorf("GET deleted book: expected 404, got %d", res.Code)
	}
}
