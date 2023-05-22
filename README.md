![CI](https://https://github.com/AntonPates/mock-http-server/.github/workflows/action.yml/badge.svg)



# Mock http server

## Usage
```bash
./build/bin/mock-server --help
Usage of ./build/bin/mock-server:
  -addr string
        address to listen (default ":8080")
  -config string
        path to config file (default "config.json")
```

## Config file
```json
[
    {
        "method": "GET",
        "path": "/healthcheck",
        "body": "OK 200",
        "status_code": 200,
        "headers": {
            "Content-Type": "text/plain"
        }
    },
    {
        "method": "GET",
        "path": "/api/v1/json",
        "body": {
            "key": "value"
        },
        "status_code": 200,
        "headers": {
            "Content-Type": "application/json"
        }
    }
]
```

## Examples

```bash
curl -i localhost:8080/healthcheck
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Sun, 21 May 2023 20:44:00 GMT
Content-Length: 6

OK 200
```

```bash
curl -i localhost:8080/api/v1/json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 21 May 2023 20:17:50 GMT
Content-Length: 16

{"key":"value"}
```

### TODO
- [ ] Add support for different methods (POST, PUT, DELETE, etc)
- [ ] Add GitHub Actions workflow
- [ ] Make repo public
- [ ] Publish docker image on docker hub
