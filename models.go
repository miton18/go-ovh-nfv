package go_ovh_nfv

type(
	Region struct {
		Region string `json:"region"`
	}
	Regions []Region

	Loadbalancer struct {
		ID            string                    `json:"id"`
		Description   string                    `json:"description"`
		Name          string                    `json:"name"`
		Region        string                    `json:"region"`
		Address       Address                   `json:"address"`
		EgressAddress EgressAddress             `json:"egressAddress"`
		Configuration LoadbalancerConfiguration `json:"configuration"`
		Status        LoadBalancerStatus        `json:"status"`
	}
	Loadbalancers []Loadbalancer

	Address struct {
		IPv4 string `json:"ipv4"`
		IPv6 string `json:"ipv6,omitempty"`
	}

	EgressAddress struct {
		IPv4 []string `json:"ipv4"`
		IPv6 []string `json:"ipv6,omitempty"`
	}

	LoadbalancerConfiguration struct {
		Applied int64 `json:"applied"`
		Latest int64 `json:"latest"`
	}

	LoadbalancerCreate struct {
		Region string `json:"region"`
		Name string `json:"name"`
		Description string `json:"description,omitempty"`
	}

	Configuration struct {
		Frontends Frontends `json:"frontends"`
		Backends Backends `json:"backends"`
		Version int64 `json:"version"`
	}
	Configurations []Configuration

	Frontends map[string]Frontend
	Frontend  struct {
		Mode       FrontendMode `json:"mode,omitempty" validate:"omitempty,oneof=http tcp udp" default:"tcp"`
		Port       int64              `json:"port,omitempty" validate:"required,min=1,max=65535"`
		Whitelists []string           `json:"whiteList" validate:"omitempty,dive,cidr"`
		Backend    BackendsSelectors  `json:"backend" validate:"required,dive"`
	}

	Backends map[string]Backend
	Backend  struct {
		Balancer      BalancerAlgorithm `json:"balancer,omitempty" validate:"omitempty,oneof=roundrobin static-rr leastconn first source" default:"roundrobin"`
		ProxyProtocol ProxyProtocol     `json:"proxyProtocol,omitempty" validate:"omitempty,oneof=v1 v2 v2-ssl v2-cn"`
		Servers       Servers                 `json:"servers,omitempty" validate:"required,gt=0"`
	}

	BackendsSelectors []BackendSelector
	BackendSelector   struct {
		Name string `json:"name,omitempty" validate:"required"`
	}

	Servers []Server
	Server  struct {
		Name    string `json:"name,omitempty" validate:"omitempty,max=100"`
		IP      string `json:"ip,omitempty" validate:"required,ip"`
		Port    int64  `json:"port,omitempty" validate:"omitempty,min=1,max=65535"`
		NoCheck bool   `json:"noCheck,omitempty" validate:"omitempty"`
		Weight  int64  `json:"weight,omitempty" validate:"omitempty,min=0,max=256" default:"100"`
	}
)
