Api examples:

register username <br/>
curl -H "Accept: application/vnd.api+json" -H 'Content-Type: application/vnd.api+json' -X POST -d \
'{"data":{"username":"example1","password":"example1"}}' http://localhost:8080/api/v1/user

get token

...
curl -H "Accept: application/vnd.api+json" -H 'Content-Type: application/vnd.api+json' -X POST -d \
'{"data":{"username":"example1","password":"example1"}}' http://localhost:8080/api/v1/user/auth


create new city weather
...

curl -H \
"Authorization: Bearer \
REPLACEWITHTOKEN" \
-H "Accept: application/vnd.api+json" -H 'Content-Type: application/vnd.api+json' \
-X POST -d '{"data":{"city":"medellin", "current_weather": "18", "forecast":"20"}}' \
http://localhost:8080/weather


get weathers
...
curl -H \
"Authorization: Bearer REPLACEWITHTOKEN" \
-H "Accept: application/vnd.api+json" -H 'Content-Type: application/vnd.api+json' \
http://localhost:8080/api/v1/weather


getmarkers
...
curl -H \
"Authorization: Bearer REPLACEWITHTOKEN" \
-H "Accept: application/vnd.api+json" -H 'Content-Type: application/vnd.api+json' \
http://localhost:8080/api/v1/wheater/[id]
