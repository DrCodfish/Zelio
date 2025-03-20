package compiler

import (
	"fmt"
	"strconv"
)

// VariableType represents the type of a variable in Zelio
type VariableType int

const (
	IntType VariableType = iota
	FloatType
	StringType
	BoolType
)

// Variable represents a variable in Zelio
type Variable struct {
	Name  string
	Type  VariableType
	Value interface{}
}

// DeclareVariable declares a new variable
func DeclareVariable(name string, varType VariableType, value interface{}) Variable {
	return Variable{
		Name:  name,
		Type:  varType,
		Value: value,
	}
}

// Println prints a message to the console
func Println(args ...interface{}) {
	fmt.Println(args...)
}

// Add adds two integers
func Add(a, b int) int {
	return a + b
}

// Subtract subtracts two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide divides two integers
func Divide(a, b int) int {
	return a / b
}

// Concatenate concatenates two strings
func Concatenate(a, b string) string {
	return a + b
}

// Length returns the length of a string
func Length(s string) int {
	return len(s)
}