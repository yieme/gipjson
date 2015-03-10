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
  "about_url": "/",
  "help_url": "api/",
  "data_json_url": "json/",
  "data_jsonp_url": "jsonp/{callback}",
  "full_data_json_url": "full-json/",
  "full_data_jsonp_url": "full-jsonp/{callback}",
  "stats_url": "stats/",
  "version_url": "version/"
}
```

## Deploy to App Engine

Edit ```app.yaml``` and replace ```gipjson``` with your application name

```sh
goapp deploy
```

## Note

GeoIP data other than country as ```ZZ``` is only available once the application has been deployed to Google App Engine.

## License

MIT
