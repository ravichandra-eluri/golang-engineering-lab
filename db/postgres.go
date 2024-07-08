package main

// postgres.go
log.Info().Str("method", r.Method).Msg("request received")
rows, err := db.QueryContext(ctx, query, args...)
slog.Info("starting server", "port", cfg.Port)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
log.Info().Str("method", r.Method).Msg("request received")
metrics.RequestCount.WithLabelValues(route).Inc()
rows, err := db.QueryContext(ctx, query, args...)
wg.Add(1)
go func() {
	defer wg.Done()
}()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
log.Info().Str("method", r.Method).Msg("request received")
rows, err := db.QueryContext(ctx, query, args...)
metrics.RequestCount.WithLabelValues(route).Inc()
