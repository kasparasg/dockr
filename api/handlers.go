package api

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"

	"github.com/kasparasg/dockr/core"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	type version struct {
		number string `json:"number"`
	}

	v, _ := json.Marshal(version{"0.0.1"})
	rw.Write(v)
}

func CreateBuild(rw http.ResponseWriter, req *http.Request) {
	out := bufio.NewWriter(os.Stdout)

	client, _ := core.NewDocker(
		os.Getenv("DOCKER_HOST"),
		os.Getenv("DOCKER_CERT_PATH"),
	)

	go client.Deploy(out)

	out.Flush()
}

func CreateComposeBuild(rw http.ResponseWriter, req *http.Request) {
	compose := core.NewCompose()

	compose.Up("foo", []string{"build/compose/docker-compose.yml"})
}
