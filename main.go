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

type Cluster struct {
	CertificateAuthorityData string `yaml:"certificate-authority-data"`
	Server                   string `yaml:"server"`
}

type Clusters struct {
	Name    string  `yaml:"name"`
	Cluster Cluster `yaml:"cluster"`
}

type Context struct {
	Cluster string `yaml:"cluster"`
	User    string `yaml:"user"`
}

type Contexts struct {
	Name    string  `yaml:"name"`
	Context Context `yaml:"context"`
}

type User struct {
	ClientCertificateData string `yaml:"client-certificate-data"`
	ClientKeyData         string `yaml:"client-key-data"`
	Password              string `yaml:"password"`
	Username              string `yaml:"username"`
}

type Users struct {
	Name string `yaml:"name"`
	User User   `yaml:"user"`
}

type c2 struct {
	APIVersion     string     `yaml:"apiVersion"`
	Clusters       []Clusters `yaml:"clusters"`
	Contexts       []Contexts `yaml:"contexts"`
	CurrentContext string     `yaml:"current-context"`
	Kind           string     `yaml:"kind"`
	Preferences    struct{}   `yaml:"preferences"`
	Users          []Users    `yaml:"users"`
}

type c1 struct {
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
	c := c2{}
	out := c2{}

	if err := yaml.Unmarshal([]byte(data), &c); err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, cont := range c.Contexts {
		out.APIVersion = c.APIVersion
		out.Kind = c.Kind
		out.CurrentContext = cont.Name
		out.Contexts = []Contexts{Contexts{Name: cont.Name}}

		for _, clus := range c.Clusters {
			if clus.Name == cont.Name {
				out.Clusters = []Clusters{Clusters{cont.Name, Cluster{clus.Cluster.CertificateAuthorityData, clus.Cluster.Server}}}
			}
		}

		for _, user := range c.Users {
			if user.Name == cont.Context.User {
				out.Users = []Users{Users{Name: user.Name}}
				out.Users[0].User.ClientCertificateData = user.User.ClientCertificateData
				out.Users[0].User.ClientKeyData = user.User.ClientKeyData
				out.Users[0].User.Password = user.User.Password
				out.Users[0].User.Username = user.User.Username
			}
		}

		out, err := yaml.Marshal(&out)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("%v\n", string(out))
	}
}
