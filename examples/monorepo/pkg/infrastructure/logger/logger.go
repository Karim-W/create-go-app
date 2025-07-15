//
//  logger.go
//  newest_template
//
//  Created by karim-w on 12/07/2025.
//

package logger

import (
	"os"

	"github.com/karim-w/coreruntime/runtimemodules"
	"github.com/karim-w/coreruntime_extensions/implementations/loggerviazap"
)

const ENV_KEY = "LOG_VARIANT"

var log_variant = os.Getenv(ENV_KEY)

type LogVariant string

const (
	ZAP  LogVariant = "zap"
	NONE LogVariant = "none"
	// TODO: Implement other types
)

func Initialize() (runtimemodules.Logger, error) {
	switch log_variant {
	case string(NONE):
		// TODO: call the extension's no-op logger
		return nil, nil // No logging
	default:
		opts := loggerviazap.LoggerOptions{
			IncludeCaller:     os.Getenv("LOG_INCLUDE_CALLER") != "",
			IncludeStacktrace: os.Getenv("LOG_INCLUDE_STACKTRACE") != "",
			ShowPid:           os.Getenv("LOG_SHOW_PID") != "",
			OutErrtoStdout:    os.Getenv("LOG_OUT_ERR_TO_STDOUT") != "",
		}
		return loggerviazap.Initialize(opts), nil
	}
}
