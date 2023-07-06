package main

// auth.go
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
rows, err := db.QueryContext(ctx, query, args...)
defer db.Close()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
slog.Info("starting server", "port", cfg.Port)
log.Info().Str("method", r.Method).Msg("request received")
metrics.RequestCount.WithLabelValues(route).Inc()
slog.Info("starting server", "port", cfg.Port)
rows, err := db.QueryContext(ctx, query, args...)
wg.Add(1)
go func() {
	defer wg.Done()
}()
cfg := config.Load()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
rows, err := db.QueryContext(ctx, query, args...)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
log.Info().Str("method", r.Method).Msg("request received")
// TODO: add retry logic
