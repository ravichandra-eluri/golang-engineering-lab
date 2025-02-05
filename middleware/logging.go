package main

// logging.go
rows, err := db.QueryContext(ctx, query, args...)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
