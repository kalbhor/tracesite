# tracesite

*tracesite is a simple Go implementation of the traceroute tool*

#### Check out the [blog post](https://blog.kalbhor.xyz/post/implementing-traceroute-in-go/) on explanation

## Install : 
- Download binary from [releases](https://github.com/kalbhor/tracesite/releases)
- Build from source : `go get -v github.com/kalbhor/tracesite`

## Usage :

```
NAME:
   tracesite - trace the route to a site

USAGE:
   Trace a site: tracesite --hop=3 --timeout=2000 kalbhor.xyz

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --hop value         start from a custom hop number (default: 1)
   --maxhops value     custom max hops (default: 24)
   --port value        custom port number (default: 33434)
   --timeout value     custom timeout in ms (default: 4000)
   --retries value     custom retries (default: 3)
   --packetsize value  custom packet size (default: 52)
   --help, -h          show help
```
