package logger

import "fmt"

const MAGENTA = "\033[35m"
const NO_COL = "\033[0m"
const BLUE = "\033[96m"

const GREEN = "\033[92m"
const YELLOW = "\033[93m"

const LIGHT_BLUE = "\033[96m"
const LIGHT_GREEN = "\033[92m"
const LIGHT_YELLOW = "\033[93m"

const START_CMD_MSG_FMT = "%s[start-up]%s %s %s\n"

const SHUT_DOWN_MSG_FMT = "%s[shutdown]%s %s %s\n"


func ShutdownDebug(msg string) {
	fmt.Printf(formatShutdownMessage(msg))
}

func StartupDebug(msg string) {
	fmt.Printf(formatStartupMessage(msg))
}

func formatShutdownMessage(msg string) string {
	return fmt.Sprintf(SHUT_DOWN_MSG_FMT, YELLOW, BLUE, msg, NO_COL)
}

func formatStartupMessage(msg string) string {
	return fmt.Sprintf(START_CMD_MSG_FMT, GREEN, YELLOW, msg, NO_COL)
}

