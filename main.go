package main

import "github.com/francismars/go-calculator/add"
import "github.com/francismars/go-calculator/sub"
import "github.com/francismars/go-calculator/div"
import "github.com/francismars/go-calculator/mult"
import ( "fmt" 
		"os"
	    "bufio"
		"strconv" )

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Two number calculation. eg: 3 + 3")

	fmt.Print("Insert first number:   ")
	scanner.Scan()
	firstNumberString := scanner.Text()

	fmt.Print("Insert Math Operator:  ")
	scanner.Scan()
	mathOperator := scanner.Text()

	fmt.Print("Insert second number:  ")
	scanner.Scan()
	secondNumberString := scanner.Text()

	firstNumber, err := strconv.Atoi(string(firstNumberString))
    if err != nil {
        fmt.Println("Insert Valid Number")
        os.Exit(2)
	}
	secondNumber, err := strconv.Atoi(string(secondNumberString))
    if err != nil {
        fmt.Println("Insert Valid Number")
        os.Exit(2)
    }

	var result int

	switch mathOperator {
		case "+":
			result = add.Two(firstNumber,secondNumber)
		case "-":
			result = sub.Two(firstNumber,secondNumber)
		case "*":
			result = mult.Two(firstNumber,secondNumber)
		case "/":
			result = div.Two(firstNumber,secondNumber)
		default:
			fmt.Println("Insert Correct Operator")
			os.Exit(2)
	}
	fmt.Println("Result:               ", result)
}