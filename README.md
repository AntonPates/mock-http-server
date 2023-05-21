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

### TODO
- [ ] Add support for different methods (POST, PUT, DELETE, etc)
- [ ] Add GitHub Actions workflow
- [ ] Make repo public
- [ ] Publish docker image on docker hub
