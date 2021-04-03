A collection of [Martian](https://github.com/devopsfaith/krakend-martian) modifiers for [KrakenD-CE](https://github.com/devopsfaith/krakend-ce)  

- Round Robin Header Selector  
Adds a new http request header with a value that selected as round robin.  
Sample configuration;
```
"github.com/devopsfaith/krakend-martian": {
    "header.RoundRobin": {
        "name": "TRON-PRO-API-KEY",
        "values": [
            "9edf0885-8371-495f-9d59-befb9ce56885",
            "634de746-81f8-43c7-be3b-dc8eaa7d856e",
            "9e5b8bc5-9d4f-4fea-a967-e41f15f3b2b0"
        ],
        "printSelectedValue": true
    }
}
```

### How to use  
#### Add to KrakenD-CE Direclty  
Create a file inside [KrakenD-CE](https://github.com/devopsfaith/krakend-ce) as `./cmd/krakend-ce/header-roundrobin.go` with content of `header-roundrobin.go` and build KrakenD-CE (`make build`)  

#### Or build as plugin  
`go build -buildmode=plugin -o krakend-martian_header-roundrobin.so ./header/roundrobin`  

Run the KrakenD-CE  
`./krakend run -c krakend.json`

### Notes  
https://github.com/google/martian  

Website development as a sysadmin  
https://www.krakend.io/blog/website-development-as-a-sysadmin/  

How to work with Golang plugins and KrakenD  
https://www.krakend.io/blog/krakend-golang-plugins/  
