//
//  tracer.go
//  newest_template
//
//  Created by karim-w on 12/07/2025.
//

package tracer

import (
	"os"

	"github.com/karim-w/coreruntime/runtimemodules"
	"github.com/karim-w/coreruntime_extensions/implementations/tracerviazaplogger"
)

const ENV_KEY = "TRACER_VARIANT"

var log_variant = os.Getenv(ENV_KEY)

type LogVariant string

const (
	LOGGED LogVariant = "logged"
	NONE   LogVariant = "none"
	// TODO: Implement other types
)

func Initialize() (runtimemodules.Tracer, error) {
	switch log_variant {
	case string(NONE):
		return nil, nil // No tracing
	default:
		return tracerviazaplogger.Initialize(), nil
	}
}
