# Using with BOSH

## Building & Creating Artifact
```
goxc
cp ~/projects/go/bin/haproxy-cli-xc/snapshot/haproxy-cli_linux_amd64.tar.gz \
  /path/to/boshrelease/blobs/haproxy-cli/haproxy-cli_linux_amd64.tar.gz
```

## Packaging:
```
mkdir -p /var/vap/store/haproxy-cli/bin
tar zxvf haproxy/haproxy-cli_linux_amd64.tar.gz \
  -C /var/vap/store/haproxy-cli/bin --strip-components 1
```

## Running

```
export HAPROXY_SOCK=/var/vcap/sys/run/haproxy/haproxy.sock

/var/vap/store/haproxy-cli/bin/haproxy-cli stats

/var/vap/store/haproxy-cli/bin/haproxy-cli info
```

