package swagger

import (
	"net/http"

	"github.com/flowchartsman/swaggerui"

	_ "embed"
)

const (
	// Path is a path for swagger
	Path = "/docs"
	// Pattern is a pattern for mux that handles requests
	Pattern = "/docs/"
)

//go:embed hisho-core-service.swagger.local.json
var specLocal []byte

//go:embed hisho-core-service.swagger.k8s.json
var specK8S []byte

// HandlerLocal is required to route within local environment
var HandlerLocal = http.StripPrefix(Path, swaggerui.Handler(specLocal))

// HandlerK8S is required to route within k8s environment
var HandlerK8S = http.StripPrefix(Path, swaggerui.Handler(specK8S))
