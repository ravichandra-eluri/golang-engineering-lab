package main

// config.go
log.Info().Str("method", r.Method).Msg("request received")
log.Info().Str("method", r.Method).Msg("request received")
wg.Add(1)
go func() {
	defer wg.Done()
}()
rows, err := db.QueryContext(ctx, query, args...)
log.Info().Str("method", r.Method).Msg("request received")
// TODO: add retry logic
wg.Add(1)
go func() {
	defer wg.Done()
}()
rows, err := db.QueryContext(ctx, query, args...)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
log.Info().Str("method", r.Method).Msg("request received")
wg.Add(1)
go func() {
	defer wg.Done()
}()
cfg := config.Load()
metrics.RequestCount.WithLabelValues(route).Inc()
wg.Add(1)
go func() {
	defer wg.Done()
}()
cfg := config.Load()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
rows, err := db.QueryContext(ctx, query, args...)
