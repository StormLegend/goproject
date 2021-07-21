package perimeter

import "math"

type Rectangle struct{
    Width float64
    Height float64
}

type Circle struct{
    Radius float64
}

type Triangle struct{
    Lenth float64
    Height float64
}

func (t Triangle)Area()float64{
    return t.Lenth*t.Height/2
}

func Perimeter(rectangle Rectangle)float64{
    return (rectangle.Width+rectangle.Height)*2
}


func (r Rectangle) Area()float64{
    return r.Height * r.Width
}

func (c Circle) Area()float64{
	return math.Pi * c.Radius*c.Radius
}
