package collector

import (
	"fmt"
	"os"

	"github.com/observiq/observiq-collector/internal/version"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service"
	"go.uber.org/zap"
)

const buildDescription = "observIQ's opentelemetry-collector distribution"

// NewSettings returns new settings for the collector.
func NewSettings(configPath string, loggingOpts []zap.Option) (service.CollectorSettings, error) {
	factories, err := DefaultFactories()
	if err != nil {
		return service.CollectorSettings{}, fmt.Errorf("failed to build factories: %w", err)
	}

	buildInfo := component.BuildInfo{
		Command:     os.Args[0],
		Description: buildDescription,
		Version:     version.Version,
	}

	provider := NewFileProvider(configPath)

	return service.CollectorSettings{
		Factories:               factories,
		BuildInfo:               buildInfo,
		LoggingOptions:          loggingOpts,
		ParserProvider:          provider,
		DisableGracefulShutdown: true,
	}, nil
}
