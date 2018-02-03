package main
import (
        _ "github.com/coredns/example"
        "github.com/coredns/coredns/coremain"
        _ "github.com/coredns/coredns/plugin/auto"
	_ "github.com/coredns/coredns/plugin/autopath"
	_ "github.com/coredns/coredns/plugin/bind"
	_ "github.com/coredns/coredns/plugin/cache"
	_ "github.com/coredns/coredns/plugin/chaos"
	_ "github.com/coredns/coredns/plugin/debug"
	_ "github.com/coredns/coredns/plugin/dnssec"
	_ "github.com/coredns/coredns/plugin/dnstap"
	_ "github.com/coredns/coredns/plugin/erratic"
	_ "github.com/coredns/coredns/plugin/errors"
	_ "github.com/coredns/coredns/plugin/etcd"
	_ "github.com/coredns/coredns/plugin/federation"
	_ "github.com/coredns/coredns/plugin/file"
	_ "github.com/coredns/coredns/plugin/health"
	_ "github.com/coredns/coredns/plugin/hosts"
	_ "github.com/coredns/coredns/plugin/kubernetes"
	_ "github.com/coredns/coredns/plugin/loadbalance"
	_ "github.com/coredns/coredns/plugin/log"
	_ "github.com/coredns/coredns/plugin/metrics"
	_ "github.com/coredns/coredns/plugin/nsid"
	_ "github.com/coredns/coredns/plugin/pprof"
	_ "github.com/coredns/coredns/plugin/proxy"
	_ "github.com/coredns/coredns/plugin/reload"
	_ "github.com/coredns/coredns/plugin/reverse"
	_ "github.com/coredns/coredns/plugin/rewrite"
	_ "github.com/coredns/coredns/plugin/root"
	_ "github.com/coredns/coredns/plugin/route53"
	_ "github.com/coredns/coredns/plugin/secondary"
	_ "github.com/coredns/coredns/plugin/template"
	_ "github.com/coredns/coredns/plugin/tls"
	_ "github.com/coredns/coredns/plugin/trace"
	_ "github.com/coredns/coredns/plugin/whoami"
	_ "github.com/coredns/example"
	_ "github.com/coredns/forward"
	_ "github.com/mholt/caddy/onevent"
        _ "github.com/mholt/caddy/startupshutdown"
        _ "github.com/arvancloud/redis"
        "github.com/coredns/coredns/core/dnsserver"
)
var directives = []string{
        "example",
        "tls",
	"reload",
	"nsid",
	"root",
	"bind",
	"debug",
	"trace",
	"health",
	"pprof",
	"prometheus",
	"errors",
	"log",
	"dnstap",
	"chaos",
	"loadbalance",
	"cache",
	"rewrite",
	"dnssec",
	"autopath",
	"reverse",
	"template",
	"hosts",
	"route53",
	"federation",
	"kubernetes",
	"file",
	"auto",
        "secondary",
        "etcd",
        "redis",
	"forward",
	"proxy",
	"erratic",
	"whoami",
	"on",
	"startup",
	"shutdown",
}

func init() {
        dnsserver.Directives = directives
}

func main() {
        coremain.Run()
}
