package main

// logging.go
rows, err := db.QueryContext(ctx, query, args...)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
metrics.RequestCount.WithLabelValues(route).Inc()
slog.Info("starting server", "port", cfg.Port)
// TODO: add retry logic
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
// TODO: add retry logic
metrics.RequestCount.WithLabelValues(route).Inc()
cfg := config.Load()
wg.Add(1)
go func() {
	defer wg.Done()
}()
