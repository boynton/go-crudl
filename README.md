# go-crudl
Simple example implementation of a web server defined in SADL.

## Usage

Use `make run` to generate, build, and run the server. Then in a separate terminal:

```
$ curl -sv -X POST -H "Content-type: application/json" -d '{"id":"item1","modified":"2021-02-13T00:02:04.609Z","data":"Hi there!"}' 'http://localhost:8080/items'
*   Trying ::1...
* TCP_NODELAY set
* Connection failed
* connect to ::1 port 8080 failed: Connection refused
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /items HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> Content-type: application/json
> Content-Length: 71
>
* upload completely sent off: 71 out of 71 bytes
< HTTP/1.1 201 Created
< Date: Sat, 13 Feb 2021 19:20:25 GMT
< Content-Length: 107
< Content-Type: text/plain; charset=utf-8
<
{
  "item": {
    "id": "item1",
    "modified": "2021-02-13T11:20:25.285Z",
    "data": "Hi there!"
  }
}
* Connection #0 to host localhost left intact
* Closing connection 0
$ curl -sv 'http://localhost:8080/items/item1'
*   Trying ::1...
* TCP_NODELAY set
* Connection failed
* connect to ::1 port 8080 failed: Connection refused
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /items/item1 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Last-Modified: Sat, 13 Feb 2021 19:20:25 GMT
< Date: Sat, 13 Feb 2021 19:20:52 GMT
< Content-Length: 149
< Content-Type: text/plain; charset=utf-8
<
{
  "item": {
    "id": "item1",
    "modified": "2021-02-13T11:20:25.285Z",
    "data": "Hi there!"
  },
  "modified": "2021-02-13T11:20:25.285Z"
}
* Connection #0 to host localhost left intact
* Closing connection 0
$ curl -sv 'http://localhost:8080/items/item1' -H "If-modified-since: Sat, 13 Feb 2021 19:20:25 GMT"
*   Trying ::1...
* TCP_NODELAY set
* Connection failed
* connect to ::1 port 8080 failed: Connection refused
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /items/item1 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> If-modified-since: Sat, 13 Feb 2021 19:20:25 GMT
>
< HTTP/1.1 304 Not Modified
< Date: Sat, 13 Feb 2021 19:21:23 GMT
<
* Connection #0 to host localhost left intact
* Closing connection 0
```


