# Mailer
<sup>An Email sending API built on fiber.</sup>

## Endpoints

```http
POST /send
```
The main endpoint of this API used to send emails, the body of the email is `application/json`.<br>

Body:
```json
{
    "target": "Name <Email_address>",
    "subject": "Email subject",
    "body": "Email content in text/html"
}
```

## Env
The environment variables are quite extensive in this API and you will need all of them for the API to work properly.

```env
EMAIL_PASSWORD=email_password
EMAIL_NAME=notifications
EMAIL_ADDRESS=noreply@example.com
SMTP_PORT=smtp.example.com
SMTP_URI=465
PORT=3600
```

## Config
The API requires an keys.toml file next to the root of the project.<br>
This file concists of the api keys and service names of the keys to give yourself or others access to the API.<br>
To add a key simply add:<br>

```toml
[[keys]]
token="random_character_string_you_generate_yourself"
service="service_name_so_you_can_identify_the_token"
```
To the keys.toml file

## Headers
The API only allows a `Content-Type` of `application/json`.<br>
All the requests need to be authorized with a self-generated `token` as shown above send via the `Authorization` header.<br>
The `Authorization` header has to have the `Bearer` keyword so the end result needs to look similar to this:
```http
Authorization: Bearer random_character_string_you_generate_yourself
```

## Deploy (comming soon)