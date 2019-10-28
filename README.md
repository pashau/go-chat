# go-chat
simple web server with chat functionality
based on "Go Programming Blueprints - Second Edition" by Mat Ryer (https://www.packtpub.com/application-development/go-programming-blueprints-second-edition)


## Notes:
- `go run main.go`
- `sudo netstat -ltnp | grep :808`
  - 1915/apache2
  - `sudo service apache2 status`
  - http://localhost:8080/
  - `ll /var/www/html/`
  - `sudo service apache2 stop`
  - `sudo service apache2 start`
- build & execute:
  - `go build -o {name}`
  - `./{name}`
- get external packages:
  - `go get` | `go get github.com/gorilla/websocket` (dependency/package => $GOPATH | ~/go/src/github.com/gorilla/websocket)
