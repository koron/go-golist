package golist

import (
	"encoding/json"
	"os"
	"os/exec"
)

// GetPackage gets information of package by `go list`.
func GetPackage(name string) (*Package, error) {
	println("name="+ name)
	if name != "" && name != "." {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		err = os.Chdir(name)
		if err != nil {
			return nil, err
		}
		defer os.Chdir(cwd)
	}
	pkg, err := list(".")
	if err != nil {
		return nil, err
	}
	return pkg, nil
}

func list(name string) (*Package, error) {
	c := exec.Command("go", "list", "-e", "-json", name)
	b, err := c.Output()
	if err != nil {
		return nil, err
	}
	pkg := new(Package)
	err = json.Unmarshal(b, pkg)
	if err != nil {
		return nil, err
	}
	return pkg, nil
}
