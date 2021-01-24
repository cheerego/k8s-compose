package compose

import (
	"github.com/cockroachdb/errors"
	"os"
)

type Collector struct {
}

func (c *Collector) Workspace() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "os.Getwd error")
	}

	return pwd, nil
}

func CollectEnvs() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "os.Getwd error")
	}

	return pwd, nil
}

func CollectEnvs() (*map[string]string, error) {

}
