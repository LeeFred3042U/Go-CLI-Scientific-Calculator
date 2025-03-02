# Scientific Calculator CLI

This is a command-line interface (CLI) scientific calculator built with Go. It handles basic arithmetic operations, trigonometric functions, logarithms, and more. Enter expressions directly, and the program will evaluate and return the result.

## Features

- **Basic Arithmetic Operations:** Addition, subtraction, multiplication, division, and exponentiation (`+`, `-`, `*`, `/`, `^`).
- **Trigonometric Functions:** `sin`, `cos`, `tan`, `cot`, `sec`, `csc` (input in degrees).
- **Logarithmic and Root Functions:** `log`, `sqrt`.
- **Error Handling:** Detects invalid expressions, division by zero, and undefined operations.

## How to Use

1. **Compile the Code:**  
   ```bash
   go build calculator.go
2. **Enter Expression:**
-   Arithmetic: 2 + 3, 4 * 5, 10 / 2
-   Exponents: 2 ^ 3
-   Trigonometry: sin(30), cos(60), tan(45)
-   Inverse Trig: cot(45), sec(60), csc(30)
-   Roots & Logs: sqrt(16), log(10)
3. **Exit Program:**
-  For quiting the program type
   ```bash
   exit
