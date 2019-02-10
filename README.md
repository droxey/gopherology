# gopherology

🔮 Go microservice that computes the numerological Life Path number for a given birthdate.

## Start Server

```bash
$ go run server.go
```

## Call API

```bash
$ curl -X POST http://localhost:1234/path \
  -H 'Content-Type: application/json' \
  -d '{"day":26,"month":6,"year":1988}'

22
```
