package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const y_for_consent_message = "press \"y\" or \"yes\" to proceed: "


func PromptInt(msg string) (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(msg)
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

    int, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		return 0, err
	}

    return int, nil
}

func PromptString(msg string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(msg)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

    return strings.TrimSpace(input), nil
}

func PromptConsent(msg string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
    consent := false

	fmt.Printf("%s%s", msg, y_for_consent_message)
	input, err := reader.ReadString('\n')

	if err != nil {
		return consent, err
	}

    input = strings.ToLower(strings.TrimSpace(input))

    if input == "y" || input == "yes" {
        consent = true
    }

    return consent, nil
}
