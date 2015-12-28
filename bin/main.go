package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/gandi/docker-machine-gandi"
)

var Version string

func main() {
	plugin.RegisterDriver(gandi.NewDriver("", ""))
}
