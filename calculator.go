package main
import (
        "bufio"
        "fmt"
        "math"
        "os"
        "strconv"
        "strings"
        "errors"
)



func main() {
        reader := bufio.NewReader(os.Stdin)

        fmt.Println("Scientific Calculator CLI")
        fmt.Println("Enter exps (e.g., 2 + 3, sin(0.5), sqrt(16), etc.) or 'exit' to quit:")

        for {
                fmt.Print("--> ")
                input, _ := reader.ReadString('\n')
                input = strings.TrimSpace(input)

                if strings.ToLower(input) == "exit" {
                break
                }

                result, err := evaluate(input)
                if err != nil {
                        fmt.Println("Error:", err)
                } else {
                        fmt.Println("=", result)
                }
        }
}



func evaluate(exp string) (float64, error) {
        exp = strings.ToLower(exp)

        //Handles Basic Arthimatic op
        if strings.Contains(exp, "+") {

	//checks if + is there and also splits the text like  1 + 2 will be split to "1 then + then 2" same for others
        parts := strings.Split(exp, "+")

	//Calculate is defined later | parts [0] would be 1 = a  and  [1] would be 2 = b same for others  
        return calculate(parts[0], parts[1], 
					
	//takes a and b float giving a + b same for others
	func(a, b float64) float64 { return a + b })

        } else if strings.Contains(exp, "-") {
                parts := strings.Split(exp, "-")
                return calculate(parts[0], parts[1], func(a, b float64) float64 { return a - b })


        } else if strings.Contains(exp, "*") {
                parts := strings.Split(exp, "*")
                return calculate(parts[0], parts[1], func(a, b float64) float64 { return a * b })


        } else if strings.Contains(exp, "/") {
                parts := strings.Split(exp, "/")
                return calculate(parts[0], parts[1], func(a, b float64) float64 { return a / b })


        } else if strings.Contains(exp, "^") { 
                parts := strings.Split(exp, "^")
                return calculate(parts[0], parts[1], func(a, b float64) float64 { return math.Pow(a, b) })
        }

        // Scientific Calculator 

        if strings.HasPrefix(exp, "sin(") && strings.HasSuffix(exp, ")") {
        valStr := exp[4 : len(exp)-1]
        val, err := strconv.ParseFloat(valStr, 64)
                if err != nil {
                        return 0, err
                }
                
                //Since we taking values in degree we have to convert it into radian
                rad := val * (math.Pi/180) 
                return math.Sin(rad), nil
	

        } else if strings.HasPrefix(exp, "cos(") && strings.HasSuffix(exp, ")") {
        valStr := exp[4 : len(exp)-1]
        val, err := strconv.ParseFloat(valStr, 64)
                if err != nil {
                return 0, err
                }			
	        rad := val * (math.Pi/180) 
                return math.Cos(rad), nil


        } else if strings.HasPrefix(exp, "tan(") && strings.HasSuffix(exp, ")") {
                valStr := exp[4 : len(exp)-1]
                val, err := strconv.ParseFloat(valStr, 64)
                if err != nil {
                        return 0, err
                }
                rad := val * (math.Pi/180) 
                return math.Tan(rad), nil


        } else if strings.HasPrefix(exp, "cot(") && strings.HasSuffix(exp, ")") {
                valStr := exp[4 : len(exp)-1]
                val, err := strconv.ParseFloat(valStr, 64)

                if err != nil {
                        return 0, err
                }

                rad := val * (math.Pi / 180)
                tanValue := math.Tan(rad)

                if tanValue == 0 {
                        return 0, errors.New("cot is undefined for this input")
                }
                return 1 / tanValue, nil       


        } else if strings.HasPrefix(exp, "sec(") && strings.HasSuffix(exp, ")") {
                valStr := exp[4 : len(exp)-1]
                val, err := strconv.ParseFloat(valStr, 64)

                if err != nil {
                        return 0, err
                }

                rad := val * (math.Pi / 180)
                cosValue := math.Cos(rad)

                if cosValue == 0 {
                        return 0, errors.New("sec is undefined for this input")
                }
                return 1 / cosValue, nil


        } else if strings.HasPrefix(exp, "csc(") && strings.HasSuffix(exp, ")") {
                valStr := exp[4 : len(exp)-1]
                val, err := strconv.ParseFloat(valStr, 64)

                if err != nil {
                        return 0, err
                }

                rad := val * (math.Pi / 180)
                sinValue := math.Sin(rad)

                if sinValue == 0 {
                        return 0, errors.New("cosec is undefined for this input")
                }
                return 1 / sinValue, nil

                
        } else if strings.HasPrefix(exp, "sqrt(") && strings.HasSuffix(exp, ")") {
                valStr := exp[5 : len(exp)-1]
                val, err := strconv.ParseFloat(valStr, 64)
                if err != nil {
                        return 0, err
                }
                if val < 0 {
                        return 0, fmt.Errorf("square root of negative number")
                }
                return math.Sqrt(val), nil


        } else if strings.HasPrefix(exp, "log(") && strings.HasSuffix(exp, ")") { 
                valStr := exp[4 : len(exp)-1]
                val, err := strconv.ParseFloat(valStr, 64)
                if err != nil {
                        return 0, err
                }
                if val <= 0 {
                        return 0, fmt.Errorf("logarithm of non-positive number")
                }
                return math.Log(val), nil
        }



	//Return Error if invalid
        val, err := strconv.ParseFloat(exp, 64)
        if err != nil {
                return 0, fmt.Errorf("invalid exp")
        }
        return val, nil
}

// Reading user input through the CLI
// Converts string inputs to float64 and performs the given mathematical operation
func calculate(operand1Str, operand2Str string, 
        // Takes two string operands and an operation function, returns the result or an error
        operation func(float64, float64) float64) (float64, error) {
                
        // Trimming leading and trailing whitespaces
        operand1Str = strings.TrimSpace(operand1Str)
        operand2Str = strings.TrimSpace(operand2Str)        

                
        //Parse converts string to float
        operand1, err := strconv.ParseFloat(operand1Str, 64)

        // Return an error if the input string cannot be parsed as a float
        if err != nil {
                return 0, err
        }
        operand2, err := strconv.ParseFloat(operand2Str, 64)
        if err != nil {
                return 0, err
        }
        return operation(operand1, operand2), nil
}



