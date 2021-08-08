package main

import (
  "fmt"
  "strconv"
  "io/ioutil"
  "net/http"
  "net/url"
  "os"
  "github.com/gin-gonic/gin"
  //"bytes"
  //"encoding/json"
)

// city represents data about a city.
type city struct {
  ID          string  `json:"id"`
  Name        string  `json:"name"`
  Code        float64 `json:"code"`
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
  router.GET("/cityByCoor/:lon/:lat", getCityByCoor)
  router.GET("/cityByName/:name", getCityByName)

  router.Run("localhost:8080")

  /*jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
  jsonValue, _ := json.Marshal(jsonData)
  response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
  if err != nil {
      fmt.Printf("The HTTP request failed with error %s\n", err)
  } else {
      data, _ := ioutil.ReadAll(response.Body)
      fmt.Println(string(data))
  */
}

func getCityByCoor(c *gin.Context) {
  lon,_ := strconv.ParseFloat(c.Param("lon"),64)
  lat,_ := strconv.ParseFloat(c.Param("lat"),64)

  getCityImageByCoor(lon,lat)

  c.IndentedJSON(http.StatusOK, gin.H{"message": "Image downloaded"})
}

func getCityByName(c *gin.Context) {
  name := c.Param("name")

  getCityImageByName(name)

  c.IndentedJSON(http.StatusOK, gin.H{"message": "Image downloaded"})
}

func getCityImageByName(city string) string {
  var path string

  //response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=guayaquil&appid=8b37876e18ca7d87a1149202b5a68c56")
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
    f, _ := os.Create("images/map.jpg")
    data, _ := ioutil.ReadAll(response.Body)
    f.Write(data)
		f.Close()
  }

  return path
}

func getCityImageByCoor(lon, lat float64) string {
  var path string
  geoc := fmt.Sprint(lon,",",lat)

  //response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=guayaquil&appid=8b37876e18ca7d87a1149202b5a68c56")
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
/*
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
  var newAlbum album

  // Call BindJSON to bind the received JSON to
  // newAlbum.
  if err := c.BindJSON(&newAlbum); err != nil {
      return
  }

  // Add the new album to the slice.
  albums = append(albums, newAlbum)
  c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
  id := c.Param("id")

  // Loop over the list of albums, looking for
  // an album whose ID value matches the parameter.
  for _, a := range albums {
      if a.ID == id {
          c.IndentedJSON(http.StatusOK, a)
          return
      }
  }
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}*/
