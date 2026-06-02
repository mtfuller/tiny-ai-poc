package example

import (
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple name", "Alice", "Hello, Alice!"},
		{"another name", "Bob", "Hello, Bob!"},
		{"empty string", "", "Hello, !"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Greet(tt.input)
			if result != tt.expected {
				t.Errorf("Greet(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCalculateAdd(t *testing.T) {
	result, err := Calculate(5, 3, "add")
	if err != nil {
		t.Errorf("Calculate(5, 3, add) returned error: %v", err)
	}
	if result != 8 {
		t.Errorf("Calculate(5, 3, add) = %v, want 8", result)
	}
}

func TestCalculateSubtract(t *testing.T) {
	result, err := Calculate(10, 4, "subtract")
	if err != nil {
		t.Errorf("Calculate(10, 4, subtract) returned error: %v", err)
	}
	if result != 6 {
		t.Errorf("Calculate(10, 4, subtract) = %v, want 6", result)
	}
}

func TestCalculateMultiply(t *testing.T) {
	result, err := Calculate(6, 7, "multiply")
	if err != nil {
		t.Errorf("Calculate(6, 7, multiply) returned error: %v", err)
	}
	if result != 42 {
		t.Errorf("Calculate(6, 7, multiply) = %v, want 42", result)
	}
}

func TestCalculateDivide(t *testing.T) {
	result, err := Calculate(20, 4, "divide")
	if err != nil {
		t.Errorf("Calculate(20, 4, divide) returned error: %v", err)
	}
	if result != 5 {
		t.Errorf("Calculate(20, 4, divide) = %v, want 5", result)
	}
}

func TestCalculateDivideByZero(t *testing.T) {
	_, err := Calculate(10, 0, "divide")
	if err == nil {
		t.Error("Calculate(10, 0, divide) should return error for division by zero")
	}
}

func TestCalculateInvalidOperation(t *testing.T) {
	_, err := Calculate(5, 3, "invalid")
	if err == nil {
		t.Error("Calculate with invalid operation should return error")
	}
}

func TestCalculateAllOperations(t *testing.T) {
	tests := []struct {
		name      string
		a         int
		b         int
		operation string
		expected  int
		wantErr   bool
	}{
		{"add positive", 5, 3, "add", 8, false},
		{"add negative", -5, 3, "add", -2, false},
		{"subtract", 10, 4, "subtract", 6, false},
		{"multiply", 6, 7, "multiply", 42, false},
		{"divide", 20, 4, "divide", 5, false},
		{"divide by zero", 10, 0, "divide", 0, true},
		{"unknown operation", 5, 3, "modulo", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculate(tt.a, tt.b, tt.operation)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Calculate() expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Calculate() unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Calculate() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}
