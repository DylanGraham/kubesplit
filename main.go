/*
kubesplit - split a merged kubernetes config file
Copyright (C) 2018 Dylan Graham

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

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

type config struct {
	APIVersion     string     `yaml:"apiVersion"`
	Clusters       []Clusters `yaml:"clusters"`
	Contexts       []Contexts `yaml:"contexts"`
	CurrentContext string     `yaml:"current-context"`
	Kind           string     `yaml:"kind"`
	Preferences    struct{}   `yaml:"preferences"`
	Users          []Users    `yaml:"users"`
}

func init() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s %s\n\n", os.Args[0], "<config>")
		os.Exit(-1)
	}
}

func readFile(b []byte) []byte {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("couldn't read file: %v", err)
	}
	return b
}

func main() {
	var bytes []byte
	bytes = readFile(bytes)

	c := config{}
	out := config{}

	if err := yaml.Unmarshal([]byte(bytes), &c); err != nil {
		log.Fatalf("error: %v", err)
	}

	for i, cont := range c.Contexts {
		out.APIVersion = c.APIVersion
		out.Kind = c.Kind
		out.CurrentContext = cont.Name
		out.Contexts = []Contexts{Contexts{cont.Name, Context{cont.Context.Cluster, cont.Context.User}}}

		for _, clus := range c.Clusters {
			if clus.Name == cont.Name {
				out.Clusters = []Clusters{Clusters{cont.Name, Cluster{clus.Cluster.CertificateAuthorityData, clus.Cluster.Server}}}
			}
		}

		for _, user := range c.Users {
			if user.Name == cont.Context.User {
				out.Users = []Users{Users{user.Name, User{user.User.ClientCertificateData, user.User.ClientKeyData, user.User.Password, user.User.Username}}}
			}
		}

		out, err := yaml.Marshal(&out)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		filename := fmt.Sprintf("config-split-%d", i)
		filename = cont.Name + ".yaml"
		if err := ioutil.WriteFile(filename, out, 0600); err != nil {
			log.Fatalf("error: %v", err)
		}
	}
}
