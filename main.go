package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkInput(input string) (int, int, string, bool, []error) {
	var errorList []error

	splitInput := strings.Split(strings.TrimSpace(input), " ")
	if len(splitInput) != 3 {
		errorList = append(errorList, errors.New("введенные данные не соответствуют требуемому формату"))
		return 0, 0, "", false, errorList
	}
	firstNum, secondNum := splitInput[0], splitInput[2]
	sign := splitInput[1]

	if checkSign(sign) {
		errorList = append(errorList, errors.New("некорректный ввод знака"))
	}

	firstInt, secondInt, arabicOrRoman, err := checkNumbers(firstNum, secondNum)
	if err != nil {
		errorList = append(errorList, err)
	}

	return firstInt, secondInt, sign, arabicOrRoman, errorList
}

func checkSign(sign string) bool {
	if sign == "+" || sign == "-" || sign == "/" || sign == "*" {
		return false
	}
	return true
}

func checkNumbers(firstString, secondString string) (int, int, bool, error) {
	var firstInt, secondInt int
	var firstArabic, secondArabic, firstRoman, secondRoman bool

	firstArabic, firstInt = isArabicNumber(firstString)
	secondArabic, secondInt = isArabicNumber(secondString)
	if firstArabic && secondArabic {
		return firstInt, secondInt, true, nil
	}

	firstRoman, firstInt = isRomanNumber(firstString)
	secondRoman, secondInt = isRomanNumber(secondString)
	if firstRoman && secondRoman {
		return firstInt, secondInt, false, nil
	}

	return 0, 0, false, errors.New("числа должны быть либо целые арабские, либо римские, от 1 до 10 включительно")
}

func isArabicNumber(num string) (bool, int) {
	numInt, err := strconv.Atoi(num)
	if err == nil && numInt > 0 && numInt < 11 {
		return true, numInt
	}

	return false, 0
}

func isRomanNumber(num string) (bool, int) {
	romanNumberMap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	value, found := romanNumberMap[num]
	return found, value
}

func calculateArabic(firstNum, secondNum int, sign string) (result int) {
	switch sign {
	case "+":
		result = firstNum + secondNum
	case "-":
		result = firstNum - secondNum
	case "/":
		result = firstNum / secondNum
	case "*":
		result = firstNum * secondNum
	}

	return
}

func arabicToRoman(num int) (string, error) {
	if num < 1 {
		return "", errors.New("исключение: римские цифры не могут быть отрицательными")
	}

	romanNumbersMap := map[int]string{
		1:   "I",
		4:   "IV",
		5:   "V",
		9:   "IX",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}

	keys := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	var result string

	for _, key := range keys {
		for num >= key {
			result += romanNumbersMap[key]
			num -= key
		}
	}

	return result, nil
}

func main() {
	fmt.Println("Введите строку в формате \"число1_знак_число2\" и нажмите \"Enter\"")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	firstNum, secondNum, sign, arabicOrRoman, err := checkInput(input)
	if err != nil {
		for _, val := range err {
			fmt.Println(val)
		}
		return
	}

	res := calculateArabic(firstNum, secondNum, sign)
	if arabicOrRoman {
		fmt.Println(res)
	} else {
		if resRoman, err := arabicToRoman(res); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resRoman)
		}
	}
}
