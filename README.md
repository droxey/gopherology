# gopherology
[![Go Report Card](https://goreportcard.com/badge/github.com/droxey/gopherology)](https://goreportcard.com/report/github.com/droxey/gopherology) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/7ed40f9f3ecf46709879d5fbac28fd9b)](https://www.codacy.com/app/droxey/gopherology?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=droxey/gopherology&amp;utm_campaign=Badge_Grade)

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
