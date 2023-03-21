package log

import (
	"fmt"
	"os"
)

func Log(a ...any) {
	fmt.Print("[LOG] ")
	fmt.Println(a...)
}

// Info
/**
 * Console val with info level into stdout
 * @param val the printing string
 * @return void
 * @example [INFO] tea console example
 */
func Info(a ...any) {
	fmt.Print("[INFO] ")
	fmt.Println(a...)
}

// Warning
/**
 * Console val with warning level into stdout
 * @param val the printing string
 * @return void
 * @example [WARNING] tea console example
 */
func Warning(a ...any) {
	fmt.Print("[WARNING] ")
	fmt.Println(a...)
}

// Debug
/**
 * Console val with debug level into stdout
 * @param val the printing string
 * @return void
 * @example [DEBUG] tea console example
 */
func Debug(a ...any) {
	fmt.Printf("[DEBUG] ")
	fmt.Println(a...)
}

// Error
/**
 * Console val with error level into stderr
 * @param val the printing string
 * @return void
 * @example [ERROR] tea console example
 */
func Error(a ...any) {
	_, _ = fmt.Fprint(os.Stderr, "[ERROR] ")
	_, _ = fmt.Fprintln(os.Stderr, a...)
}
