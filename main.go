package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

var data = `
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LSLS0tCg==
    server: https://ap-2.elb.amazonaws.com
  name: blah.k8s.local
`

type config struct {
	APIVersion string `yaml:"apiVersion"`
}

func main() {
	c := config{}

	if err := yaml.Unmarshal([]byte(data), &c); err != nil {
		fmt.Printf("yay: %v", err)
	}

	fmt.Printf("apiVersion: %s", c.APIVersion)
}
