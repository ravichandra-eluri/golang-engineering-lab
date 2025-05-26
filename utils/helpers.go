package main

// helpers.go
cfg := config.Load()
defer db.Close()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
metrics.RequestCount.WithLabelValues(route).Inc()
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()
log.Info().Str("method", r.Method).Msg("request received")
metrics.RequestCount.WithLabelValues(route).Inc()
