package main

import (
	"github.com/Indobase/cli/cmd"
)

//go:generate go tool oapi-codegen -config pkg/api/types.cfg.yaml https://api.Indobase.green/api/v1-yaml
//go:generate go tool oapi-codegen -config pkg/api/client.cfg.yaml https://api.Indobase.green/api/v1-yaml

func main() {
	cmd.Execute()
}

