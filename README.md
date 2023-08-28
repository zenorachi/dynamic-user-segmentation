[![Golang](https://img.shields.io/badge/Go-v1.21-EEEEEE?logo=go&logoColor=white&labelColor=00ADD8)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

<div align="center">
    <h1>Dynamic User Segmentation service</h1>
    <h5>
        A microservice written in the Go programming language is designed to add users to specific groups (segments) and remove users from segments with the capability to automate these processes. Additionally, it offers the ability to generate comprehensive reports on all operations, including the option to specify particular users, in the form of CSV files or links to CSV files.
    </h5>
    <p>
        English | <a href="README.ru.md">Russian</a> 
    </p>
</div>

### Technologies used:
- [Golang](https://go.dev), [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/), [Nginx](https://nginx.org/ru/)
- [REST](https://ru.wikipedia.org/wiki/REST), [Swagger UI](https://swagger.io/tools/swagger-ui/)
- [JWT Authentication](https://jwt.io/)

### Installation
```shell
git clone git@github.com/zenorachi/dynamic-user-segmentation.git
```

### Getting started
1. **In order for the service to correctly process requests for obtaining a link to a CSV file, it is necessary to:**
    * Register the application in [Google Cloud](https://developers.google.com/workspace/guides/create-project);
    * Create a service account and generate a secret key for it;
    * Add the received secret key to the `secrets` directory.;
    * Modify the environment variable `GDRIVE_CREDENTIALS` in the .env file.
> **Hint:** the service can be launched without integrating with Google Drive.
> In this case, when requesting a link to a CSV file, an error will occur stating that the GDrive service is unavailable.
2. **Setting up environment variables (create a .env file in the project root):**
```dotenv
# Database
export DB_HOST=
export DB_PORT=
export DB_USER=
export DB_NAME=
export DB_SSLMODE=
export DB_PASSWORD=

# Local database
export LOCAL_DB_PORT=

# Postgres service
export POSTGRES_PASSWORD=

# Password Hasher
export HASH_SALT=
export HASH_SECRET=

# Path to Google Drive credentials.json
export GDRIVE_CREDENTIALS=./secrets/your_credentials_file.json

# Gin mode (optional, default - release)
export GIN_MODE=
```
> **Hint:**
if you are running the project using Docker, set `DB_HOST` to "**postgres**" (as the service name of Postgres in the docker-compose).
3. **Compile and run the project:**
```shell
make
```
4. **To test the service's functionality, you can navigate to the address 
http://localhost:8080/docs/index.html to access the Swagger documentation.**

---

### [Examples of requests](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md)

**[Users](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#Users)**
* [Registration](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#1-registration)
* [Authentication](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#2-authentication)
* [Refresh token](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#3-refresh-token)

**[Segments](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#Segments)**
* [Create a segment](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#1-create-a-segment)
* [Create a segment with an indication of the percentage of automatic addition](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#2-create-a-segment-with-an-indication-of-the-percentage-of-automatic-addition)
* [Delete a segment by name](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#3-delete-a-segment-by-name)
* [Delete a segment by ID](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#4-delete-a-segment-by-id)
* [Get all segments](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#5-get-all-segments)
* [Get a segment by ID](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#6-get-a-segment-by-id)

**[User segment addition / removal operations](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#user-segment-addition--removal-operations)**
* [Add segments to a user by a list of names](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#1-add-segments-to-a-user-by-a-list-of-names)
* [Add segments to the user by a list of names with an indication of the ttl](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#2-add-segments-to-the-user-by-a-list-of-names-with-an-indication-of-the-ttl)
* [Add segments to a user by a list of ID](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#3-add-segments-to-a-user-by-a-list-of-id)
* [Add segments to the user by a list of ID with an indication of the ttl](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#4-add-segments-to-the-user-by-a-list-of-id-with-an-indication-of-the-ttl)
* [Delete segments from a user by a list of names](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#5-delete-segments-from-a-user-by-a-list-of-names)
* [Delete segments from a user by a list of ID](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#6-delete-segments-from-a-user-by-a-list-of-id)

**[User-segment relations](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#user-segment-relations)**
* [Get active segments of a user](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#1-get-active-segments-of-a-user)
* [Get active users of a segment](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#2-get-active-users-of-a-segment)

**[Reports](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#Reports)**
* [Getting operation history](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#1-get-operation-history)
* [Getting operation history in the form of a CSV file](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#2-get-operation-history-in-the-form-of-a-csv-file)
* [Getting operation history in the form of a CSV file link](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/docs/examples/01-requests.md#3-get-operation-history-in-the-form-of-a-csv-file-link)

---

### Additional features
1. **Running tests:**
```shell
make test
```
2. **Running the linter:**
```shell
make lint
```
3. **Creating migration files:**
```shell
make migrate-create
```
4. **Migrations up / down:**
```shell
make migrate-up
```
```shell
make migrate-down
```