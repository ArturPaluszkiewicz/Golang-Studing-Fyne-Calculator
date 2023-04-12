package calculator

type Calculator struct {
	FirstNumber float64
	SecondNumber float64
	Result float64
}

func (cal *Calculator) Add(){
	cal.Result = cal.FirstNumber + cal.SecondNumber
}

func (cal *Calculator) Minus(){
	cal.Result = cal.FirstNumber - cal.SecondNumber
}

func (cal *Calculator) Divine(){
	cal.Result = cal.FirstNumber / cal.SecondNumber
}

func (cal *Calculator) Multiple(){
	cal.Result = cal.FirstNumber * cal.SecondNumber
}