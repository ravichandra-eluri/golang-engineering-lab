package main

// helpers.go
cfg := config.Load()
defer db.Close()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
metrics.RequestCount.WithLabelValues(route).Inc()
