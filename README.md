# gopherology

🔮 Go microservice that computes the numerological Life Path number for a given birthdate.

## Usage

### `/api/path`

#### Request

```bash
curl -X POST https://gopherology.herokuapp.com/api/path \
  -H 'Content-Type: application/json' \
  -d '{"day":26,"month":6,"year":1988}'
```

#### Response

```json
{
  "number": 22,
  "more_info":" https://www.tokenrock.com/numerology/my_life_path/?num=22"
}
```

## Development

```bash
$ export PORT=1324; go run main.go
```
