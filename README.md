# HAProxy CLI Utility & Library

## HAProxy Configuration

First be sure to enable the haproxy UDS stats socket in the config then reload:
```
global
  stats socket /tmp/haproxy.sock mode 600 level admin
  stats socket ipv4@127.0.0.1:9000 level admin
  stats timeout 2m
```

## CLI Usage
Once haproxy is configured and reloaded, collect haproxy stats and/or info as follows:

```
export HAPROXY_SOCK=/tmp/haproxy.sock
haproxy-cli stats
haproxy-cli info
```

For using haproxy-cli with BOSH releases see `docs/bosh.md`

## Library Usage

Be sure that the `HAPROXY_SOCK` environment variable is set and pointing to the 
location of the HAProxy socket file as described above.

Import the library, if you set the import name to `hap` then you can name your
object that you interact with `haproxy`.
```
import (
	hap "github.com/wayneeseguin/haproxy-cli/haproxy"
)
```
Create a new haproxy object:
```
haproxy = &hap.Haproxy{Socket: os.Getenv("HAPROXY_SOCK")}
```

Then you can gather statistics as a JSON array of statistics K/V objects,
```
output, err := haproxy.Stats("all")
```

Or information about the haproxy process as a JSON KV Object,
```
output, err := haproxy.Info()
```
