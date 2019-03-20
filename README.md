## Playground and example structure for getting started with Go Template to build Webserver or API Servers 

This playground is based on
* [Echo Go Web Framework](https://echo.labstack.com/)
* [Golang Template](https://golang.org/pkg/text/template/)

# The main idea is to build/playground robust http routers with go templates

# **Still work in progress**

# Requirements 
* Golang > 1.12

# Installation
* this project is installed with use of the GOlang module Tool, for more information , please follow this link [Golang Modules](https://github.com/golang/go/wiki/Modules)

```bash
##no use of GOPATH and act as a project outside of the GOPATH/src
export GO111MODULE=on
##run the programm
go run main.go
##build the binary
go build .
#run ober the binary
./webserver-go-templates
```

# Installation on minishift / openshift
go to the [README-minishift](./minishift/README.md)