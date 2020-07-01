package go_ovh_nfv

type BalancerAlgorithm string

const (
	RoundRobin      BalancerAlgorithm = "roundrobin"
	StaticRR        BalancerAlgorithm = "static-rr"
	LeastConnection BalancerAlgorithm = "leastconn"
	First           BalancerAlgorithm = "first"
	Source          BalancerAlgorithm = "source"
)

type FrontendMode string

const (
	TCP FrontendMode = "tcp"
	HTTP FrontendMode = "http"
)

type (
	LoadBalancerStatus    string
)

const (
	Created LoadBalancerStatus =  "CREATED"
Applying LoadBalancerStatus = "APPLYING"
Running LoadBalancerStatus = "RUNNING"
Deleting LoadBalancerStatus = "DELETING"
Error LoadBalancerStatus = "ERROR"
Frozen LoadBalancerStatus = "FROZEN"
)

type ProxyProtocol string

const (
	None  ProxyProtocol = ""
	V1    ProxyProtocol = "v1"
	V2    ProxyProtocol = "v2"
	V2SSL ProxyProtocol = "v2-ssl"
	V2CN  ProxyProtocol = "v2-cn"
)
