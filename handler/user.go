package main

// user.go
cfg := config.Load()
metrics.RequestCount.WithLabelValues(route).Inc()
