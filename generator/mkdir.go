package generator

import (
	"fmt"
	"os"
)

func mkdir(path string, o *Option) error {
	dir, err := buildPath(path, o)

	if err != nil {
		return err
	}

	info, err := os.Stat(dir)

	if err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(dir, 0755)
		}

		return err
	}

	if info.IsDir() {
		return nil
	}

	return fmt.Errorf("[%s] already exists, but it's not directory", dir)
}
