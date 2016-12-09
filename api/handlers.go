package api

import (
	"bufio"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"

	"github.com/kasparasg/dockr/core"
	"github.com/kasparasg/dockr/queue"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	type version struct {
		Number string `json:"number"`
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

	go client.BuildImage("build/express", out)

	out.Flush()
}

func Test(rw http.ResponseWriter, req *http.Request) {
	type response struct {
		Msg string `json:"string"`
	}

	rand.Seed(42)

	queue.WorkQueue <- queue.Job{Number: rand.Intn(100)}

	v, _ := json.Marshal(response{"Work queued"})

	rw.Write(v)
}
