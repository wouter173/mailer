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

## Deployment

You can, if you want to (but I wouldn't), run mailer like this:
```sh
$ go run main/app.go
```

or build mailer and run it:

```
$ go build -o mailer main/app.go
$ ./mailer
```
Also note that if you do this you will probably want a .env file in this repo or adjecent to the executable.

### Docker
Now for the cool way to run mailer you will use docker.<br>
Keep in mind when using the docker image that you will need to mount the keys.toml file to `/app/keys.toml`.<br>
Also keep in mind that in docker you can set the env variables with -e when running or by mounting the .env file to `/app/.env`.<br>
First of all build the image with:

```sh
$ make build
```

Or the more customizable way:

```sh
$ sudo docker build -t mailer:latest .
```
<br>
After you built the image you can run it in docker:

```sh
$ sudo docker run -v path/to/your/keys.toml:/app/keys.toml -v path/to/your/.env:/app/.env --network="host" -d mailer:latest
```

Great you now are a mailer pro.
Good luck mailing people!