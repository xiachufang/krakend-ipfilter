# Example

## Run

```asciidoc
go run main.go -c krakend.json -d
```

deny from `127.0.0.0/8`

```asciidoc
$ curl -v localhost:8989/api/get-user/x1ah
* Uses proxy env variable all_proxy == 'http://127.0.0.1:7890'
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 7890 (#0)
> GET http://localhost:8989/api/get-user/x1ah HTTP/1.1
> Host: localhost:8989
> User-Agent: curl/7.64.1
> Accept: */*
> Proxy-Connection: Keep-Alive
>
< HTTP/1.1 403 Forbidden
< Connection: keep-alive
< Date: Wed, 31 Mar 2021 10:20:45 GMT
< Keep-Alive: timeout=4
< Proxy-Connection: keep-alive
< Content-Length: 0
<
* Connection #0 to host 127.0.0.1 left intact
* Closing connection 0
```

Add `127.0.0.1` to `allow` list

```asciidoc
$ curl -v localhost:8989/api/get-user/x1ah
* Uses proxy env variable all_proxy == 'http://127.0.0.1:7890'
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 7890 (#0)
> GET http://localhost:8989/api/get-user/x1ah HTTP/1.1
> Host: localhost:8989
> User-Agent: curl/7.64.1
> Accept: */*
> Proxy-Connection: Keep-Alive
>
< HTTP/1.1 200 OK
< Content-Length: 1136
< Cache-Control: public, max-age=3600
< Connection: keep-alive
< Content-Type: application/json; charset=utf-8
< Date: Wed, 31 Mar 2021 10:23:17 GMT
< Keep-Alive: timeout=4
< Proxy-Connection: keep-alive
< X-Krakend: Version undefined
< X-Krakend-Completed: true
<
* Connection #0 to host 127.0.0.1 left intact
{"avatar_url":"https://avatars.githubusercontent.com/u/14919255?v=4","bio":"打字员","blog":"","company":"@xiachufang","created_at":"2015-10-01T05:19:27Z","email":null,"events_url":"https://api.github.com/users/x1ah/events{/privacy}","followers":64,"followers_url":"https://api.github.com/users/x1ah/followers","following":60,"following_url":"https://api.github.com/users/x1ah/following{/other_user}","gists_url":"https://api.github.com/users/x1ah/gists{/gist_id}","gravatar_id":"","hireable":true,"html_url":"https://github.com/x1ah","id":14919255,"location":"BeiJing","login":"x1ah","name":"x1ah","node_id":"MDQ6VXNlcjE0OTE5MjU1","organizations_url":"https://api.github.com/users/x1ah/orgs","public_gists":5,"public_repos":60,"received_events_url":"https://api.github.com/users/x1ah/received_events","repos_url":"https://api.github.com/users/x1ah/repos","site_admin":false,"starred_url":"https://api.github.com/users/x1ah/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/x1ah/subscriptions","twitter_username":null,"type":"User","updated_at":"2021-03-29T15:10:13Z","url":"https://api.github.com/users/x1ah"}* Closing connection 0
```
