# Files

cli.go: source code for the "cli" application
cli: cli application

# Usage

On linux/mac: ./cli

# Scenario 2 

General metrics to monitor: CPU (for each core), RAM, Network bandwidth, storage
Specific metrics:
	ssl offload:
		- TLS certificate expiration
		- TLS version/ciphers used
	Reverse proxy metrics: Each of these has to be viewable per target (upstream server and URI)
		- Requests per second per target
		- Clients error metrics (4xx)
		- Server error rate (5xx)
		- request processing time (performance)
		- Dropped connections
		- Active connections
		- request length (for network sizing)
		- upstream connect time between the reverse proxy and the upstream server
		- http user agent: identify trends in used devices/browsers
		- status of the upstream servers
	Security ?
		- Not treated here, I assume there is a WAF.
How to monitor:
	- Using an agent installed on the ssl offloading server (ex: Datadog), which will send metrics to a dashboard (ex: Datadog, Grafana ...). Then alerts will have to be created on this dashboard tool to warn the SRE team in case of an issue (ex: certificate expiration).

Challenges:
	- 