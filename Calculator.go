package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
	"unicode"
)

// нет проверки на длину числа, можно выполнить только 1 операцию
func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(400, 600))
	answer := widget.NewLabel("")
	buttonOne := widget.NewButton("1", func() {
		buf := answer.Text
		if buf == "" {
			answer.SetText("1")
		} else if strings.Contains(buf, "=") {
			answer.SetText("1")
		} else if !unicode.IsDigit(rune(buf[len(buf)-1])) {
			answer.SetText(buf + strconv.Itoa(1))
		} else {
			number, err := strconv.ParseInt(buf, 10, 64)
			if err != nil {
				for i := range buf {
					if !unicode.IsDigit(rune(buf[i])) {
						switch buf[i] {
						case '+':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "+"), 10, 64)
						case '-':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "-"), 10, 64)
						case '*':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "*"), 10, 64)
						case '/':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "/"), 10, 64)
						}
						break
					}
				}
				leftPart := strings.TrimRight(buf, "+-*/")
				answer.SetText(fmt.Sprintf("%v%v", leftPart, number*10+1))
			} else {
				answer.SetText(fmt.Sprintf("%v", number*10+1))
			}
		}

	})
	buttonTwo := widget.NewButton("2", func() {
		buf := answer.Text
		if buf == "" {
			answer.SetText("2")
		} else if strings.Contains(buf, "=") {
			answer.SetText("2")
		} else if !unicode.IsDigit(rune(buf[len(buf)-1])) {
			answer.SetText(buf + strconv.Itoa(2))
		} else {
			number, err := strconv.ParseInt(buf, 10, 64)
			if err != nil {
				for i := range buf {
					if !unicode.IsDigit(rune(buf[i])) {
						switch buf[i] {
						case '+':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "+"), 10, 64)
						case '-':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "-"), 10, 64)
						case '*':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "*"), 10, 64)
						case '/':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "/"), 10, 64)
						}
						break
					}
				}
				leftPart := strings.TrimRight(buf, "+-*/")
				answer.SetText(fmt.Sprintf("%v%v", leftPart, number*10+2))
			} else {
				answer.SetText(fmt.Sprintf("%v", number*10+2))
			}
		}

	})
	buttonThree := widget.NewButton("3", func() {
		buf := answer.Text
		if buf == "" {
			answer.SetText("3")
		} else if strings.Contains(buf, "=") {
			answer.SetText("3")
		} else if !unicode.IsDigit(rune(buf[len(buf)-1])) {
			answer.SetText(buf + strconv.Itoa(3))
		} else {
			number, err := strconv.ParseInt(buf, 10, 64)
			if err != nil {
				for i := range buf {
					if !unicode.IsDigit(rune(buf[i])) {
						switch buf[i] {
						case '+':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "+"), 10, 64)
						case '-':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "-"), 10, 64)
						case '*':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "*"), 10, 64)
						case '/':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "/"), 10, 64)
						}
						break
					}
				}
				leftPart := strings.TrimRight(buf, "+-*/")
				answer.SetText(fmt.Sprintf("%v%v", leftPart, number*10+3))
			} else {
				answer.SetText(fmt.Sprintf("%v", number*10+3))
			}
		}

	})
	buttonFour := widget.NewButton("4", func() {
		buf := answer.Text
		if buf == "" {
			answer.SetText("4")
		} else if strings.Contains(buf, "=") {
			answer.SetText("4")
		} else if !unicode.IsDigit(rune(buf[len(buf)-1])) {
			answer.SetText(buf + strconv.Itoa(4))
		} else {
			number, err := strconv.ParseInt(buf, 10, 64)
			if err != nil {
				for i := range buf {
					if !unicode.IsDigit(rune(buf[i])) {
						switch buf[i] {
						case '+':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "+"), 10, 64)
						case '-':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "-"), 10, 64)
						case '*':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "*"), 10, 64)
						case '/':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "/"), 10, 64)
						}
						break
					}
				}
				leftPart := strings.TrimRight(buf, "+-*/")
				answer.SetText(fmt.Sprintf("%v%v", leftPart, number*10+4))
			} else {
				answer.SetText(fmt.Sprintf("%v", number*10+4))
			}
		}

	})
	buttonFive := widget.NewButton("5", func() {
		buf := answer.Text
		if buf == "" {
			answer.SetText("5")
		} else if strings.Contains(buf, "=") {
			answer.SetText("5")
		} else if !unicode.IsDigit(rune(buf[len(buf)-1])) {
			answer.SetText(buf + strconv.Itoa(5))
		} else {
			number, err := strconv.ParseInt(buf, 10, 64)
			if err != nil {
				for i := range buf {
					if !unicode.IsDigit(rune(buf[i])) {
						switch buf[i] {
						case '+':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "+"), 10, 64)
						case '-':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "-"), 10, 64)
						case '*':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "*"), 10, 64)
						case '/':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "/"), 10, 64)
						}
						break
					}
				}
				leftPart := strings.TrimRight(buf, "+-*/")
				answer.SetText(fmt.Sprintf("%v%v", leftPart, number*10+5))
			} else {
				answer.SetText(fmt.Sprintf("%v", number*10+5))
			}
		}

	})
	buttonSix := widget.NewButton("6", func() {
		buf := answer.Text
		if buf == "" {
			answer.SetText("6")
		} else if strings.Contains(buf, "=") {
			answer.SetText("6")
		} else if !unicode.IsDigit(rune(buf[len(buf)-1])) {
			answer.SetText(buf + strconv.Itoa(6))
		} else {
			number, err := strconv.ParseInt(buf, 10, 64)
			if err != nil {
				for i := range buf {
					if !unicode.IsDigit(rune(buf[i])) {
						switch buf[i] {
						case '+':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "+"), 10, 64)
						case '-':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "-"), 10, 64)
						case '*':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "*"), 10, 64)
						case '/':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "/"), 10, 64)
						}
						break
					}
				}
				leftPart := strings.TrimRight(buf, "+-*/")
				answer.SetText(fmt.Sprintf("%v%v", leftPart, number*10+6))
			} else {
				answer.SetText(fmt.Sprintf("%v", number*10+6))
			}
		}

	})
	buttonSeven := widget.NewButton("7", func() {
		buf := answer.Text
		if buf == "" {
			answer.SetText("7")
		} else if strings.Contains(buf, "=") {
			answer.SetText("7")
		} else if !unicode.IsDigit(rune(buf[len(buf)-1])) {
			answer.SetText(buf + strconv.Itoa(7))
		} else {
			number, err := strconv.ParseInt(buf, 10, 64)
			if err != nil {
				for i := range buf {
					if !unicode.IsDigit(rune(buf[i])) {
						switch buf[i] {
						case '+':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "+"), 10, 64)
						case '-':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "-"), 10, 64)
						case '*':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "*"), 10, 64)
						case '/':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "/"), 10, 64)
						}
						break
					}
				}
				leftPart := strings.TrimRight(buf, "+-*/")
				answer.SetText(fmt.Sprintf("%v%v", leftPart, number*10+7))
			} else {
				answer.SetText(fmt.Sprintf("%v", number*10+7))
			}
		}

	})
	buttonEight := widget.NewButton("8", func() {
		buf := answer.Text
		if buf == "" {
			answer.SetText("8")
		} else if strings.Contains(buf, "=") {
			answer.SetText("8")
		} else if !unicode.IsDigit(rune(buf[len(buf)-1])) {
			answer.SetText(buf + strconv.Itoa(8))
		} else {
			number, err := strconv.ParseInt(buf, 10, 64)
			if err != nil {
				for i := range buf {
					if !unicode.IsDigit(rune(buf[i])) {
						switch buf[i] {
						case '+':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "+"), 10, 64)
						case '-':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "-"), 10, 64)
						case '*':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "*"), 10, 64)
						case '/':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "/"), 10, 64)
						}
						break
					}
				}
				leftPart := strings.TrimRight(buf, "+-*/")
				answer.SetText(fmt.Sprintf("%v%v", leftPart, number*10+8))
			} else {
				answer.SetText(fmt.Sprintf("%v", number*10+8))
			}
		}

	})
	buttonNine := widget.NewButton("9", func() {
		buf := answer.Text
		if buf == "" {
			answer.SetText("9")
		} else if strings.Contains(buf, "=") {
			answer.SetText("9")
		} else if !unicode.IsDigit(rune(buf[len(buf)-1])) {
			answer.SetText(buf + strconv.Itoa(9))
		} else {
			number, err := strconv.ParseInt(buf, 10, 64)
			if err != nil {
				for i := range buf {
					if !unicode.IsDigit(rune(buf[i])) {
						switch buf[i] {
						case '+':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "+"), 10, 64)
						case '-':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "-"), 10, 64)
						case '*':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "*"), 10, 64)
						case '/':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "/"), 10, 64)
						}
						break
					}
				}
				leftPart := strings.TrimRight(buf, "+-*/")
				answer.SetText(fmt.Sprintf("%v%v", leftPart, number*10+9))
			} else {
				answer.SetText(fmt.Sprintf("%v", number*10+9))
			}
		}

	})
	buttonZero := widget.NewButton("0", func() {
		buf := answer.Text
		if buf == "" {
			answer.SetText("0")
		} else if strings.Contains(buf, "=") {
			answer.SetText("0")
		} else if !unicode.IsDigit(rune(buf[len(buf)-1])) {
			answer.SetText(buf + strconv.Itoa(0))
		} else {
			number, err := strconv.ParseInt(buf, 10, 64)
			if err != nil {
				for i := range buf {
					if !unicode.IsDigit(rune(buf[i])) {
						switch buf[i] {
						case '+':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "+"), 10, 64)
						case '-':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "-"), 10, 64)
						case '*':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "*"), 10, 64)
						case '/':
							number, _ = strconv.ParseInt(strings.TrimLeft(buf, "/"), 10, 64)
						}
						break
					}
				}
				leftPart := strings.TrimRight(buf, "+-*/")
				answer.SetText(fmt.Sprintf("%v%v", leftPart, number*10))
			} else {
				answer.SetText(fmt.Sprintf("%v", number*10))
			}
		}

	})
	buttonAdding := widget.NewButton("+", func() {
		answer.SetText(answer.Text + "+")
	})
	buttonSubtracting := widget.NewButton("-", func() {
		answer.SetText(answer.Text + "-")
	})
	buttonMultiplication := widget.NewButton("*", func() {
		answer.SetText(answer.Text + "*")
	})
	buttonDivision := widget.NewButton("/", func() {
		answer.SetText(answer.Text + "/")
	})
	buttonRemoveAll := widget.NewButton("del", func() {
		answer.SetText("")
	})
	buttonResult := widget.NewButton("=", func() {
		buf := answer.Text
		for i := range buf {
			if !unicode.IsDigit(rune(buf[i])) {
				number1, _ := strconv.ParseInt(buf[0:i], 10, 64)
				number2, _ := strconv.ParseInt(buf[i+1:], 10, 64)
				switch buf[i] {
				case '+':
					answer.SetText(fmt.Sprintf("%v=%v", buf, number1+number2))
				case '-':
					answer.SetText(fmt.Sprintf("%v=%v", buf, number1-number2))
				case '*':
					answer.SetText(fmt.Sprintf("%v=%v", buf, number1*number2))
				case '/':
					answer.SetText(fmt.Sprintf("%v=%v", buf, number1/number2))
				default:
					break
				}
				break
			}
		}
	})

	w.SetContent(container.NewGridWithRows(5,
		container.NewGridWithColumns(1, answer),
		container.NewGridWithColumns(4, buttonSeven, buttonEight, buttonNine, buttonAdding),
		container.NewGridWithColumns(4, buttonFour, buttonFive, buttonSix, buttonSubtracting),
		container.NewGridWithColumns(4, buttonOne, buttonTwo, buttonThree, buttonMultiplication),
		container.NewGridWithColumns(4, buttonRemoveAll, buttonZero, buttonDivision, buttonResult),
	))
	w.ShowAndRun()
}
