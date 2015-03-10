package hello

import (
    "fmt"
    "net/http"
    "time"
)

const VERSION     = "1.8.7"
const NAME        = "gipjson"
const DESCRIPTION = "GeoIP"
const API_STUB    = ""
var   help_usage      = 0
var   json_usage      = 0
var   jsonlong_usage  = 0
var   stats_usage     = 0
var   version_usage   = 0
var   started         = time.Now()

func init() {
  http.HandleFunc(API_STUB+"/api", func(w http.ResponseWriter, r *http.Request) {
    help_usage = help_usage + 1
    http.ServeFile(w, r, "public/api.json")
  })
  http.HandleFunc(API_STUB+"/jsonlong", jsonLongHandler)
  http.HandleFunc(API_STUB+"/json", jsonHandler)
  http.HandleFunc(API_STUB+"/json/", jsonpHandler)
  http.HandleFunc(API_STUB+"/jsonlong/", jsonpLongHandler)
  http.HandleFunc(API_STUB+"/version", versionHandler)
  http.HandleFunc(API_STUB+"/public/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, r.URL.Path[1:]) })
  http.HandleFunc(API_STUB+"/stats", statsHandler)
  http.HandleFunc(API_STUB+"/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "public/index.html")
  })
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
  stats_usage = stats_usage + 1
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprint(w, "{ ")
  fmt.Fprint(w, "\"since\": \"",   started.Format("2006-01-02 15:04:05"), "\"")
  fmt.Fprint(w, ", \"help\":",     help_usage)
  fmt.Fprint(w, ", \"json\":",     json_usage)
  fmt.Fprint(w, ", \"jsonlong\":", jsonlong_usage)
  fmt.Fprint(w, ", \"stats\":",    stats_usage)
  fmt.Fprint(w, ", \"version\":",  version_usage)
  fmt.Fprint(w, ", \"total\":",    help_usage + json_usage + jsonlong_usage + version_usage + stats_usage)
  fmt.Fprint(w, " }")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
  version_usage = version_usage + 1
  w.Header().Set("Content-Type", "text/plain")
  fmt.Fprint(w, VERSION)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    jsonData(w, r)
}

func jsonLongHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    jsonlongData(w, r)
}

func jsonpHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, ";", r.URL.Path[len(API_STUB)+6:], "(")
    jsonData(w, r)
    fmt.Fprint(w, ");")
}

func jsonpLongHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, "; ", r.URL.Path[len(API_STUB)+10:], "(")
    jsonlongData(w, r)
    fmt.Fprint(w, ");")
}

func jsonData(w http.ResponseWriter, r *http.Request) {
  json_usage = json_usage + 1

  fmt.Fprint(w, "{")

  country := r.Header.Get("X-AppEngine-Country")
  fmt.Fprint(w, "\"c\":\"", country, "\"")

  city := r.Header.Get("X-AppEngine-City")
  if city != "" {
    fmt.Fprint(w, ",\"ci\":\"", city, "\"")
  }

  region := r.Header.Get("X-AppEngine-Region")
  if region != "" {
    fmt.Fprint(w, ",\"r\":\"", region, "\"")
  }

  latlong := r.Header.Get("X-AppEngine-CityLatLong")
  if latlong != "" {
    fmt.Fprint(w, ",\"ll\":\"", latlong, "\"")
  }

  fmt.Fprint(w, "}")
}

func jsonlongData(w http.ResponseWriter, r *http.Request) {
  jsonlong_usage = jsonlong_usage + 1

  fmt.Fprint(w, "{ ")

  country := r.Header.Get("X-AppEngine-Country")
  fmt.Fprint(w, "\"country\": \"", country, "\"")

  city := r.Header.Get("X-AppEngine-City")
  if city != "" {
    fmt.Fprint(w, ", \"city\": \"", city, "\"")
  }

  region := r.Header.Get("X-AppEngine-Region")
  if region != "" {
    fmt.Fprint(w, ", \"region\": \"", region, "\"")
  }

  latlong := r.Header.Get("X-AppEngine-CityLatLong")
  if latlong != "" {
    fmt.Fprint(w, ", \"latlong\": \"", latlong, "\"")
  }

  fmt.Fprint(w, " }")
}
