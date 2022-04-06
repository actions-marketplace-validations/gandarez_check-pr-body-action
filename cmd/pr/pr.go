package pr

import (
	"os"

	"github.com/gandarez/check-pr-body-action/internal/exitcode"
	"github.com/gandarez/check-pr-body-action/internal/pr"

	"github.com/apex/log"
)

// Run validates if a pull request contains or not the specific words.
func Run() {
	p, err := LoadParams()
	if err != nil {
		log.Fatalf("failed to load parameters: %s", err)
		os.Exit(exitcode.ErrDefault)
	}

	pr := pr.New(p.Client)

	err = pr.CheckBody(p.PrNumber, p.Contains, p.NotContains)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(exitcode.ErrDefault)
	}

	os.Exit(exitcode.Success)
}
