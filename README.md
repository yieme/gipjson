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
  "help_url": "api/",
  "version_url": "version/",
  "about_url": "/",
  "full_data_json_url": "full-json/",
  "full_data_jsonp_url": "full-jsonp/{callback}",
  "data_json_url": "json/",
  "data_jsonp_url": "jsonp/{callback}",
  "stats_url": "stats/"
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
