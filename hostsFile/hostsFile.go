package hostsFile

import (
	"bytes"
	"os"
	"os/user"
)

func WriteHostsFile(domains []string) error {
	file, err := os.OpenFile("/etc/hosts", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	data := bytes.NewBufferString("\n")
	for _, line := range domains {
		_, err = data.WriteString("127.0.0.1 " + line + " # Auto generated by local proxy\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func IsPrivileged() (bool, error) {
	u, err := user.Current()
	if err != nil {
		return false, err
	}
	return u.Uid == "0", nil
}
