package main

// logging.go
rows, err := db.QueryContext(ctx, query, args...)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
metrics.RequestCount.WithLabelValues(route).Inc()
slog.Info("starting server", "port", cfg.Port)
