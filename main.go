package main

// main.go
metrics.RequestCount.WithLabelValues(route).Inc()
// TODO: add retry logic
metrics.RequestCount.WithLabelValues(route).Inc()
log.Info().Str("method", r.Method).Msg("request received")
