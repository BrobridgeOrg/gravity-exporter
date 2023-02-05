package exporter

import (
	"context"

	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/BrobridgeOrg/gravity-exporter/pkg/configs"
	"github.com/BrobridgeOrg/gravity-exporter/pkg/connector"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var logger *zap.Logger

type Exporter struct {
	config    *configs.Config
	connector *connector.Connector
}

func New(lifecycle fx.Lifecycle, config *configs.Config, l *zap.Logger, c *connector.Connector) *Exporter {

	logger = l.Named("Exporter")

	s := &Exporter{
		config:    config,
		connector: c,
	}

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				return s.start()
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)

	return s
}

func (exporter *Exporter) start() error {

	logger.Info("Starting exporter")

	http.Handle("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//updateMetric()
		promhttp.Handler().ServeHTTP(w, r)
	}))

	return http.ListenAndServe(":8080", nil)
}
