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

func TestCompute_ShouldReturn2_GivenTheInputIs1(t *testing.T) {
	input := "2"
	result := Compute(input)

	if result != "2" {
		t.Errorf("expected is 2")
	}
}

func TestCompute_ShouldReturnfoo_GivenTheInputIs3(t *testing.T) {
	input := "3"
	result := Compute(input)

	if result != "FooFoo" {
		t.Errorf("expected is FooFoo")
	}
}

func TestCompute_ShouldReturn4_GivenTheInputIs1(t *testing.T) {
	input := "4"
	result := Compute(input)

	if result != "4" {
		t.Errorf("expected is 4")
	}
}

func TestCompute_ShouldReturnBar_GivenTheInputIs5(t *testing.T) {
	input := "5"
	result := Compute(input)

	if result != "BarBar" {
		t.Errorf("expected is BarBar")
	}
}

func TestCompute_ShouldReturnQix_GivenTheInputIs7(t *testing.T) {
	input := "7"
	result := Compute(input)

	if result != "QixQix" {
		t.Errorf("Expected %v but got %v", "Qix", result)
	}
}

func TestCompute_ShouldReturnfoo_GivenTheInputIs6(t *testing.T) {
	input := "6"
	result := Compute(input)

	if result != "Foo" {
		t.Errorf("Expected %v but got %v", "Foo", result)
	}
}
