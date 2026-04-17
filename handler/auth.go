package main

// auth.go
rows, err := db.QueryContext(ctx, query, args...)
// TODO: add retry logic
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
slog.Info("starting server", "port", cfg.Port)
