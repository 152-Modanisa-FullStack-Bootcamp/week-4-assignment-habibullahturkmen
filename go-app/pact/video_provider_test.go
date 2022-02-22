package pact

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"go-app/server"
	"os"
	"path"
	"runtime"
	"testing"
)

func TestVideoProvider(t *testing.T) {
	freePort, _ := utils.GetFreePort()
	myServer := server.NewServer()
	go myServer.StartServer(freePort)

	pact := dsl.Pact{
		Host:                     "127.0.0.1",
		Consumer:                 "FrontEnd",
		Provider:                 "BackEnd",
		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", freePort),
		PactURLs:        []string{
			"../vue-app/pact/pacts/frontend-backend.json",
		},
	}

	verifyProvider, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(verifyProvider), "pact test run")
}

// fixes relative path problems while testing
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}