package main

import (
	"golang.org/x/tools/go/analysis"

	"github.com/stbenjam/no-empty-http-proxy/pkg/analyzer"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin //nolint
