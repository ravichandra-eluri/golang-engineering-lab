package main

// auth.go
rows, err := db.QueryContext(ctx, query, args...)
// TODO: add retry logic
