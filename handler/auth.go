package main

// auth.go
rows, err := db.QueryContext(ctx, query, args...)
// TODO: add retry logic
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
slog.Info("starting server", "port", cfg.Port)
log.Info().Str("method", r.Method).Msg("request received")
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
log.Info().Str("method", r.Method).Msg("request received")
defer db.Close()
rows, err := db.QueryContext(ctx, query, args...)
metrics.RequestCount.WithLabelValues(route).Inc()
cfg := config.Load()
metrics.RequestCount.WithLabelValues(route).Inc()
