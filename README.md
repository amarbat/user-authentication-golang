# user-authentication-golang
Simple MongoDB Go Authentication api using JWT token.
gin-gonic/gin Backend framework with MongoDB Go driver.

Setup db: use docker to download mongodb and start on default docker mongodb port: 27017
.env:
  port as port number
  MONGO_URL as /mongodb://localhost:27017 format


##Sample usage

Please use some tool to send requests. I used Postman for testing the apis.

#Routes
  ##Signup
  POST:http://localhost:8000/users/signup
  Body(raw-JSON):
{
"first_name":"Hello",
"last_name":"World",
"email":"worlds@gmail.com",
"Password":"password",
"phone":"1234567889"
}

should return: {
    "InsertedID": "60a8001418f113af23a21e80"
}
id might be different.

  ##Login
  POST:http://localhost:8000/users/login
  Body(raw-JSON):
{
"email":"worlds@gmail.com",
"Password":"password"
}

should return:
{
    "ID": "60a74524d31e909feb050e0c",
    "first_name": "Hello",
    "last_name": "World",
    "Password": "$2a$14$m93Z7xqm380em3dx1FRj.uv.50whXXKMy3MnLUxpsQ.BkpIf8xZCO",
    "email": "worlds@gmail.com",
    "phone": "7034857504",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Indvcmxkc0BnbWFpbC5jb20iLCJGaXJzdF9uYW1lIjoiSGVsbG8iLCJMYXN0X25hbWUiOiJXb3JsZCIsIlVpZCI6IjYwYTc0NTI0ZDMxZTkwOWZlYjA1MGUwYyIsImV4cCI6MTYyMTY2MTYwNn0.C9VB3jBSnk3z2bE7d9xhyY3rw2Ks66FONM0aueeKdFc",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X25hbWUiOiIiLCJMYXN0X25hbWUiOiIiLCJVaWQiOiIiLCJleHAiOjE2MjIxODAwMDZ9.PXIwKzI_T3Eo5lKS9sQXjKquMmvgivdB15gI7P4X23k",
    "created_at": "2021-05-21T05:29:08Z",
    "updated_at": "2021-05-21T05:33:26Z",
    "user_id": "60a74524d31e909feb050e0c"
}

## API - 1
GET:http://localhost:8000/api-1
Headers: token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Indvcmxkc0BnbWFpbC5jb20iLCJGaXJzdF9uYW1lIjoiSGVsbG8iLCJMYXN0X25hbWUiOiJXb3JsZCIsIlVpZCI6IjYwYTc0NTI0ZDMxZTkwOWZlYjA1MGUwYyIsImV4cCI6MTYyMTY2MTYwNn0.C9VB3jBSnk3z2bE7d9xhyY3rw2Ks66FONM0aueeKdFc

Should return: {
    "success": "Access granted for api-1"
}


## API -2 
GET:http://localhost:8000/api-2
HEaders: token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Indvcmxkc0BnbWFpbC5jb20iLCJGaXJzdF9uYW1lIjoiSGVsbG8iLCJMYXN0X25hbWUiOiJXb3JsZCIsIlVpZCI6IjYwYTc0NTI0ZDMxZTkwOWZlYjA1MGUwYyIsImV4cCI6MTYyMTY2MTM0OH0.V5vW-6uAJUGk2kKj5ZyyBv54oi8WltdegkTVzPdS-5I

Should return {
    "success": "Access granted for api-2"
}

