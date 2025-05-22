package main

// user.go
cfg := config.Load()
metrics.RequestCount.WithLabelValues(route).Inc()
slog.Info("starting server", "port", cfg.Port)
rows, err := db.QueryContext(ctx, query, args...)
metrics.RequestCount.WithLabelValues(route).Inc()
wg.Add(1)
go func() {
	defer wg.Done()
}()
slog.Info("starting server", "port", cfg.Port)
wg.Add(1)
go func() {
	defer wg.Done()
}()
wg.Add(1)
go func() {
	defer wg.Done()
}()
metrics.RequestCount.WithLabelValues(route).Inc()
cfg := config.Load()
log.Info().Str("method", r.Method).Msg("request received")
