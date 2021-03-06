package testutils

import (
	"net/http"
	"net/http/httptest"

	"github.com/selectel/go-selvpcclient/selvpcclient"
)

// TestEnv represents a testing environment for all resources.
type TestEnv struct {
	Mux    *http.ServeMux
	Server *httptest.Server
	Client *selvpcclient.ServiceClient
}

// SetupTestEnv prepares the new testing environment.
func SetupTestEnv() *TestEnv {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	testEnv := &TestEnv{
		Mux:    mux,
		Server: server,
	}

	return testEnv
}

// TearDownTestEnv releases the testing environment.
func (testEnv *TestEnv) TearDownTestEnv() {
	testEnv.Server.Close()
	testEnv.Server = nil
	testEnv.Mux = nil
	testEnv.Client = nil
}
