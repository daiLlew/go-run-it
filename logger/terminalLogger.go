package logger

import "fmt"

const MAGENTA = "\033[35m"
const RED = "\033[31m"
const NO_COL = "\033[0m"
const BLUE = "\033[96m"

const GREEN = "\033[92m"
const YELLOW = "\033[93m"

const LIGHT_BLUE = "\033[96m"
const LIGHT_GREEN = "\033[92m"
const LIGHT_YELLOW = "\033[93m"

const START_CMD_MSG_FMT = "%s[start-up]%s %s %s\n"

const SHUT_DOWN_MSG_FMT = "%s[shutdown]%s %s %s\n"
const ERROR_MSG_FMT = "%s[ERROR]%s %s %s\n"


func ShutdownDebug(msg string) {
	fmt.Printf(fmt.Sprintf(SHUT_DOWN_MSG_FMT, YELLOW, BLUE, msg, NO_COL))
}

func StartupDebug(msg string) {
	fmt.Printf(fmt.Sprintf(START_CMD_MSG_FMT, YELLOW, BLUE, msg, NO_COL))
}

func LogError(msg string) {
	fmt.Printf(fmt.Sprintf(ERROR_MSG_FMT, RED, MAGENTA, msg, NO_COL))
}

