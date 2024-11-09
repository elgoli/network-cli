# network-cli
This is a cli tool to lookup DNS A, AAA, NS and MX records.
## Commands
### iplookup
To find host IP address using the local resolver:
```
$ ncli iplookup --host <hostName>
```
### mxlookup
To find DNS MX records for a given domain name:
```
$ ncli mxlookup --host <hostName>
```
### nslookup
To find DNS NS records for a given domain name:
```
$ ncli nslookup --host <hostName>
```
## How to access ncli on Linux distributions
To access the commands in system wide and for all users, `ncli` bin file can be created and added to the path as following:
```
$ git clone https://github.com/elnazdev/network-cli.git
$ cd network-cli
$ make build
$ cp bin/ncli /usr/local/bin
```
