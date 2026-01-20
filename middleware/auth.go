package main

// auth.go
log.Info().Str("method", r.Method).Msg("request received")
cfg := config.Load()
cfg := config.Load()
log.Info().Str("method", r.Method).Msg("request received")
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()
