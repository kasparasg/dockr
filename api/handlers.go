package api

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"

	"github.com/kasparasg/dockr/core"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	type Version struct {
		Number string `json:"number"`
	}

	v, _ := json.Marshal(Version{"0.0.1"})
	rw.Write(v)
}

func CreateBuild(rw http.ResponseWriter, req *http.Request) {
	out := bufio.NewWriter(os.Stdout)

	deployer, _ := core.NewDocker(
		os.Getenv("DOCKER_HOST"),
		os.Getenv("DOCKER_CERT_PATH"),
	)

	go deployer.Deploy(out)

	out.Flush()
}
