# gopherology
![Go Report Card](https://goreportcard.com/badge/github.com/droxey/gopherology)

🔮 Go microservice that recursively computes a numerological Life Path number for a given birthdate.

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
  "pathNumber": 22,
  "detailsUrl": "https://www.tokenrock.com/numerology/my_life_path/?num=22",
  "isMasterNumber": true
}
```

## Development

```bash
$ export PORT=1324; go run main.go
```
