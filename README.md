Api examples:

register username
curl -H "Accept: application/vnd.api+json" -H 'Content-Type: application/vnd.api+json' -X POST -d \
'{"data":{"username":"example1","password":"example1"}}' http://localhost:8080/api/v1/user

get token
curl -H "Accept: application/vnd.api+json" -H 'Content-Type: application/vnd.api+json' -X POST -d \
'{"data":{"username":"example1","password":"example1"}}' http://localhost:8080/api/v1/user/auth

create new city weather
curl -H \
"Authorization: Bearer \
REPLACEWITHTOKEN" \
-H "Accept: application/vnd.api+json" -H 'Content-Type: application/vnd.api+json' \
-X POST -d '{"data":{"city":"medellin", "current_weather": "18", "forecast":"20"}}' \
http://localhost:8080/weathers


get weathers
curl -H \
"Authorization: Bearer REPLACEWITHTOKEN" \
-H "Accept: application/vnd.api+json" -H 'Content-Type: application/vnd.api+json' \
http://localhost:8080/weathers


get markers
curl -H \
"Authorization: Bearer REPLACEWITHTOKEN" \
-H "Accept: application/vnd.api+json" -H 'Content-Type: application/vnd.api+json' \
http://localhost:8080/wheaters/[id]
