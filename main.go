package main

import (
	calc "Calculator/calculator"
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}


func main(){
	myCalc := calc.Calculator {}
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")

	f1 := binding.BindFloat(&myCalc.FirstNumber)
	str1 := binding.FloatToString(f1)
	f2 := binding.BindFloat(&myCalc.SecondNumber)
	str2 := binding.FloatToString(f2)
	f3 := binding.BindFloat(&myCalc.Result)
	str3 := binding.FloatToString(f3)

	
	input1 := widget.NewEntryWithData(str1)
	input2 := widget.NewEntryWithData(str2)
	inputPanel := container.New(layout.NewGridLayout(2), input1, input2)

	addButton := widget.NewButton("Dodaj", func() {
		myCalc.Add()
		f3.Set(myCalc.Result)
	})
	minusButton := widget.NewButton("Odejmij", func() {
		myCalc.Minus()
		f3.Set(myCalc.Result)
	})
	divButton := widget.NewButton("Podziel", func() {
		myCalc.Divine()
		f3.Set(myCalc.Result)
	})
	multiButton := widget.NewButton("Pomnoz", func() {
		myCalc.Multiple()
		f3.Set(myCalc.Result)
	})
	checkPanel := container.New(layout.NewGridLayout(4), addButton, minusButton, divButton, multiButton)

	resultPanel := widget.NewLabelWithData(str3)

	animationPanel := canvas.NewRectangle(color.Black)
	animationPanel.Resize(fyne.NewSize(100,500))

	mainPanel := container.New(layout.NewGridLayout(1), inputPanel, resultPanel, checkPanel, animationPanel)

	myWindow.SetContent(mainPanel)
	myWindow.Resize(fyne.NewSize(500,500))
	myWindow.ShowAndRun()

	fmt.Println(myCalc.FirstNumber)
	fmt.Println(myCalc.SecondNumber)
	fmt.Println(myCalc.Result)
}
