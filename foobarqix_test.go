package main

//gofmt => format the file =>remove all unused inporst and intendadion
import "testing"

func TestCompute_ShouldReturn1_GivenTheInputIs1(t *testing.T) {
	input := "1"
	result := Compute(input)

	if result != "1" {
		t.Errorf("expected is 1")
	}
}

func TestCompute_ShouldReturn2_GivenTheInputIs2(t *testing.T) {
	input := "2"
	result := Compute(input)

	if result != "2" {
		t.Errorf("expected is 2")
	}
}

func TestCompute_ShouldReturnfoofoo_GivenTheInputIs3(t *testing.T) {
	input := "3"
	result := Compute(input)

	if result != "FooFoo" {
		t.Errorf("expected is FooFoo")
	}
}

func TestCompute_ShouldReturn4_GivenTheInputIs4(t *testing.T) {
	input := "4"
	result := Compute(input)

	if result != "4" {
		t.Errorf("expected is 4")
	}
}

func TestCompute_ShouldReturnBarBar_GivenTheInputIs5(t *testing.T) {
	input := "5"
	result := Compute(input)

	if result != "BarBar" {
		t.Errorf("expected is BarBar")
	}
}

func TestCompute_ShouldReturnQixQix_GivenTheInputIs7(t *testing.T) {
	input := "7"
	result := Compute(input)

	if result != "QixQix" {
		t.Errorf("Expected %v but got %v", "QixQix", result)
	}
}

func TestCompute_ShouldReturnfoo_GivenTheInputIs6(t *testing.T) {
	input := "6"
	result := Compute(input)

	if result != "Foo" {
		t.Errorf("Expected %v but got %v", "Foo", result)
	}
}

func TestCompute_ShouldReturn8_GivenTheInputIs8(t *testing.T) {
	input := "8"
	result := Compute(input)

	if result != "8" {
		t.Errorf("Expected %v but got %v", "8", result)
	}
}

func TestCompute_ShouldReturnfoo_GivenTheInputIs9(t *testing.T) {
	input := "9"
	result := Compute(input)

	if result != "Foo" {
		t.Errorf("Expected %v but got %v", "Foo", result)
	}
}

func TestCompute_ShouldReturn10_GivenTheInputIs10(t *testing.T) {
	input := "10"
	result := Compute(input)

	if result != "Bar" {
		t.Errorf("Expected %v but got %v", "Bar", result)
	}
}

func TestCompute_ShouldReturn13_GivenTheInputIs13(t *testing.T) {
	input := "13"
	result := Compute(input)

	if result != "Foo" {
		t.Errorf("Expected %v but got %v", "Foo", result)
	}
}

func TestCompute_ShouldReturnfoobarbar_GivenTheInputIs15(t *testing.T) {
	input := "15"
	result := Compute(input)

	if result != "FooBarBar" {
		t.Errorf("Expected %v but got %v", "FooBarBar", result)
	}
}

func TestCompute_ShouldReturnfooqix_GivenTheInputIs21(t *testing.T) {
	input := "21"
	result := Compute(input)

	if result != "FooQix" {
		t.Errorf("Expected %v but got %v", "FooQix", result)
	}
}

func TestCompute_ShouldReturnfoofoofoo_GivenTheInputIs33(t *testing.T) {
	input := "33"
	result := Compute(input)

	if result != "FooFooFoo" {
		t.Errorf("Expected %v but got %v", "FooFooFoo", result)
	}
}

func TestCompute_ShouldReturnfoobar_GivenTheInputIs51(t *testing.T) {
	input := "51"
	result := Compute(input)

	if result != "FooBar" {
		t.Errorf("Expected %v but got %v", "FooBar", result)
	}
}

func TestCompute_ShouldReturnbarfoo_GivenTheInputIs53(t *testing.T) {
	input := "53"
	result := Compute(input)

	if result != "BarFoo" {
		t.Errorf("Expected %v but got %v", "BarFoo", result)
	}
}
