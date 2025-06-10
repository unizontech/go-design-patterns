// strategy.go
package strategy

import "fmt"

type Strategy interface {
	Execute(a, b int) int
}

type Add struct{}
func (Add) Execute(a, b int) int { return a + b }

type Multiply struct{}
func (Multiply) Execute(a, b int) int { return a * b }

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c Context) Execute(a, b int) {
	fmt.Println("Result:", c.strategy.Execute(a, b))
}

