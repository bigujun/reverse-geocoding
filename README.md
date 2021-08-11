# Reverse Geocoding API

Minimal API writen in go to return city and state from a coordinate

Currently only have data from Brazil

The data was taken from IBGE at [BR_Localidades_2010_v1.kml](https://geoftp.ibge.gov.br/organizacao_do_territorio/estrutura_territorial/localidades/Google_KML/BR_Localidades_2010_v1.kml)

Application will load data from the kml file and serve a endpoint to search nearest place from a pair of cordinates (latitude,longitude)  


## Routes

### GET /api/place?cord={latitude},{longitude}
Returns json with nearest city found for the given coordinates

Ex: /api/place?cord=-32.0559762,-52.145971

Respose:
```json
{
    "id":20070,
    "type":"URBANO",
    "city":"RIO GRANDE",
    "state":"RIO GRANDE DO SUL",
    "lat":-32.0087315314711,
    "long":-52.1192092383167
}
```

# License
[MIT](LICENCE)