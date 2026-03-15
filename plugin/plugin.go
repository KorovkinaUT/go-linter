package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/KorovkinaUT/go-linter/internal/analyzer"
)

type MyPlugin struct{}

var _ register.LinterPlugin = (*MyPlugin)(nil)

func (*MyPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}, nil
}

func (*MyPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}

func New(settings any) (register.LinterPlugin, error) {
	return &MyPlugin{}, nil
}

func init() {
	register.Plugin("loglint", New)
}
