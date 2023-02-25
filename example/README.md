# Example

## Run
**In order to simulate the IP address when testing locally, please add `X-Forwarded-For` header as the request ip.**

### Deny specified IP and deny other IP

```shell
go run main.go -c deny_all.json -d
```

allow from `127.0.0.0/8`

```shell
$ curl -i localhost:8989/api/get-user/x1ah -H 'X-Forwarded-For: 127.0.1.1'
HTTP/1.1 200 OK
Cache-Control: public, max-age=3600
Content-Type: application/json; charset=utf-8
X-Krakend: Version undefined
X-Krakend-Completed: true
Date: Sat, 25 Feb 2023 12:51:44 GMT
Content-Length: 1144

{"avatar_url":"https://avatars.githubusercontent.com/u/14919255?v=4","bio":null,"blog":"https://when.run","company":null,"created_at":"2015-10-01T05:19:27Z","email":null,"events_url":"https://api.github.com/users/x1ah/events{/privacy}","followers":79,"followers_url":"https://api.github.com/users/x1ah/followers","following":73,"following_url":"https://api.github.com/users/x1ah/following{/other_user}","gists_url":"https://api.github.com/users/x1ah/gists{/gist_id}","gravatar_id":"","hireable":true,"html_url":"https://github.com/x1ah","id":14919255,"location":"ShenZhen, China","login":"x1ah","name":"x1ah","node_id":"MDQ6VXNlcjE0OTE5MjU1","organizations_url":"https://api.github.com/users/x1ah/orgs","public_gists":5,"public_repos":78,"received_events_url":"https://api.github.com/users/x1ah/received_events","repos_url":"https://api.github.com/users/x1ah/repos","site_admin":false,"starred_url":"https://api.github.com/users/x1ah/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/x1ah/subscriptions","twitter_username":null,"type":"User","updated_at":"2023-01-31T16:44:41Z","url":"https://api.github.com/users/x1ah"}
```

deny other

```shell
$ curl -i localhost:8989/api/get-user/x1ah -H 'X-Forwarded-For: 1.1.1.1'
HTTP/1.1 403 Forbidden
Date: Sat, 25 Feb 2023 12:52:33 GMT
Content-Length: 0
```

### Deny specified IP and allow other IP

```shell
go run main.go -c deny_all.json -d
```

deny from `127.0.0.0/8`

```shell
$ curl -i localhost:8989/api/get-user/x1ah -H 'X-Forwarded-For: 127.0.1.1'
HTTP/1.1 403 Forbidden
Date: Sat, 25 Feb 2023 11:42:59 GMT
Content-Length: 0
```

allow other

```shell
$ curl -i localhost:8989/api/get-user/x1ah -H 'X-Forwarded-For: 9.9.9.9'
HTTP/1.1 200 OK
Cache-Control: public, max-age=3600
Content-Type: application/json; charset=utf-8
X-Krakend: Version undefined
X-Krakend-Completed: true
Date: Sat, 25 Feb 2023 12:32:49 GMT
Content-Length: 1144

{"avatar_url":"https://avatars.githubusercontent.com/u/14919255?v=4","bio":null,"blog":"https://when.run","company":null,"created_at":"2015-10-01T05:19:27Z","email":null,"events_url":"https://api.github.com/users/x1ah/events{/privacy}","followers":79,"followers_url":"https://api.github.com/users/x1ah/followers","following":73,"following_url":"https://api.github.com/users/x1ah/following{/other_user}","gists_url":"https://api.github.com/users/x1ah/gists{/gist_id}","gravatar_id":"","hireable":true,"html_url":"https://github.com/x1ah","id":14919255,"location":"ShenZhen, China","login":"x1ah","name":"x1ah","node_id":"MDQ6VXNlcjE0OTE5MjU1","organizations_url":"https://api.github.com/users/x1ah/orgs","public_gists":5,"public_repos":78,"received_events_url":"https://api.github.com/users/x1ah/received_events","repos_url":"https://api.github.com/users/x1ah/repos","site_admin":false,"starred_url":"https://api.github.com/users/x1ah/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/x1ah/subscriptions","twitter_username":null,"type":"User","updated_at":"2023-01-31T16:44:41Z","url":"https://api.github.com/users/x1ah"}
```
