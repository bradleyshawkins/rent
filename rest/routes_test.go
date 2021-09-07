package rest_test

import "os"

func getServiceURL() string {
	u := os.Getenv("SERVICE_URL")
	if u == "" {
		u = "http://127.0.0.1:8080"
	}
	return u
}
