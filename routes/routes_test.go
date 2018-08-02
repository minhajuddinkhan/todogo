package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type RouteService struct {
	URL string
}

var (
	routeService = RouteService{
		URL: "http://localhost:3000",
	}
)

func (r *RouteService) Curl(route string, method string, body interface{}) (*http.Response, error) {

	b, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, r.URL+route, bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	return client.Do(req)

}

func TestAuthRouteWithCorrectCreds(t *testing.T) {

	body := LoginRequest{
		Username: "Rameez",
		Password: "123",
	}

	resp, err := routeService.Curl("/login", "POST", body)
	if err != nil {
		t.Errorf(fmt.Sprintf("Test failed: %s", err.Error()))
	}

	decoder := json.NewDecoder(resp.Body)
	auth := Authorization{}
	err = decoder.Decode(&auth)
	if err != nil {
		t.Errorf(fmt.Sprintf("Test failed: %s", err.Error()))

	}

	if len(auth.Authorization) == 0 {
		t.Errorf(fmt.Sprintf("Test failed: %s", err.Error()))

	}

}

func TestAuthRouteInCorrectCreds(t *testing.T) {

	body := LoginRequest{
		Username: "Rameez",
		Password: "1236",
	}

	_, err := routeService.Curl("/login", "POST", body)
	if err == nil {
		t.Errorf(fmt.Sprintf("Test failed: %s", err.Error()))
		return
	}

	// decoder := json.NewDecoder(resp.Body)
	// auth := Authorization{}
	// err = decoder.Decode(&auth)
	// if err != nil {
	// 	t.Errorf(fmt.Sprintf("Test failed: %s", err.Error()))

	// }

	// if len(auth.Authorization) == 0 {
	// 	t.Errorf(fmt.Sprintf("Test failed: %s", err.Error()))

	// }

}
