package pkg

import (
	"fmt"
	"os"
	"strings"
)

func ParseSpecificFlagString(flagName string) (string, error) {
	prefix := "-"
	if strings.Contains(flagName, "-") {
		prefix = "--"
	}

	for i, arg := range os.Args {
		if !strings.HasPrefix(arg, prefix) {
			continue
		}

		if arg == fmt.Sprintf("%s%s", prefix, flagName) {
			return os.Args[i+1], nil
		} else if strings.Contains(arg, fmt.Sprintf("%s=", flagName)) {
			return strings.Replace(arg, fmt.Sprintf("%s%s=", prefix, flagName), "", 1), nil
		}
	}

	return "", fmt.Errorf("failed to parse %s flag", flagName)
}
