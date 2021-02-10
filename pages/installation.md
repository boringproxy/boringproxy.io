# The program is shipped in one executable file that acts both as the server and the client.

Acquire it as below for your system:

## Linux amd64
```bash
curl -LO https://github.com/boringproxy/boringproxy/releases/download/v0.5.0/boringproxy-linux-amd64d

### Make executable
chmod +x boringproxy-linux-amd64

### Allow binding to ports 80 and 443
sudo setcap cap_net_bind_service=+ep boringproxy-linux-amd64
```

## Linux i386
```bash
curl -LO https://github.com/boringproxy/boringproxy/releases/download/v0.5.0/boringproxy-linux-386

### Make executable
chmod +x boringproxy-linux-386

### Allow binding to ports 80 and 443
sudo setcap cap_net_bind_service=+ep boringproxy-linux-386
```
##Linux ARM
```bash
curl -LO https://github.com/boringproxy/boringproxy/releases/download/v0.5.0/boringproxy-linux-arm

### Make executable
chmod +x boringproxy-linux-arm

### Allow binding to ports 80 and 443
sudo setcap cap_net_bind_service=+ep boringproxy-linux-arm
```
##Linux arm64
```bash
curl -LO https://github.com/boringproxy/boringproxy/releases/download/v0.5.0/boringproxy-linux-arm64

### Make executable
chmod +x boringproxy-linux-arm64

### Allow binding to ports 80 and 443
sudo setcap cap_net_bind_service=+ep boringproxy-linux-arm64
```
## Windows and macOS
See the <a href="https://github.com/boringproxy/boringproxy/releases">releases page</a> for downloads.
Note: macOS support is untested. 

