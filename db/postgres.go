package main

// postgres.go
log.Info().Str("method", r.Method).Msg("request received")
rows, err := db.QueryContext(ctx, query, args...)
