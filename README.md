# JWT golang auth server

## Project startup
1. configure a .env file at the proj root directory:
   - JWT_SECRET
   - JWT_TTL (seconds)
   - APP_SMTP_KEY (google app key)
   - EMAIL_SENDER (email/gmail that uses app key)
2. if the .env is not your choise, modify docker-compose.yml and add the proper(described above) env variables to the application container
3. Run at the terminal `make clean; make up` (The GNU/make is required to start project)

## API:

### Endpoint
#### Method: POST
#### URI: /auth?guid={{user_id}}
Creates `Access` & `Refresh` key pair

### Endpoint
#### Method: PUT
#### URI: /auth
Will refresh your key pair
#### body
``` json
{
   "Access": "{{access key you have got}},
   "Refresh": "{{refresh key you have got}}"
}
```

Updates `Access` & `Refresh` key pair

## Curl templates to test the rest application:

#### Create key-pair
```
curl --location --request POST 'http://192.168.31.76:8080/auth?guid=e3dc718a-b94d-44d2-a7f1-56891fd51999'
```


#### Refresh key-pair
```
curl --location --request PUT 'http://192.168.31.76:8080/auth' \
--header 'Content-Type: application/json' \
--data '{
"Access": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc2NDgwNzEsImlhdCI6MTcyNzY0NzE3MSwiaXAiOiIxOTIuMTY4LjMxLjc2Iiwic2Vzc2lvbiI6IjM5Zjc1YmFjLWRiNjMtNDY2Ni1iMmEzLWE2ZDgyOTI5ZmJiMSIsInN1YiI6ImUzZGM3MThhLWI5NGQtNDRkMi1hN2YxLTU2ODkxZmQ1MTk5OSJ9.qQuUn4ZQzx60paVZp5V4njfdRnb6REJGuu770LnCifJ-8Yuw6YtZKYaawE4HlIwHAT8ZqBr0bTxh8GM5u4uZRA",
"Refresh": "MTkyLjE2OC4zMS43NiAyMDI0LTA5LTMwIDAwOjU5OjMx"
}'
```

Author: `@countenum404`, Denis Shabashov