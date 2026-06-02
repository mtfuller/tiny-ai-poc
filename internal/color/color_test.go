package color

import (
	"strings"
	"testing"
)

func TestColorize(t *testing.T) {
	text := "hello"
	result := Colorize(red, text)
	
	if !strings.Contains(result, text) {
		t.Errorf("Colorize() should contain text %v", text)
	}
	if !strings.Contains(result, red) {
		t.Errorf("Colorize() should contain color code")
	}
	if !strings.HasSuffix(result, reset) {
		t.Errorf("Colorize() should end with reset code")
	}
}

func TestRedColorFunction(t *testing.T) {
	text := "error"
	result := Red(text)
	
	if !strings.Contains(result, text) {
		t.Errorf("Red() should contain text %v", text)
	}
	if !strings.Contains(result, red) {
		t.Errorf("Red() should contain red color code")
	}
}

func TestGreenColorFunction(t *testing.T) {
	text := "success"
	result := Green(text)
	
	if !strings.Contains(result, text) {
		t.Errorf("Green() should contain text %v", text)
	}
	if !strings.Contains(result, green) {
		t.Errorf("Green() should contain green color code")
	}
}

func TestYellowColorFunction(t *testing.T) {
	text := "warning"
	result := Yellow(text)
	
	if !strings.Contains(result, text) {
		t.Errorf("Yellow() should contain text %v", text)
	}
	if !strings.Contains(result, yellow) {
		t.Errorf("Yellow() should contain yellow color code")
	}
}

func TestBlueColorFunction(t *testing.T) {
	text := "info"
	result := Blue(text)
	
	if !strings.Contains(result, text) {
		t.Errorf("Blue() should contain text %v", text)
	}
	if !strings.Contains(result, blue) {
		t.Errorf("Blue() should contain blue color code")
	}
}

func TestCyanColorFunction(t *testing.T) {
	text := "cyan"
	result := Cyan(text)
	
	if !strings.Contains(result, text) {
		t.Errorf("Cyan() should contain text %v", text)
	}
	if !strings.Contains(result, cyan) {
		t.Errorf("Cyan() should contain cyan color code")
	}
}

func TestMagentaColorFunction(t *testing.T) {
	text := "magenta"
	result := Magenta(text)
	
	if !strings.Contains(result, text) {
		t.Errorf("Magenta() should contain text %v", text)
	}
	if !strings.Contains(result, magenta) {
		t.Errorf("Magenta() should contain magenta color code")
	}
}

func TestWhiteColorFunction(t *testing.T) {
	text := "white"
	result := White(text)
	
	if !strings.Contains(result, text) {
		t.Errorf("White() should contain text %v", text)
	}
	if !strings.Contains(result, white) {
		t.Errorf("White() should contain white color code")
	}
}

func TestBoldFunction(t *testing.T) {
	text := "bold"
	result := Bold(text)
	
	if !strings.Contains(result, text) {
		t.Errorf("Bold() should contain text %v", text)
	}
	if !strings.Contains(result, bold) {
		t.Errorf("Bold() should contain bold code")
	}
}
