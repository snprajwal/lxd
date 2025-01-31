package osarch

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// GetLSBRelease returns a map with Linux distribution information.
func GetLSBRelease() (map[string]string, error) {
	osRelease, err := getLSBRelease("/etc/os-release")
	if os.IsNotExist(err) {
		return getLSBRelease("/usr/lib/os-release")
	}

	return osRelease, err
}

func getLSBRelease(filename string) (map[string]string, error) {
	osRelease := make(map[string]string)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return osRelease, nil
		}

		return osRelease, err
	}

	for i, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		tokens := strings.SplitN(line, "=", 2)
		if len(tokens) != 2 {
			return osRelease, fmt.Errorf("%s: invalid format on line %d", filename, i+1)
		}

		osRelease[tokens[0]] = strings.Trim(tokens[1], `'"`)
	}

	return osRelease, nil
}
