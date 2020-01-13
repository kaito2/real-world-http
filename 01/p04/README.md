## Run

```
$ go run 01/p04/main.go
```

## curl with http1.0

```
$ curl --http1.0 http://localhost:18888/greeting
<html><body>Hello</body></html>
```

Server log

```
$ go run 01/p04/main.go
2020/01/13 15:13:38 start http listening :18888
GET /greeting HTTP/1.0
Host: localhost:18888
Connection: close
Accept: */*
User-Agent: curl/7.54.0
```

