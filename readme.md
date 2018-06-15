# kubesplit

kubesplit is a simple tool to split merged kubernetes config files.

## Install

`go get github.com/DylanGraham/kubesplit`

## Usage

`kubesplit <config>`

## Example

```
$ kubesplit config
$ ls
config config-split-1 config-split-2
```

## Why?

Tools have varying levels of support for multiple kubernetes config files and the KUBECONFIG environment variable.

I find myself in the situation of having some clean separate files, such as `config-gke` and `config-local`, but then also a merged default file at `config` which gets data added from tools such as `kops` upon cluster initialisation.

This tool aims to keep cluster config neatly organised into one file per config, allowing you flexibility to specify multiple files with KUBECONFIG or specify your chosen config when required, eg:

`KUBECONFIG=~/.kube/config-x istioctl get virtualservices`

## Recommended tools
https://github.com/ahmetb/kubectx

Fast way to switch between clusters and namespaces in kubectl!


## License

GPLv3
```
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
```
