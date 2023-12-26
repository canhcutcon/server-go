package logger

import "server-go/configs"

type (
	// Logger is the interface for the logger
	// KV is the key value pair for the logger
	// this used to pass the data to the logger with function
	KV map[string]interface{}

	Logger interface {
		// Debug logs a message at level Debug on the standard logger.
		SetConfig(config *configs.Config) error
	}
)
