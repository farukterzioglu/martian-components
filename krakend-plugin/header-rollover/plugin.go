package main

import (
	"github.com/devopsfaith/krakend-martian/register"
	"github.com/farukterzioglu/martian-components/header/rollover/modifier"
)

func init() {
	register.Set("header.RollOver", []register.Scope{register.ScopeRequest}, func(b []byte) (interface{}, error) {
		return modifier.AppendModifierFromJSON(b)
	})
}

func main() {

}
