package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/Knetic/govaluate"
)

var useDegrees = true // Toggle for degrees/radians

func main() {
	fmt.Println("Scientific Calculator CLI")
	fmt.Println("Enter expressions (e.g., 2 + 3, sin(30), sqrt(16), etc.) or 'exit' to quit:")
	fmt.Println("Type 'toggle' to switch between degrees and radians.")

	reader := bufio.NewReader(os.Stdin)

	for {
	
		fmt.Print("--> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

	
		if input == "exit" {
		break
	}

	result, err := calculate(preprocessExpression(input))
	
		if err != nil {
			fmt.Println("Error:", err)
	
		} else {
			fmt.Println("=", result)
	}
}
}

func preprocessExpression(expr string) string {
	// Handle implicit multiplication and missing parentheses
	expr = regexp.MustCompile(`(\d)([a-zA-Z])`).ReplaceAllString(expr, "$1*$2")
	expr = regexp.MustCompile(`([a-zA-Z]+)(\d+(\.\d+)?)`).ReplaceAllString(expr, "$1($2)")
	expr = strings.ReplaceAll(expr, "^", "**") // Convert ^ to ** for power

	return expr
}

func calculate(expression string) (interface{}, error) {
	functions := map[string]govaluate.ExpressionFunction{
		"sin": func(args ...interface{}) (interface{}, error) {
			return round(math.Sin(convertAngle(args[0]))), nil
		},
		"cos": func(args ...interface{}) (interface{}, error) {
			return round(math.Cos(convertAngle(args[0]))), nil
		},
		"tan": func(args ...interface{}) (interface{}, error) {
			return round(math.Tan(convertAngle(args[0]))), nil
		},
		"sqrt": func(args ...interface{}) (interface{}, error) {
			num := args[0].(float64)
			if num < 0 {
				return nil, fmt.Errorf("sqrt of negative number is undefined")
			}
			return round(math.Sqrt(num)), nil
		},
		"log": func(args ...interface{}) (interface{}, error) {
			num := args[0].(float64)
			if num <= 0 {
				return nil, fmt.Errorf("log of non-positive number is undefined")
			}
			return round(math.Log(num)), nil
		},
		"pow": func(args ...interface{}) (interface{}, error) {
			base := args[0].(float64)
			exponent := args[1].(float64)
			return round(math.Pow(base, exponent)), nil
		},
	}

	expr, err := govaluate.NewEvaluableExpressionWithFunctions(expression, functions)
	if err != nil {
		return nil, fmt.Errorf("invalid expression: %v", err)
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		return nil, fmt.Errorf("evaluation error: %v", err)
	}

	return result, nil
}

func convertAngle(val interface{}) float64 {
	angle := val.(float64)
	if useDegrees {
		return math.Pi * angle / 180
	}
	return angle
}

func round(val float64) float64 {
	return math.Round(val*1e10) / 1e10
}
