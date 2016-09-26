package core

import (
	"io"

	log "github.com/Sirupsen/logrus"
	docker "github.com/fsouza/go-dockerclient"
)

type Docker struct {
	api    string
	certs  string
	client *docker.Client
}

func NewDocker(api string, certs string) (*Docker, error) {
	log.Info("Using Docker api: ", api)

	var client *docker.Client
	var err error

	client, err = docker.NewTLSClient(
		api,
		certs+"/cert.pem",
		certs+"/key.pem",
		certs+"/ca.pem",
	)

	if err != nil {
		return nil, err
	}

	return &Docker{api, certs, client}, nil
}

func (d *Docker) Deploy(out io.Writer) error {
	d.BuildImage("build/express", out)

	return nil
}

func (d *Docker) BuildImage(context string, out io.Writer) error {
	opts := docker.BuildImageOptions{
		ContextDir:   context,
		OutputStream: out,
	}

	err := d.client.BuildImage(opts)

	if err != nil {
		return err
	}

	return nil
}
