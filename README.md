# gipjson <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/e/ec/World_map_blank_without_borders.svg/200px-World_map_blank_without_borders.svg.png" align="right" />

GO Google GeoIP uber-server

- [GO](http://golang.org) and [Google](http://google.com) [X-AppEngine-*](http://vikinghammer.com/2013/01/29/appengine-http-headers/) headers for uber-speed
- [Google](https://appengine.google.com/) data accuracy
- JSON & JSONP formats
- Self documenting API
- Usage Statistics

## Demo

[Click Here](http://gipjson.appspot.com)

## Install

1. Install [AppEngine SDK for Go](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go)
2. Clone repo

```sh
git clone https://github.com/yieme/gipjson.git
cd gipjson
```

3. Edit ```app.yaml``` and replace ```gipjson``` with your **uber-server-87917** Project ID

## Run locally

```sh
goapp serve
```

## Run in cloud

Start [AppEngine Project](https://console.developers.google.com/project), ex: ```uber-server-87917```


> Note: GeoIP data other than country as ```ZZ``` is available once the project has been deployed.
> So you will see ```{"c":"ZZ"}``` instead of ```{"c":"US","ci":"middleton","r":"id","ll":"43.706828,-116.620136"}```


## Deploy to App Engine

```sh
goapp deploy
```

## Short form GeoData as JSON

Small footprint optimized

```sh
curl http://uber-server-87917.appspot.com/json/
```

Produces

```js
{"c":"US","ci":"middleton","r":"id","ll":"43.706828,-116.620136"}
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
## License

MIT
