package domain


type Rectangle struct {
	X		int	`json:"x"`
	Y		int	`json:"y"`
	Width	int	`json:"width"`
	Height	int	`json:"height"`
}

func rangeCheck(val int, min int, max int) bool{
	return val>=min && val<=max
}

func (first *Rectangle) checkIntersection(second Rectangle) bool{
	xOverlap := rangeCheck(first.X, second.X, second.X+second.Width) || rangeCheck(second.X, first.X, first.X+first.Width)
	yOverlap := rangeCheck(first.Y, second.Y, second.Y + second.Height) || rangeCheck(second.Y, first.Y, first.Y+first.Height)
	return xOverlap && yOverlap
}




