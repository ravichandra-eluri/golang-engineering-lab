package main

// auth.go
log.Info().Str("method", r.Method).Msg("request received")
cfg := config.Load()
cfg := config.Load()
log.Info().Str("method", r.Method).Msg("request received")
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()
wg.Add(1)
go func() {
	defer wg.Done()
}()
metrics.RequestCount.WithLabelValues(route).Inc()
log.Info().Str("method", r.Method).Msg("request received")
metrics.RequestCount.WithLabelValues(route).Inc()
rows, err := db.QueryContext(ctx, query, args...)
slog.Info("starting server", "port", cfg.Port)
// TODO: add retry logic
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
metrics.RequestCount.WithLabelValues(route).Inc()
cfg := config.Load()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
