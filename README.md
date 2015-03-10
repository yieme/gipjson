# go-google-geoip

Google App Engine GO Lang GeoIP server

## Install

Clone repo

```sh
git clone https://github.com/yieme/go-google-geoip.git
cd go-google-geoip
```

## Run locally

```sh
goapp serve
```

## Usage

Built in help

```sh
curl http://localhost:8080/api
```

Renders

```js
{
  "help_url": "/api",
  "version_url": "/version",
  "about_url": "/",
  "geodata_long_json_url": "/jsonlong",
  "geodata_long_jsonp_url": "/jsonlong/{callback}",
  "geodata_json_url": "/json",
  "geodata_jsonp_url": "/json/{callback}",
  "stats_url": "/stats"
}
```

## Deploy to App Engine

Edit ```app.yaml``` and replace ```gipjson``` with your application name

```sh
goapp deploy
```


## License

MIT
