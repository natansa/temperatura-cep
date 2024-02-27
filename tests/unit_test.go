package tests

import (
	"testing"

	"github.com/natansa/temperatura-cep/services"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	result := services.CelsiusToFahrenheit(22)
	if result != 71.6 {
		t.Errorf("Expected 71.6, got %f", result)
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	result := services.CelsiusToKelvin(20.1)
	if result != 293.25 {
		t.Errorf("Expected 293.25, got %f", result)
	}
}
