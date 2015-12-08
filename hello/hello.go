package hello

import (
    "fmt"
    "net/http"
    "time"
    "io/ioutil"
//    "log"
)

const NAME         = "gipjson"
const VERSION      = "2.3.0"
const DESCRIPTION  = "Uberfast GeoIP JSON server"
const API_STUB     = ""
const ALLOW_DOMAIN = "localhost"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func arrayToString(intArray []byte, e error) string {
  check(e)
  var str string
  for _, value := range intArray {
    str += string(int(value))
  }
  return str
}


var   error_usage      = 0
var   about_usage      = 0
var   help_usage       = 0
var   json_usage       = 0
var   fulljson_usage   = 0
var   stats_usage      = 0
var   version_usage    = 0
var   started          = time.Now()
var   page404a, err404 = ioutil.ReadFile("public/404.html")
var   page404          = arrayToString(page404a, err404)
var   page403a, err403 = ioutil.ReadFile("public/403.html")
var   page403          = arrayToString(page403a, err403)
var   indexA, errIndex = ioutil.ReadFile("public/index.html")
var   indexHtml        = arrayToString(indexA, errIndex)



func init() {
  http.HandleFunc(API_STUB+"/", func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
      about_usage = about_usage + 1
      fmt.Fprint(w, indexHtml)
    } else {
      error_usage = error_usage + 1
      w.WriteHeader(http.StatusNotFound)
      fmt.Fprint(w, page404)
    }
  })
  http.HandleFunc(API_STUB+"/api/", func(w http.ResponseWriter, r *http.Request) {
    help_usage = help_usage + 1
    http.ServeFile(w, r, "public/api.json")
  })
  http.HandleFunc(API_STUB+"/stats/", statsHandler)
  http.HandleFunc(API_STUB+"/version/", versionHandler)
  http.HandleFunc(API_STUB+"/json/", jsonHandler)
  http.HandleFunc(API_STUB+"/full-json/", fullJsonHandler)
  http.HandleFunc(API_STUB+"/jsonp/", jsonpHandler)
  http.HandleFunc(API_STUB+"/full-jsonp/", fullJsonpHandler)
  http.HandleFunc(API_STUB+"/public/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
  })
}



func statsHandler(w http.ResponseWriter, r *http.Request) {
  stats_usage = stats_usage + 1
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprint(w, "{ ")
  fmt.Fprint(w, "\"service\": \"",  NAME, "/", VERSION, "\"")
  fmt.Fprint(w, ", \"since\": \"",  started.Format("2006-01-02 15:04:05"), "\"")
  fmt.Fprint(w, ", \"about\":",     about_usage)
  fmt.Fprint(w, ", \"error\":",     error_usage)
  fmt.Fprint(w, ", \"help\":",      help_usage)
  fmt.Fprint(w, ", \"json\":",      json_usage)
  fmt.Fprint(w, ", \"full-json\":", fulljson_usage)
  fmt.Fprint(w, ", \"stats\":",     stats_usage)
  fmt.Fprint(w, ", \"version\":",   version_usage)
  fmt.Fprint(w, ", \"total\":",     about_usage + help_usage + json_usage + fulljson_usage + version_usage + stats_usage)
  fmt.Fprint(w, " }")
}



func versionHandler(w http.ResponseWriter, r *http.Request) {
  version_usage = version_usage + 1
  w.Header().Set("Content-Type", "text/plain")
  fmt.Fprint(w, VERSION)
}



func jsonHandler(w http.ResponseWriter, r *http.Request) {
  if pseudoCorsCheck(w, r) {
    w.Header().Set("Content-Type", "application/json")
    jsonData(w, r)
  }
}



func fullJsonHandler(w http.ResponseWriter, r *http.Request) {
  if pseudoCorsCheck(w, r) {
    w.Header().Set("Content-Type", "application/json")
    jsonlongData(w, r)
  }
}



func jsonpHandler(w http.ResponseWriter, r *http.Request) {
  if pseudoCorsCheck(w, r) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, ";", r.URL.Path[len(API_STUB)+7:], "(")
    jsonData(w, r)
    fmt.Fprint(w, ");")
  }
}



func fullJsonpHandler(w http.ResponseWriter, r *http.Request) {
  if pseudoCorsCheck(w, r) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, "; ", r.URL.Path[len(API_STUB)+12:], "(")
    jsonlongData(w, r)
    fmt.Fprint(w, ");")
  }
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
  fulljson_usage = fulljson_usage + 1

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

  fmt.Fprint(w, ", \"ua\": \"", r.UserAgent(), "\"")
  fmt.Fprint(w, " }")
}



func pseudoCorsCheck(w http.ResponseWriter, r *http.Request) bool {
/*  refer := r.Referer()
  log.Println(refer)
  domainLen := len(ALLOW_DOMAIN)
  log.Println(ALLOW_DOMAIN)
  log.Println(domainLen)
  allow1 := "http://"  + ALLOW_DOMAIN
  allow2 := "https://" + ALLOW_DOMAIN
  log.Println(allow1)
  log.Println(refer[0:len(allow1)])
  log.Println(allow2)
  log.Println(refer[0:len(allow2)])
  switch true {
  case refer[0:len(allow1)] == allow1:
    return true
  case refer[0:len(allow2)] == allow2:
    return true
  default:
    error_usage = error_usage + 1
    w.WriteHeader(http.StatusForbidden)
    fmt.Fprint(w, page403)
    return false
  } */
  return true
}
