Add to KrakenD-CE Direclty  
Create a file inside [KrakenD-CE](https://github.com/devopsfaith/krakend-ce) as `./cmd/krakend-ce/header-roller.go` with content of `header-roller.go` and build KrakenD-CE  

Or build as plugin  
`go build -buildmode=plugin -o krakend-martian_header-ro.so ./header/rollover`  

`./krakend run -c krakend.json`

### Notes  
https://github.com/google/martian  

Website development as a sysadmin  
https://www.krakend.io/blog/website-development-as-a-sysadmin/  

How to work with Golang plugins and KrakenD  
https://www.krakend.io/blog/krakend-golang-plugins/  
