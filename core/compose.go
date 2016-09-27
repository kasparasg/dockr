package core

import (
	"context"

	compose "github.com/docker/libcompose/docker"
	ctx "github.com/docker/libcompose/docker/ctx"
	"github.com/docker/libcompose/project"
	"github.com/docker/libcompose/project/options"
)

type Compose struct {
}

func NewCompose() *Compose {
	return &Compose{}
}

func (c *Compose) Up(projectName string, composeFiles []string) error {
	project, err := compose.NewProject(&ctx.Context{
		Context: project.Context{
			ComposeFiles: composeFiles,
			ProjectName:  projectName,
		},
	}, nil)

	if err != nil {
		return err
	}

	return project.Up(context.Background(), options.Up{})
}
