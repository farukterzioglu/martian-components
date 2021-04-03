package main

import (
	"github.com/devopsfaith/krakend-martian/register"
	"github.com/farukterzioglu/martian-components/header/roundrobin/modifier"
)

func init() {
	register.Set("header.RoundRobin", []register.Scope{register.ScopeRequest}, func(b []byte) (interface{}, error) {
		return modifier.FromJSON(b)
	})
}

func main() {

}
