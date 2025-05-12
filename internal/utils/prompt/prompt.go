package prompt

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

var terminalSupported = false

func init() {
	terminalSupported = CheckTerminalSupported()
}

// CheckTerminalSupported returns true if the current terminal supports
// line editing features.
func CheckTerminalSupported() bool {
	bad := map[string]bool{"": true, "dumb": true, "cons25": true}

	return !bad[strings.ToLower(os.Getenv("TERM"))]
}

// Password prompts the user for a password. Set confirmation to true
// to require the user to confirm the password.
func Password(label string, confirmation bool) string {
	prompt := promptui.Prompt{
		Label:   label,
		Mask:    '*',
		Pointer: promptui.PipeCursor,
	}
	password, err := prompt.Run()
	FatalErrorCheck(err)

	if confirmation {
		validate := func(input string) error {
			if input != password {
				return errors.New("passwords do not match")
			}

			return nil
		}

		confirmPrompt := promptui.Prompt{
			Label:    "Confirm password",
			Validate: validate,
			Mask:     '*',
			Pointer:  promptui.PipeCursor,
		}

		_, err := confirmPrompt.Run()
		FatalErrorCheck(err)
	}

	return password
}

// Confirm prompts user to confirm the operation.
func Confirm(label string) bool {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
		Pointer:   promptui.PipeCursor,
	}
	result, err := prompt.Run()
	if err != nil {
		if !errors.Is(err, promptui.ErrAbort) {
			PrintErrorMsgf("prompt error: %v", err)
		} else {
			PrintWarnMsgf("Aborted.")
		}
		os.Exit(1)
	}

	if result != "" && strings.ToUpper(result[:1]) == "Y" {
		return true
	}

	return false
}

// Input prompts for an input string.
func Input(label string) string {
	prompt := promptui.Prompt{
		Label:   label,
		Pointer: promptui.PipeCursor,
	}
	result, err := prompt.Run()
	FatalErrorCheck(err)

	return result
}

// Select prompts create choice menu for select by user.
func Select(label string, items []string) int {
	prompt := promptui.Select{
		Label:   label,
		Items:   items,
		Pointer: promptui.PipeCursor,
	}

	choice, _, err := prompt.Run()
	FatalErrorCheck(err)

	return choice
}

// InputWithSuggestion prompts the user for an input string with a suggestion.
func InputWithSuggestion(label, suggestion string) string {
	prompt := promptui.Prompt{
		Label:   label,
		Default: suggestion,
		Pointer: promptui.PipeCursor,
	}
	result, err := prompt.Run()
	FatalErrorCheck(err)

	return result
}

// InputWithRange prompts the user for an input integer within a specified range.
func InputWithRange(label string, def, minimum, maximum int) int {
	prompt := promptui.Prompt{
		Label:     label,
		Default:   fmt.Sprintf("%v", def),
		IsVimMode: true,
		Pointer:   promptui.PipeCursor,
		Validate: func(input string) error {
			num, err := strconv.Atoi(input)
			if err != nil {
				return err
			}
			if num < minimum || num > maximum {
				return fmt.Errorf("enter a number between %v and %v", minimum, maximum)
			}

			return nil
		},
	}
	result, err := prompt.Run()
	FatalErrorCheck(err)

	num, err := strconv.Atoi(result)
	FatalErrorCheck(err)

	return num
}

func PrintErrorMsgf(format string, args ...any) {
	format = "[ERROR] " + format
	if terminalSupported {
		// Print error msg with red color
		format = fmt.Sprintf("\033[31m%s\033[0m", format)
	}
	fmt.Printf(format+"\n", args...)
}

func PrintSuccessMsgf(format string, a ...any) {
	if terminalSupported {
		// Print successful msg with green color
		format = fmt.Sprintf("\033[32m%s\033[0m", format)
	}
	fmt.Printf(format+"\n", a...)
}

func PrintWarnMsgf(format string, a ...any) {
	if terminalSupported {
		// Print warning msg with yellow color
		format = fmt.Sprintf("\033[33m%s\033[0m", format)
	}
	fmt.Printf(format+"\n", a...)
}

func PrintInfoMsgf(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

func PrintInfoMsgBoldf(format string, a ...any) {
	if terminalSupported {
		format = fmt.Sprintf("\033[1m%s\033[0m", format)
	}
	fmt.Printf(format+"\n", a...)
}

func PrintLine() {
	fmt.Println()
}

func FatalErrorCheck(err error) {
	if err != nil {
		if terminalSupported {
			fmt.Printf("\033[31m%s\033[0m\n", err.Error())
		} else {
			fmt.Printf("%s\n", err.Error())
		}

		os.Exit(1)
	}
}
