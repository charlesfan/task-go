# task-go
Simple task API application.
## Run service
1. Clone this code:
```sh
$ git clone https://github.com/charlesfan/task-go.git
```
2. Run with docker-compos:
```sh
$ docker compose up --build -d
```
## Test
```sh
$ make test
```
## REST API
### Crearte task information
#### Request
`POST /samdbox/tasks`

    curl -i -H 'Accept: application/json' -X POST -d 'name=task-01&status=0' http://localhost:8080/sandbox/tasks
    
#### Response
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Thu, 03 Apr 2025 10:49:38 GMT
    Content-Length: 94

    {"code":20000000,"msg":"success","data":{"id":576605225058570240,"name":"task-01","status":0}}

### Get tasks list
#### Request
`GET /samdbox/tasks`

    curl -i -H 'Accept: application/json' -X GET http://localhost:8080/sandbox/tasks

#### Response
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Thu, 03 Apr 2025 10:56:40 GMT
    Content-Length: 150

    {"code":20000000,"msg":"success","data":[{"id":576605225058570240,"name":"task-01","status":0},{"id":576604996888432640,"name":"task-02","status":0}]}

### Put task
#### Request
`PUT /samdbox/tasks/{id}`

    curl -i -H 'Accept: application/json' -X PUT -d 'name=task-02&status=1' http://localhost:8080/sandbox/tasks/576604996888432640

#### Response
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Thu, 03 Apr 2025 11:02:15 GMT
    Content-Length: 94

    {"code":20000000,"msg":"success","data":{"id":576604996888432640,"name":"task-02","status":1}}

### Delete task
#### Request
`DELETE /samdbox/tasks/{id}`

    curl -i -H 'Accept: application/json' -X DELETE http://localhost:8080/sandbox/tasks/576604996888432640

#### Response
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Thu, 03 Apr 2025 11:05:02 GMT
    Content-Length: 33

    {"code":20000000,"msg":"success"}
