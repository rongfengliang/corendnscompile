# coredns cutomer plugin build demo project

>  include default Corefile with build plugin

- add example plugin
- add redis plugin

## init priject
```bash
go get github.com/coredns/coredns
```
### add  plugin  [example]
```bash
go get github.com/coredns/example
```
### add plugin  [redis]
```bash
go get github.com/arvancloud/redis
```
## build project
```bash
go build
```
## run 
```bash
./main 
```


