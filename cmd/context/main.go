package main

import (
	"context"
	"fmt"
)

type Context string

const (
	ContextKey Context = "username"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ContextKey", "John Doe")

	value := ctx.Value("ContextKey")
	strValue, ok := value.(int)
	fmt.Println(strValue)

	fmt.Println(ok)

}
