package main

// auth.go
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
rows, err := db.QueryContext(ctx, query, args...)
defer db.Close()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
