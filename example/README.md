# Example

## Run

```shell
go run main.go -c krakend.json -d
```

deny from `127.0.0.0/8`

```shell
$ curl -v localhost:8989/api/get-user/x1ah -H 'X-Real-Ip: 127.0.0.1'
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8989 (#0)
> GET /api/get-user/x1ah HTTP/1.1
> Host: localhost:8989
> User-Agent: curl/7.64.1
> Accept: */*
> X-Real-Ip: 127.0.0.1
>
< HTTP/1.1 403 Forbidden
< Date: Thu, 24 Jun 2021 10:27:09 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
* Closing connection 0

```

Add `127.0.0.1` to `allow` list

```shell
$ curl localhost:8989/api/get-user/x1ah -H 'X-Real-Ip: 127.0.0.1'
{"avatar_url":"https://avatars.githubusercontent.com/u/14919255?v=4","bio":"打字员","blog":"https://when.run","company":"@xiachufang","created_at":"2015-10-01T05:19:27Z","email":null,"events_url":"https://api.github.com/users/x1ah/events{/privacy}","followers":65,"followers_url":"https://api.github.com/users/x1ah/followers","following":69,"following_url":"https://api.github.com/users/x1ah/following{/other_user}","gists_url":"https://api.github.com/users/x1ah/gists{/gist_id}","gravatar_id":"","hireable":true,"html_url":"https://github.com/x1ah","id":14919255,"location":"BeiJing","login":"x1ah","name":"x1ah","node_id":"MDQ6VXNlcjE0OTE5MjU1","organizations_url":"https://api.github.com/users/x1ah/orgs","public_gists":5,"public_repos":67,"received_events_url":"https://api.github.com/users/x1ah/received_events","repos_url":"https://api.github.com/users/x1ah/repos","site_admin":false,"starred_url":"https://api.github.com/users/x1ah/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/x1ah/subscriptions","twitter_username":null,"type":"User","updated_at":"2021-06-10T12:31:47Z","url":"https://api.github.com/users/x1ah"}
```
