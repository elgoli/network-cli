# network-cli
This is an implementation of network Command Line Interface in Go
## Commands
### ipLookup
To find host IP address using the local resolver, you can run the network-cli service locally with `ipLookup` command and `--host` switch as following:
```
network-cli ipLookup --host <hostName>
```
### mxLookup
To find DNS MX records for a given domain name, you can run the network-cli service locally with `mxLookup` command and `--host` switch as following:
```
network-cli mxLookup --host <hostName>
```
### nsLookup
To find DNS NS records for a given domain name, you can run the network-cli service locally with `nsLookup` command and `--host` switch as following:
```
network-cli nsLookup --host <hostName>
```
## How to access network-cli on Linux distributions
To access the network-cli commands in system wide and for all users, network-cli bin file can be created and added to the path as follwoing:
```
git clone https://gitlab.com/elgol/network-cli.git
cd network-cli
make build
cp bin/network-cli /usr/local/bin
```
