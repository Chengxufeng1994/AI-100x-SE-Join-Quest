package test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"
)

var t *testing.T

var opts = godog.Options{
	Format:    "cucumber",
	Paths:     []string{"../feature/chinesechess.feature"},
	Tags:      "~@Ignore",
	Randomize: time.Now().UnixNano(),
	Output:    colors.Colored(os.Stdout),
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts) // godog v0.11.0 and later
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = []string{"../feature"}
	status := godog.TestSuite{
		Name:                "godog",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})
	InitializeGeneralScenario(ctx)
}
