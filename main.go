package onion

import ()

type Onion struct{}

func (c Onion) Do(a int) (int, error) {
	return a, nil
}

// A Client manipulate int and returns int or errors in case of failure.
type Client interface {
	Do(int) (int, error)
}

// ClientFunc is a function type that implements the Client interface.
type ClientFunc func(int) (int, error)

func (f ClientFunc) Do(a int) (int, error) {
	return f(a)
}

// A Decorator wraps a Client with extra behaviour.
type Decorator func(Client) Client

// Add returns a Decorator that increases Client's input.
func Add(n int) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(a int) (int, error) {
			return c.Do(a + n)
		})
	}
}

// Mul returns a Decorator that multiplies Client's input.
func Mul(m int) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(a int) (int, error) {
			return c.Do(a * m)
		})
	}
}

// Decorate decorates a Client c with all the given Decorators, in order.
func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}
