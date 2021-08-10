package main

import (
  "fmt"
  "errors"
  "strconv"
  "io/ioutil"
  "net/http"
  "net/url"
  "os"
  "github.com/gin-gonic/gin"
  "encoding/json"
)

// city represents data about a city.
type city struct {
  ID          float64  `json:"id"`
  Name        string  `json:"name"`
  Country     string  `json:"country"`
  Timezone    float64 `json:"timeZone"`
  Lon         float64 `json:"lon"`
  Lat         float64 `json:"lat"`
  Temperature float64 `json:"temperature"`
  FeelsLike   float64 `json:"feelsLike"`
  Pressure    float64 `json:"pressure"`
  Humidity    float64 `json:"humidyty"`
  WindSpeed   float64 `json:"windSpeed"`
  WindDeg     float64 `json:"windDeg"`
  URL         string  `json:"url"`
}

func main() {
  router := gin.Default()

  //router.GET("/cities/", getCities)
  router.GET("/cityByCoor/:lon/:lat", getCityByCoor)
  router.GET("/cityByName/:name", getCityByName)

  router.Run(apiPath)
}

var apiPath = "localhost:8080"
var cities = []city{}
var ci = city{}

func getCityByName(c *gin.Context) {
  name := c.Param("name")

  if err := getCityInfoByName(name); err == nil {
    c.IndentedJSON(http.StatusOK, ci)
  } else {
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
  }
}

func getCityInfoByName(name string) error {
  endpoint, _ := url.Parse("http://api.openweathermap.org/data/2.5/weather")
  queryParams := endpoint.Query()
  queryParams.Set("q",name)
  queryParams.Set("appid","8b37876e18ca7d87a1149202b5a68c56")
  endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())

  if err != nil {
    return errors.New("The HTTP request failed with error "+err.Error())
  } else {
    var result map[string]interface{}

    data, _ := ioutil.ReadAll(response.Body)
    json.Unmarshal([]byte(data), &result)

    if result["cod"] != "404" {
      ci.ID   = result["id"].(float64)
      ci.Name = fmt.Sprintf("%v", result["name"])
      ci.Country = fmt.Sprintf("%v", result["sys"].(map[string]interface{})["country"])
      ci.Timezone = result["timezone"].(float64)
      ci.Lon = result["coord"].(map[string]interface{})["lon"].(float64)
      ci.Lat = result["coord"].(map[string]interface{})["lat"].(float64)
      ci.Temperature = result["main"].(map[string]interface{})["temp"].(float64)
      ci.FeelsLike = result["main"].(map[string]interface{})["feels_like"].(float64)
      ci.Pressure  = result["main"].(map[string]interface{})["pressure"].(float64)
      ci.Humidity  = result["main"].(map[string]interface{})["humidity"].(float64)
      ci.WindSpeed = result["wind"].(map[string]interface{})["speed"].(float64)
      ci.WindDeg  = result["wind"].(map[string]interface{})["deg"].(float64)
      ci.URL = "/images/"+name+".jpg"

      getCityImageByName(name)
    } else {
      return errors.New("Can't find the specified city :(")
    }
  }
  return nil
}

func getCityImageByName(city string) string {
  var path string

  endpoint, _ := url.Parse("https://image.maps.ls.hereapi.com/mia/1.6/mapview")
  queryParams := endpoint.Query()
  queryParams.Set("ci",city)
  queryParams.Set("apiKey","BvSa9_R77Lc9_ogrJWCjXIjXpUXASxw9foTmnrrQFPI")
  queryParams.Set("sb","km")
  queryParams.Set("ndot","true")
  queryParams.Set("ppi","320") // image quality
  queryParams.Set("w","1280")
  queryParams.Set("h","720")
  queryParams.Set("z","11") // zoom
  endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())

  if err != nil {
    fmt.Printf("The HTTP request failed with error %s\n", err)
  } else {
    f, _ := os.Create("images/"+city+".jpg")
    data, _ := ioutil.ReadAll(response.Body)
    f.Write(data)
		f.Close()
  }

  return path
}

func getCityByCoor(c *gin.Context) {
  lon,_ := strconv.ParseFloat(c.Param("lon"),64)
  lat,_ := strconv.ParseFloat(c.Param("lat"),64)

  if err := getCityInfoByCoor(lon,lat); err == nil {
    c.IndentedJSON(http.StatusOK, ci)
  } else {
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
  }
}

func getCityInfoByCoor(lon, lat float64) error {
  endpoint, _ := url.Parse("http://api.openweathermap.org/data/2.5/weather")
  queryParams := endpoint.Query()
  queryParams.Set("lon",fmt.Sprintf("%v", lon))
  queryParams.Set("lat",fmt.Sprintf("%v", lat))
  queryParams.Set("appid","8b37876e18ca7d87a1149202b5a68c56")
  endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())

  if err != nil {
    return errors.New("The HTTP request failed with error "+err.Error())
  } else {
    var result map[string]interface{}

    data, _ := ioutil.ReadAll(response.Body)
    json.Unmarshal([]byte(data), &result)

    if result["cod"] != "404" {
      ci.ID   = result["id"].(float64)
      ci.Name = fmt.Sprintf("%v", result["name"])
      ci.Country = fmt.Sprintf("%v", result["sys"].(map[string]interface{})["country"])
      ci.Timezone = result["timezone"].(float64)
      ci.Lon = result["coord"].(map[string]interface{})["lon"].(float64)
      ci.Lat = result["coord"].(map[string]interface{})["lat"].(float64)
      ci.Temperature = result["main"].(map[string]interface{})["temp"].(float64)
      ci.FeelsLike = result["main"].(map[string]interface{})["feels_like"].(float64)
      ci.Pressure  = result["main"].(map[string]interface{})["pressure"].(float64)
      ci.Humidity  = result["main"].(map[string]interface{})["humidity"].(float64)
      ci.WindSpeed = result["wind"].(map[string]interface{})["speed"].(float64)
      ci.WindDeg  = result["wind"].(map[string]interface{})["deg"].(float64)
      ci.URL = "/images/"+ci.Name+".jpg"

      getCityImageByCoor(lon,lat)
    } else {
      return errors.New("Can't find the specified city :(")
    }
  }
  return nil
}

func getCityImageByCoor(lon, lat float64) string {
  var path string
  geoc := fmt.Sprint(lat,",",lon)

  endpoint, _ := url.Parse("https://image.maps.ls.hereapi.com/mia/1.6/mapview")
  queryParams := endpoint.Query()
  queryParams.Set("c",geoc)
  queryParams.Set("apiKey","BvSa9_R77Lc9_ogrJWCjXIjXpUXASxw9foTmnrrQFPI")
  queryParams.Set("sb","km")
  queryParams.Set("ndot","true")
  queryParams.Set("ppi","320") // image quality
  queryParams.Set("w","1280")
  queryParams.Set("h","720")
  queryParams.Set("z","11") // zoom
  endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())

  if err != nil {
    fmt.Printf("The HTTP request failed with error %s\n", err)
  } else {
    f, _ := os.Create("images/"+ci.Name+".jpg")
    data, _ := ioutil.ReadAll(response.Body)
    f.Write(data)
		f.Close()
  }

  return path
}
