package InfixToPostfix

import (
	"errors"
	"github.com/farzanehshahi/InfixToPostfix/stack"
)

var (
	fullStackError = errors.New("stack is full")
)

type Converter struct {
	stack stack.Stack
}

func NewConverter(stack stack.Stack) *Converter {
	return &Converter{stack: stack}
}

// ConvertInfixToPostfix converts infix expression to postfix expression
func (c *Converter) ConvertInfixToPostfix(infix string) (string, error) {

	postfix := ""

	for _, ch := range infix {
		if c.isOperand(ch) {
			postfix += string(ch)
		} else if ch == '(' {
			if pushed := c.stack.Push(ch); !pushed {
				return "", fullStackError
			}
		} else if ch == ')' {
			for top, exist := c.stack.Top(); exist && top.(rune) != '('; top, exist = c.stack.Top() {
				op, _ := c.stack.Pop()
				postfix += string(op.(rune))
			}
			c.stack.Pop() // remove '(' from stack
		} else {
			if lastOperator, exist := c.stack.Top(); !exist {
				if pushed := c.stack.Push(ch); !pushed {
					return "", fullStackError
				}
			} else if c.getOperatorPriority(lastOperator.(rune)) < c.getOperatorPriority(ch) {
				if pushed := c.stack.Push(ch); !pushed {
					return "", fullStackError
				}
			} else {
				for top, exist := c.stack.Top(); exist && c.getOperatorPriority(top.(rune)) >= c.getOperatorPriority(ch); top, exist = c.stack.Top() {
					op, _ := c.stack.Pop()
					postfix += string(op.(rune))
				}
				if pushed := c.stack.Push(ch); !pushed {
					return "", fullStackError
				}
			}
		}
	}

	for !c.stack.IsEmpty() {
		op, _ := c.stack.Pop()
		postfix += string(op.(rune))
	}

	return postfix, nil
}

func (c *Converter) getOperatorPriority(operator rune) int {
	if operator == '+' || operator == '-' {
		return 1
	} else if operator == '*' || operator == '/' {
		return 2
	} else if operator == '^' {
		return 3
	} else {
		return 0
	}
}

func (c *Converter) isOperand(ch rune) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
		return true
	}
	return false
}
