package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/gandarez/check-pr-body-action/cmd/pr"
)

func main() {
	log.SetHandler(cli.Default)

	pr.Run()
}
