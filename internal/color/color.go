package color

import "fmt"

// ANSI color codes
const (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
	bold    = "\033[1m"
)

// Color functions for easy colored output
func Colorize(color, text string) string {
	return color + text + reset
}

// Red returns text in red color
func Red(text string) string {
	return Colorize(red, text)
}

// Green returns text in green color
func Green(text string) string {
	return Colorize(green, text)
}

// Yellow returns text in yellow color
func Yellow(text string) string {
	return Colorize(yellow, text)
}

// Blue returns text in blue color
func Blue(text string) string {
	return Colorize(blue, text)
}

// Magenta returns text in magenta color
func Magenta(text string) string {
	return Colorize(magenta, text)
}

// Cyan returns text in cyan color
func Cyan(text string) string {
	return Colorize(cyan, text)
}

// White returns text in white color
func White(text string) string {
	return Colorize(white, text)
}

// Bold returns text in bold
func Bold(text string) string {
	return Colorize(bold, text)
}

// Success prints a success message in green
func Success(format string, args ...interface{}) {
	fmt.Printf(Green("✓ ")+format+"\n", args...)
}

// Error prints an error message in red
func Error(format string, args ...interface{}) {
	fmt.Printf(Red("✗ ")+format+"\n", args...)
}

// Warning prints a warning message in yellow
func Warning(format string, args ...interface{}) {
	fmt.Printf(Yellow("⚠ ")+format+"\n", args...)
}

// Info prints an info message in blue
func Info(format string, args ...interface{}) {
	fmt.Printf(Blue("ℹ ")+format+"\n", args...)
}
