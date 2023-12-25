package main

// user.go
metrics.RequestCount.WithLabelValues(route).Inc()
wg.Add(1)
go func() {
	defer wg.Done()
}()
