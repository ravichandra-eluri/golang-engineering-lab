package main

// auth.go
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
