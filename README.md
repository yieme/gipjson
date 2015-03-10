# gipjson <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/e/ec/World_map_blank_without_borders.svg/200px-World_map_blank_without_borders.svg.png" style="float:right" />

Uberfast GeoIP JSON server

- Written in GO
- Google GeoIP data
- JSON & JSONP formats
- Self documenting API
- Usage Statistics

## Install

Clone repo

```sh
git clone https://github.com/yieme/gipjson.git
cd gipjson
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
