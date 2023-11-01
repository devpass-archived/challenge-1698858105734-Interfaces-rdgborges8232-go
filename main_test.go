package main

import (
	"testing"
)

func TestShape(t *testing.T) {
	rectangle := Rectangle{Width: 5, Height: 8}
	circle := Circle{Radius: 3}

	expectedRectangleArea := 40.000000
	expectedRectanglePerimeter := 26.000000
	expectedCircleArea := 28.260000
	expectedCirclePerimeter := 18.840000

	if rectangle.Area() != expectedRectangleArea {
		t.Errorf("Expected rectangle area to be %f, but got %f", expectedRectangleArea, rectangle.Area())
	}

	if rectangle.Perimeter() != expectedRectanglePerimeter {
		t.Errorf("Expected rectangle perimeter to be %f, but got %f", expectedRectanglePerimeter, rectangle.Perimeter())
	}

	if circle.Area() != expectedCircleArea {
		t.Errorf("Expected circle area to be %f, but got %f", expectedCircleArea, circle.Area())
	}

	if circle.Perimeter() != expectedCirclePerimeter {
		t.Errorf("Expected circle perimeter to be %f, but got %f", expectedCirclePerimeter, circle.Perimeter())
	}
}