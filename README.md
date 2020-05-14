# Sample-api 
## Using [chi](https://github.com/go-chi/chi) router 

```shell script
$ go get github.com/dorokhin/sample-api
$ $GOPATH/bin/sample-api

# Go to http://localhost:8080/api/users
```

```shell script
curl -i http://localhost:8080/api/users
# HTTP/1.1 200 OK
# Content-Type: application/json
# Content-Length: 170
```

```json
[
  {"id":1,"first_name":"Andrew","last_name":"Dorokhin","email":"andrew@dorokhin.moscow"},
  {"id":2,"first_name":"Vasily","last_name":"Gorev","email":"vasya_gor@email.com"}
]
```