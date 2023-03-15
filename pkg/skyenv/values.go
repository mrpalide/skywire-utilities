package skyenv

// Constants for new default services.
const (
	ServiceConfAddr     = "http://conf.skywire.skycoin.com"
	TpDiscAddr          = "http://tpd.skywire.skycoin.com"
	DmsgDiscAddr        = "http://dmsgd.skywire.skycoin.com"
	ServiceDiscAddr     = "http://sd.skycoin.com"
	RouteFinderAddr     = "http://rf.skywire.skycoin.com"
	UptimeTrackerAddr   = "http://ut.skywire.skycoin.com"
	AddressResolverAddr = "http://ar.skywire.skycoin.com"
	SetupPK             = "0324579f003e6b4048bae2def4365e634d8e0e3054a20fc7af49daf2a179658557"
	NetworkMonitorPK    = "0380ea88f0ad0aa4d93c330ba5f97aabca1d892190b94db69eee140b549d2817dd"
)

// Constants for testing deployment.
const (
	TestServiceConfAddr     = "http://conf.skywire.dev"
	TestTpDiscAddr          = "http://tpd.skywire.dev"
	TestDmsgDiscAddr        = "http://dmsgd.skywire.dev"
	TestServiceDiscAddr     = "http://sd.skywire.dev"
	TestRouteFinderAddr     = "http://rf.skywire.dev"
	TestUptimeTrackerAddr   = "http://ut.skywire.dev"
	TestAddressResolverAddr = "http://ar.skywire.dev"
	TestSetupPK             = "026c2a3e92d6253c5abd71a42628db6fca9dd9aa037ab6f4e3a31108558dfd87cf"
	TestNetworkMonitorPK    = "0218905f5d9079bab0b62985a05bd162623b193e948e17e7b719133f2c60b92093"
)

// GetStunServers gives back default Stun Servers
func GetStunServers() []string {
	return []string{
		"139.162.12.30:3478",
		"170.187.228.181:3478",
		"172.104.161.184:3478",
		"170.187.231.137:3478",
		"143.42.74.91:3478",
		"170.187.225.78:3478",
		"143.42.78.123:3478",
		"139.162.12.244:3478",
	}
}

// DNSServer is value for DNS Server Address
const DNSServer = "1.1.1.1"
