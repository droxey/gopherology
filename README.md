# gopherology

🔮 Go microservice that computes the numerological Life Path number for a given birthdate.

## Usage

### POST `/api/path`: Life Path Number

```bash
$ curl -X POST http://gopherology.herokuapp.com/api/path \
  -H 'Content-Type: application/json' \
  -d '{"day":26,"month":6,"year":1988}'
```

### POST `/api/path`: Output

```json
{
  "year": 1988,
  "month": 6,
  "day": 26,
  "path": 22,
  "message": "Your Life Path Number is 22"
}
```

## Development

```bash
$ export PORT=1324; go run main.go
```
