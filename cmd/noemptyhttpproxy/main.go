package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/stbenjam/no-empty-http-proxy/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
