package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(400, 600))
	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()
	answer := widget.NewLabel("")
	button1 := widget.NewButton("+", func() {
		number1, error1 := strconv.ParseInt(entry1.Text, 10, 64)
		number2, error2 := strconv.ParseInt(entry2.Text, 10, 64)
		if error1 != nil || error2 != nil {
			answer.SetText("")
		} else {
			answer.SetText(fmt.Sprintf("%v+%v=%v", number1, number2, number1+number2))
		}

	})
	button2 := widget.NewButton("-", func() {
		number1, error1 := strconv.ParseInt(entry1.Text, 10, 64)
		number2, error2 := strconv.ParseInt(entry2.Text, 10, 64)
		if error1 != nil || error2 != nil {
			answer.SetText("")
		} else {
			answer.SetText(fmt.Sprintf("%v-%v=%v", number1, number2, number1-number2))
		}
	})
	button3 := widget.NewButton("*", func() {
		number1, error1 := strconv.ParseInt(entry1.Text, 10, 64)
		number2, error2 := strconv.ParseInt(entry2.Text, 10, 64)
		if error1 != nil || error2 != nil {
			answer.SetText("")
		} else {
			answer.SetText(fmt.Sprintf("%v*%v=%v", number1, number2, number1*number2))
		}
	})
	button4 := widget.NewButton("/", func() {
		number1, error1 := strconv.ParseInt(entry1.Text, 10, 64)
		number2, error2 := strconv.ParseInt(entry2.Text, 10, 64)
		if error1 != nil || error2 != nil {
			answer.SetText("")
		} else {
			answer.SetText(fmt.Sprintf("%v/%v=%v", number1, number2, number1/number2))
		}
	})

	w.SetContent(container.NewVBox(entry1, entry2, button1, button2, button3, button4, answer))
	w.ShowAndRun()
}
