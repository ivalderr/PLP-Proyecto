package main

import (
  "fmt"
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

  router.Run("localhost:8080")
}

var cities = []city{}

func getCityByCoor(c *gin.Context) {
  lon,_ := strconv.ParseFloat(c.Param("lon"),64)
  lat,_ := strconv.ParseFloat(c.Param("lat"),64)

  getCityImageByCoor(lon,lat)

  c.IndentedJSON(http.StatusOK, gin.H{"message": "Image downloaded"})
}

func getCityByName(c *gin.Context) {
  name := c.Param("name")

  getCityImageByName(name)
  data := getCityInfoByName(name)

  //c.IndentedJSON(http.StatusOK, gin.H{"message": "Image downloaded"})
  c.IndentedJSON(http.StatusOK, data)
}

func getCityInfoByName(name string) string {
  endpoint, _ := url.Parse("http://api.openweathermap.org/data/2.5/weather")
  queryParams := endpoint.Query()
  queryParams.Set("q",name)
  queryParams.Set("appid","8b37876e18ca7d87a1149202b5a68c56")
  endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())

  if err != nil {
    fmt.Printf("The HTTP request failed with error %s\n", err)
  } else {
    f, _ := os.Create("info.json")
    data, _ := ioutil.ReadAll(response.Body)
    f.Write(data)
    f.Close()

    var result map[string]interface{}
    json.Unmarshal([]byte(data), &result)

    if result["cod"] != "404" {
      ci := city {
        ID   : result["id"].(float64),
        Name : fmt.Sprintf("%v", result["name"]),
        Country : fmt.Sprintf("%v", result["sys"].(map[string]interface{})["country"]),
        Timezone : result["timezone"].(float64),
        Lon : result["coord"].(map[string]interface{})["lon"].(float64),
        Lat : result["coord"].(map[string]interface{})["lat"].(float64),
        Temperature : result["main"].(map[string]interface{})["temp"].(float64),
        FeelsLike : result["main"].(map[string]interface{})["feels_like"].(float64),
        Pressure  : result["main"].(map[string]interface{})["pressure"].(float64),
        Humidity  : result["main"].(map[string]interface{})["humidity"].(float64),
        WindSpeed : result["wind"].(map[string]interface{})["speed"].(float64),
        WindDeg  : result["wind"].(map[string]interface{})["deg"].(float64),
        URL : " ",
      }
      return ci.Name
    }
  }
  return "OK"
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

func getCityImageByCoor(lon, lat float64) string {
  var path string
  geoc := fmt.Sprint(lon,",",lat)

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
    f, _ := os.Create("images/map.jpg")
    data, _ := ioutil.ReadAll(response.Body)
    f.Write(data)
		f.Close()
  }

  return path
}
