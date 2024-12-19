## Notification service:

Deploy DB
> make deploy

login to DB with password `postgres`
> psql -h localhost -U postgres -d notification

run the following query in db
> CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE all the tables and enums in DB from `./schema` directory

install sqlc in your PC
install dependencies
> go get -u .

run service:
> go run .

the service will create a user by default and expects every api to be logged in, to user id can later on be retrieved from context

### Endpoints:
Create notification:
`POST   /notification`

> curl --location 'localhost:8080/notification' \
--header 'Content-Type: application/json' \
--data '{
    "current_price": 10.10,
    "percent_change": 2.3,
    "volume": 34
}'

Get all posts:
`GET    /notifications`
> curl --location 'localhost:8080/notifications'

Send Notification:
`POST   /notification/send`
> curl --location 'localhost:8080/notification/send' \
--header 'Content-Type: application/json' \
--data-raw '{
"notification_id": "{{notification_id}}",
"emails": ["test@mail.com"]
}'
