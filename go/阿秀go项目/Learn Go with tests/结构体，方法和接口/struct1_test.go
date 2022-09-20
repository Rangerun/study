package struct1

import "testing"

/*func Perimeter(rectangle Rectangle) float64 {
    return 2 * (rectangle.width + rectangle.height)
}

func Area(rectangle Rectangle) float64 {
    return rectangle.width * rectangle.height
}

func TestPerimeter(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    got := Perimeter(rectangle)
    want := 40.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}

func TestArea(t *testing.T) {
    rectangle := Rectangle{12.0, 6.0}
    got := Area(rectangle)
    want := 72.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}*/


type Rectangle struct{
	Width	float64
	Height  float64
}

type Circle struct {
	Radius float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (circle Circle) Area() float64 {
	return 314.159265358979
}

type Shape interface{
	Area() float64
}

type Triangle struct {
	Base   float64
    Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}

// 我们可以通过如下命令来运行列表中指定的测试用例： go test -run TestArea/Rectangle
func TestArea(t *testing.T) {

    areaTests := []struct {
        name    string
        shape   Shape
        hasArea float64
    }{
        {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
        {name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
        {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
    }

    for _, tt := range areaTests {
        // using tt.name from the case to use it as the `t.Run` test name
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
            }
        })

    }

}