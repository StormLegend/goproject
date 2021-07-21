package perimeter

import "testing"


/*type Rectangle struct {
    Width float64
    Height float64
}
*/


type Shape interface{
    Area() float64
}

func TestPerimeter(t *testing.T){
    checkPerimeter := func(t testing.TB, got, want float64){
	t.Helper()
	if got != want {
	    t.Errorf("want %f but got %f",want,got)
	}
    }
    t.Run("ordinary calculate perimeter",func(t *testing.T){
	rectangle := Rectangle{10.0,10.0}
        got := Perimeter(rectangle)
	want := 40.0
	checkPerimeter(t,got,want)
    })
}

func TestAreaNew(t *testing.T){
    areaTests := []struct{
        shape Shape
	want float64
    }{
        {Rectangle{12.0,6.0}, 72.0},
	{Circle{10.0}, 314.1592653589793},
	{Triangle{12.0,6.0}, 36.0},
    }

    for _, tt := range areaTests{
        if tt.shape.Area() != tt.want{
	    t.Errorf("%#v got %g but want %g",tt.shape, tt.shape.Area(), tt.want)
	}
    }
}

func TestArea(t *testing.T){
    checkArea := func(t testing.TB, shape Shape, want float64){
        t.Helper()
	if shape.Area() != want{
            t.Errorf("want %f but got %f",want,shape.Area())
	}
    }
    t.Run("rectangles area calculate", func(t *testing.T){
        rectangle := Rectangle{12.0,6.0}
        want := 72.0
	checkArea(t, rectangle, want)
    })
    t.Run("circles calculate", func(t *testing.T){
        circle := Circle{10.0}
	want := 314.1592653589793
	checkArea(t, circle, want)
    })
}
