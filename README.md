# jwt-authentication

JWT authentication hands on.

## Libraries

- [gorilla/mux](https://github.com/gorilla/mux)
- [auth0/go-jwt-middleware](https://github.com/auth0/go-jwt-middleware)
- [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)

## Select Signing Methods and Key Types

- HS256
- RS256

## Generate RSA private key & public key

```
# Create private key
openssl genrsa > demo.rsa

# Create public key
openssl rsa -pubout < demo.rsa > demo.rsa.pub
```

## Run local server: http://localhost:8080

```
# TYPE:HS256
JWT_SIGNING_KEY_TYPE=HS256 HS256_SIGNING_KEY=YOUR_SIGNINGKEY go run main.go

# TYPE:RS256
JWT_SIGNING_KEY_TYPE=RS256 RS256_PRIVATE_SECRET_PATH=YOUR_SECRET_PATH RS256_PUBLIC_PATH=YOUR_PUBLIC_PATH go run main.go
```

## How to request to server endpoints

```
# Get Jwt token
curl -XGET http://localhost:8080/auth

# Request with jwt token to root endpoint example
curl -XGET http://localhost:8080 -H "Authorization:Bearer {jwt-token}"
## If work is well, you get bellow message:
{"code":200,"status":"OK","message":"Hello World"}
```

## Referenced below pages or books, thank you

[Go言語で理解するJWT認証 実装ハンズオン](https://qiita.com/po3rin/items/740445d21487dfcb5d9f)
[jwt.io](https://jwt.io)
