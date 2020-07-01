# nenech
This repo is Slack-Gateway
dev_env: go1.13.8 darwin/amd64(macos)

GateWay for Slack App by golang from JSON file

## About
This GW does a POST for any Slack webhook-url on behalf of the host.

Please write `configs/xx.json` for your App you want to connect to this GW.

for example
```json
{
	"endpoint": "test",
	"username": "nenech",
	"channel": "test",
	"webhook_url": "your-webhook-url",
	"title": "Message",
	"message": "test message",
	"color": "good"
}
```

This GW adds a POST route to the web API based on this JSON file.

You can send a message to slack by sending a POST request to the specified endpoint.

## Usage

### Server side
```sh
git clone https://github.com/silmin/nenech
```

move `./nenech`

```sh
go run *.go
```

### Client side
```sh
curl http://<your-server-addr>/<specified-endpoint>
```
The contents of the configuration file corresponding to the endpoint will be posted to Slack.


In option,
```sh
curl http://<your-server-addr>/<specified-endpoint>\?username\=hoge
```
Username can be set to "hoge".

You can set `username`/`title`/`message`/`color(good|warning|danger|(color code))`.
