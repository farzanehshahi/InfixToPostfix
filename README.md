# Infix To Postfix

Infix expression is a form in which operators are placed between operands. This form is easy for humans to read, and it is possible to prioritize terms with parentheses. But the computer does not prioritize the operators in this way. Postfix form is used in computers to recognize the priority of operators. In postfix form, operators appear after the operands.

Stack is used to convert infix expression to postfix expression because of its FILO structure.

The working method is that the received string of the infix form is checked from left to right, the operands go to the output, and the operators are pushed in the stack according to their priority.

The general rule for operators is that the new operator is compared to the operator on top of the stack, if it has higher priority, it is pushed onto the stack. Otherwise, the operators are popped off the stack until they have less than or equal priority to the new operator. Parentheses also have priority.

An open parenthesis has the highest priority outside the stack and the lowest priority inside the stack. The only operator that can pop opening parentheses from the stack is closing parentheses.


## Example

infix expression: `A+B` ⟶ postfix expression: `AB+`

infix expression: `(A+B)*(C/D)+12-F` ⟶ postfix expression: `AB+CD/*12+F-`


## Usage

First, make a stack of desired size and initialize the converter with it.

Then give your input string to the `ConvertInfixToPostfix` function. The output will be a string in the postfix form.

```go
package main

import (
	"fmt"
	"github.com/farzanehshahi/InfixToPostfix"
	"github.com/farzanehshahi/InfixToPostfix/stack"
)

func main() {
	// create stack
	stack := stack.New(50)
	// initialize the converter
	converter := InfixToPostfix.NewConverter(stack)

	// convert infix to postfix
	infix := "x/(z^2+w)*h/(a-b)^3+36"
	postfix, err := converter.ConvertInfixToPostfix(infix)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("postfix:", postfix) // output: xz2^w+/h*ab-3^/36+
}
```

          