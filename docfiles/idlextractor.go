package docfiles

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// APIs parses and returns all contained *.bs API definitions
func APIs(dir string) ([]*API, error) {
	var res []*API
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to list %s: %e", dir, err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".bs") {
			text, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return res, fmt.Errorf("failed to read bikeshed file: %e", err)
			}
			apis, err := parseBikeshed(text)
			if err != nil {
				return res, fmt.Errorf("failed to parse bikeshed file: %e", err)
			}
			for _, api := range apis {
				res = append(res, api)
			}
		}
	}
	return res, nil
}
