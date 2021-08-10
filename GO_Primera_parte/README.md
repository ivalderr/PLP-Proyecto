<b>Proyecto:</b> Implementación de RESTful API en Golang.<br>

Enpoints retornan información de ciudades según su nombre o geo locación. <br>
Se utilizaron dos RESTful API:<br><br>
<b>Open Weather API (Free):</b> para obtener la información de la ciudad.<br>
https://openweathermap.org/price <br>
<b>Map Image API:</b> para obtener el mapa de la ciudad.<br>
https://developer.here.com/documentation/map-image/dev_guide/topics/quick-start-show-default-location.html

<b>API Endpoints</b> <br>

Retorna información según las geo coordenadas dadas <br>
/cityByGeoCoordinates/:lon/:lat <br>
Retorna Información de una ciudad según su nombre <br>
/cityByName/:name <br>

<b>Ejemplos:</b> <br><br>
http://localhost:8080/cityByCoor/-118.2437/34.0522 <br>
http://localhost:8080/cityByName/los angeles <br>
http://localhost:8080/cityByName/los%20angeles <br>

<b>Respuesta</b>

{
    "id": 5368361,
    "name": "Los Angeles",
    "country": "US",
    "timeZone": -25200,
    "lon": -118.2437,
    "lat": 34.0522,
    "temperature": 298.09,
    "feelsLike": 298.31,
    "pressure": 1014,
    "humidyty": 64,
    "windSpeed": 0,
    "windDeg": 0,
    "url": "/images/Los Angeles.jpg"
}
