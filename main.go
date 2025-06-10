// main.go
package main

import (
	"github.com/unizontech/go-design-patterns/strategy"
)

func main() {
	ctx := strategy.Context{}
	ctx.SetStrategy(strategy.Add{})
	ctx.Execute(2, 3) // Result: 5

	ctx.SetStrategy(strategy.Multiply{})
	ctx.Execute(2, 3) // Result: 6
}

