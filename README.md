# host-ip-helper
a service for expose IPs of host

# build

+ on windows
  ```CMD
  @REM for windows
  go build -o host-ip-helper.exe main.go

  @REM for other platform
  SET CGO_ENABLED=0 SET GOOS=darwin SET GOARCH=amd64 go build -o host-ip-helper main.go
  SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build -o host-ip-helper main.go
  ```

+ on linux
  ```bash
  # for linux
  go build -o host-ip-helper main.go
  # for other platform
  CGO_ENABLED=0 SET GOOS=windows SET GOARCH=amd64 go build -o host-ip-helper.exe main.go
  CGO_ENABLED=0 SET GOOS=darwin SET GOARCH=amd64 go build -o host-ip-helper main.go
  ```

+ on macOS
  ```bash
  # for macOS
  go build -o host-ip-helper main.go
  # for other platform
  CGO_ENABLED=0 SET GOOS=windows SET GOARCH=amd64 go build -o host-ip-helper.exe main.go
  CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build -o host-ip-helper main.go
  ```

# install

+ for windows `.\install.cmd`
+ for *nix `./install.sh`
