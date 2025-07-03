package test

import (
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"
)

var opts = godog.Options{
	Format: "cucumber",
	Paths: []string{
		"../features/order.feature",
		"../features/double11.feature",
	},
	Randomize: time.Now().UnixNano(),
	Output:    colors.Colored(os.Stdout),
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts) // godog v0.11.0 and later
}

func TestMain(m *testing.M) {
	pflag.Parse()

	status := godog.TestSuite{
		Name:                "AI-100x-SE-Join-Quest",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
