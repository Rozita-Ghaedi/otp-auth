# OTP Auth

## MakeFile 

to build and run dockers

``` 
make build
make run
make stop
...
```



## Swagger documentation 

```http://localhost:8080/swagger/index.html```


## Curl Command 

**request-otp**

```
curl -X POST http://localhost:8080/auth/request-otp \
  -H "Content-Type: application/json" \
  -d '{"identifier": "yourPhone"}'
```
response : 

    {"code_dev":"999697","status":"ok"}



**verify-otp**

```
curl -X POST http://localhost:8080/auth/verify-otp \ 
  -H "Content-Type: application/json" \
  -d '{"identifier": "yourPhone", "code": "736243"}'
```

 response :

    {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTc4NTM3OTgsImlhdCI6MTc1NzI0ODk5OCwic3ViIjoiMThlM2VhNGYtMGRiZS00OGExLWI5NGEtODBlMTY0ZmY4MDNjIn0.GKNlrfACrNpC-CrQI0be4ZNOsfU7IX89FGWhU8iXd84","user":{"id":"18e3ea4f-0dbe-48a1-b94a-80e164ff803c","identifier":"09398988490","verified":true,"createdAt":"2025-09-07T12:39:35.018102Z","lastLogin":"2025-09-07T12:43:18.649412Z"}}
