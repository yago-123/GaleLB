# GaleLB: multi-node L4 load balancer

Supports: 
- [ ] L3-Based Forwarding (Stateless IP Routing in `XDP`)
- [ ] L4-Based Forwarding (Stateful `NAT` with Connection Tracking in `XDP + TC`)]

## Requirements 

## Architecture

## Configuration
Load balancer configuration:
```toml
[node_health]

# number of continuous health checks that must be passed before being eligible for routing destination
checks_before_routing = 3
# duration of deadline between health checks, after this period, nodes will be removed from the routing ring
checks_timeout = "5s"

# number of times nodes can fail to send health checks before they are blacklisted
# ex: the node will be added and removed 5 times to the routing table before they will start to be completly ignored.
# use -1 if want to disable this option
black_list_after_fails = 5

# duration of the ban 
black_list_expiry = "5m"
```

Node configuration:
```toml
[load_balancer]
addresses = [
    { ip = "192.168.1.1", port = 7070 },
    { ip = "192.168.1.2", port = 8080 },
    { ip = "192.168.1.3", port = 9090 }
]
```

## Example

To run the Gale load balancer, make sure you have the necessary permissions. Elevated privileges are required to adjust the `rlimit` and load the `XDP` module. You can run the load balancer with the following command:

```bash
$ sudo ./bin/gale-lb --config cmd/lb.toml
```


## Dependencies 
Install dependencies for building eBPF programs: 
```bash
$ sudo apt-get install clang llvm libbpf-dev gcc make
$ sudo apt install linux-headers-$(uname -r)
```

Install linter: 
```bash
$ go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5
```

Install `protobuf` and `protoc`: 
```bash 
$ go get google.golang.org/protobuf/cmd/protoc-gen-go
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
