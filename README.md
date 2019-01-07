# gopherology

🔮 Go microservice that computes the numerological Life Path number for a given birthdate.

## Start Server

```bash
$ go run server.go
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v3.3.dev
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:9999
```

## Call API

```bash
$ curl -X POST http://localhost:9999/path \
  -H 'Content-Type: application/json' \
  -d '{"day":26,"month":6,"year":1988}'
```
