# jwt-authentication

JWT authentication hands on.

## Libraries

- [gorilla/mux](https://github.com/gorilla/mux)
- [auth0/go-jwt-middleware](https://github.com/auth0/go-jwt-middleware)
- [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)

## Run local server: http://localhost:8080

`SIGNINGKEY=YOUR_SIGNINGKEY go run main.go`

## How to request to server endpoints

```
# Get Jwt token
curl -XGET http://localhost:8080/auth

# Request with jwt token to root endpoint example
curl -XGET http://localhost:8080 -H "Authorization:Bearer {jwt-token}"
```

## Referenced below pages or books

[Go言語で理解するJWT認証 実装ハンズオン](https://qiita.com/po3rin/items/740445d21487dfcb5d9f)
[jwt.io](https://jwt.io)
