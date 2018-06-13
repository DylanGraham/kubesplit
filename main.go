package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBJQ0FURS0tLS0tCg==
    server: https://api-blah.ap-southeast-2.elb.amazonaws.com
  name: blah.k8s.local
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBGSUNBVEUtLS0tLQo=
    server: https://api-blah-blah.ap-southeast-2.elb.amazonaws.com
  name: blah.blah.k8s.local
contexts:
- context:
    cluster: blah.k8s.local
    user: blah.k8s.local
  name: blah.k8s.local
- context:
    cluster: blah.blah.k8s.local
    namespace: default
    user: blah.blah.k8s.local
  name: blah.blah.k8s.local
current-context: blah.k8s.local
kind: Config
preferences: {}
users:
- name: blah.k8s.local
  user:
    client-certificate-data: LS0tLS1CRjlzeDVwck9yTT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    client-key-data: LS0tLS1CRUdJTiBSUFJJVkFURSBLRVktLS0tLQo=
    password: MWt1wbazqUWFOdeXcsITuE
    username: admin
- name: blah.k8s.local-basic-auth
  user:
    password: MWt1wbazqUWFOdeXcsITuE
    username: admin
- name: blah.blah.k8s.local
  user:
    client-certificate-data: LS0tLS1CRUdwV1Y3bytMUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    client-key-data: LS0tLS1CRUdJTiBSU0EJVkFURSBLRVktLS0tLQo=
    password: xCxHeFZ4bCZ1HfPl6VsaSx
    username: admin
- name: blah.blah.k8s.local-basic-auth
  user:
    password: xCxH2TZ4bCZ1HfPl6VsaSx
    username: admin
- name: developer/127-0-0-1:8443
  user:
    token: t92DXuXe8zkHyFpFgNGkeWI
- name: system:admin/127-0-0-1:8443
  user:
    client-certificate-data: LS0tLSRPT0s1YmM9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    client-key-data: LS0tLS1CRCBSU0EgUFJJVkFURSBLRVktLS0tLQo=
`

type config struct {
	APIVersion string `yaml:"apiVersion"`
	Clusters   []struct {
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
			Server                   string `yaml:"server"`
		} `yaml:"cluster"`
		Name string `yaml:"name"`
	} `yaml:"clusters"`
	Contexts []struct {
		Context struct {
			Cluster string `yaml:"cluster"`
			User    string `yaml:"user"`
		} `yaml:"context"`
		Name string `yaml:"name"`
	} `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
	Kind           string `yaml:"kind"`
	Preferences    struct {
	} `yaml:"preferences"`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			ClientCertificateData string `yaml:"client-certificate-data"`
			ClientKeyData         string `yaml:"client-key-data"`
			Password              string `yaml:"password"`
			Username              string `yaml:"username"`
		} `yaml:"user"`
	} `yaml:"users"`
}

func main() {
	c := config{}

	if err := yaml.Unmarshal([]byte(data), &c); err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, cont := range c.Contexts {
		fmt.Printf("apiVersion: %s\n", c.APIVersion)
		fmt.Printf("kind: %s\n", c.Kind)
		fmt.Printf("current-context: %s\n", cont.Name)
		fmt.Printf("%v: %v\n", cont.Name, cont.Context)

		for _, clus := range c.Clusters {
			if clus.Name == cont.Name {
				fmt.Printf("%v: %v\n", clus.Name, clus.Cluster)
			}
		}

		for _, user := range c.Users {
			if user.Name == cont.Context.User {
				fmt.Printf("%v: %v\n", user.Name, user.User)
			}
		}

		fmt.Println()
	}
}
