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

---

## Technologies used:
- [Golang](https://go.dev), [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/), [Nginx](https://nginx.org/ru/)
- [REST](https://ru.wikipedia.org/wiki/REST), [Swagger UI](https://swagger.io/tools/swagger-ui/)
- [JWT Authentication](https://jwt.io/)

---

## Installation
```shell
git clone git@github.com/zenorachi/dynamic-user-segmentation.git
```

---

## Getting started
#### [Detailed Guide to Google Drive Integration](./docs/examples/01-google-drive-setup.md)
1. **Google Drive integration:**
    * Register the application in [Google Cloud](https://developers.google.com/workspace/guides/create-project);
    * Create a service account and generate a secret key for it;
    * Add the received secret key to the `secrets/credentials` directory.;
    * Modify the environment variable `GDRIVE_CREDENTIALS` in the .env file.
> **Hint:** the service can be launched without integrating with Google Drive. In this case, you need to leave the value of the `GDRIVE_CREDENTIALS` variable empty.
> In this case, when requesting a link to a CSV file, an error will occur stating that the Google Drive service is unavailable.
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
export GDRIVE_CREDENTIALS=./secrets/credentials/your_credentials_file.json

# GIN mode (optional, default - release)
export GIN_MODE=

# Nginx & HTTPS
# name of the app's service (app)
export APP_HOST=

# app's port (as in main.yml)
export APP_PORT=

# HTTPS connection (443, for example)
export HTTPS_PORT=
```
> **Hint:**
if you are running the project using Docker, set `DB_HOST` to "**postgres**" (as the service name of Postgres in the docker-compose).

3. **(Optional) Adding certificates for Nginx to work correctly:**
> You need to generate certificates and place them in the `secrets/certs` directory in order to be able to access the service via HTTPS.
You can use the [**minica**](https://github.com/jsha/minica) utility.

4. **Compile and run the project:**
```shell
make
```
5. **To test the service's functionality, you can navigate to the address 
http://localhost:8080/docs/index.html to access the Swagger documentation.**
> **Hint:** to complete the authorization in Swagger UI after receiving the JWT token, you need 
to enter `Bearer <your_token>` (without "<" and ">" symbols) in the input field.

---

## [Examples of requests](./docs/examples/02-requests.md)

**[Users](./docs/examples/02-requests.md#Users)**
* [Registration](./docs/examples/02-requests.md#1-registration)
* [Authentication](./docs/examples/02-requests.md#2-authentication)
* [Refresh token](./docs/examples/02-requests.md#3-refresh-token)

**[Segments](./docs/examples/02-requests.md#Segments)**
* [Create a segment](./docs/examples/02-requests.md#1-create-a-segment)
* [Create a segment with an indication of the percentage of automatic addition](./docs/examples/02-requests.md#2-create-a-segment-with-an-indication-of-the-percentage-of-automatic-addition)
* [Delete a segment by name](./docs/examples/02-requests.md#3-delete-a-segment-by-name)
* [Delete a segment by ID](./docs/examples/02-requests.md#4-delete-a-segment-by-id)
* [Get all segments](./docs/examples/02-requests.md#5-get-all-segments)
* [Get a segment by ID](./docs/examples/02-requests.md#6-get-a-segment-by-id)

**[User segment addition / removal operations](./docs/examples/02-requests.md#user-segment-addition--removal-operations)**
* [Add segments to a user by a list of names](./docs/examples/02-requests.md#1-add-segments-to-a-user-by-a-list-of-names)
* [Add segments to the user by a list of names with an indication of the ttl](./docs/examples/02-requests.md#2-add-segments-to-the-user-by-a-list-of-names-with-an-indication-of-the-ttl)
* [Add segments to a user by a list of ID](./docs/examples/02-requests.md#3-add-segments-to-a-user-by-a-list-of-id)
* [Add segments to the user by a list of ID with an indication of the ttl](./docs/examples/02-requests.md#4-add-segments-to-the-user-by-a-list-of-id-with-an-indication-of-the-ttl)
* [Delete segments from a user by a list of names](./docs/examples/02-requests.md#5-delete-segments-from-a-user-by-a-list-of-names)
* [Delete segments from a user by a list of ID](./docs/examples/02-requests.md#6-delete-segments-from-a-user-by-a-list-of-id)

**[User-segment relations](./docs/examples/02-requests.md#user-segment-relations)**
* [Get active segments of a user](./docs/examples/02-requests.md#1-get-active-segments-of-a-user)
* [Get active users of a segment](./docs/examples/02-requests.md#2-get-active-users-of-a-segment)

**[Reports](./docs/examples/02-requests.md#Reports)**
* [Getting operation history](./docs/examples/02-requests.md#1-get-operation-history)
* [Getting operation history in the form of a CSV file](./docs/examples/02-requests.md#2-get-operation-history-in-the-form-of-a-csv-file)
* [Getting operation history in the form of a CSV file link](./docs/examples/02-requests.md#3-get-operation-history-in-the-form-of-a-csv-file-link)

---

## Additional features
1. **Run tests**
```shell
make test
```
> **Hint:** for the tests to work correctly for configurations, it is necessary to comment
out the first line in the Makefile: `include .env`.
2. **Run the linter**
```shell
make lint
```
3. **Create migration files**
```shell
make migrate-create
```
4. **Migrations up / down**
```shell
make migrate-up
```
```shell
make migrate-down
```
5. **Stop all running containers**
```shell
make stop
```

---

## Decisions
In the process of project implementation, I made the following decisions regarding certain contentious issues:

* **How to implement a *many-to-many* relationship between users and segments?**
> **Decision:** it seemed to me that it would be more appropriate to use a linking table called `relations`.
This approach enables easy scalability of the application in the future without altering the current structure.
It simplifies the process of adding segments to a user and removing them, as well as streamlines the associated queries.

* **Which request method is better to use when adding segments to a user (*POST* or *PUT*)?**
> **Decision:** it seemed to me that it would be more appropriate to use the *POST* method since new relations between users and segments are being created.

* **Why in the connecting table `relations`, are the user ID and segment ID not implemented as references?**
> **Decision:** Initially, it was set up this way. The advantages are that no action is required 
to update relations if a segment is deleted. On the downside, if a segment is deleted, all associated relations
are also removed. While this cleanup is beneficial, it results in the loss of operational history.
Hence, the decision was made to forego the use of references and implement a TRIGGER function.
This function responds to segment deletions by erasing all connections to the deleted segment and recording the actions in the `operations` table.

* **How to implement automatic removal of users from segments?**
> **Decision:** there was a choice between creating a separate service that would periodically check whether
the relationship between a user and segments has expired (for example, using `time.Ticker` or the `cron` utility),
or simply launching a separate goroutine if a TTL is specified in the request. This goroutine, using a `select` statement,
would wait for the specified time (or the context to be canceled) and then proceed to remove segments from users.
The chosen approach is relatively simple to implement, especially considering that the functionality to record
operations in the `operations` table was already implemented in the segment removal method for users. 
This greatly simplifies the process. However, it's worth noting that this approach has its downsides,
such as potentially launching a large number of goroutines (so many that it could deplete the available memory),
or other unfavorable scenarios that could disrupt the service. Therefore, if the application were not in a "test" mode,
it would definitely be advisable to use the first approach. But in this case, to expedite the project,
I decided to go with the second option.

* **How to implement automatic user addition?**
> **Decision:** couldn't come up with anything better than to write a TRIGGER function that responds to new records being added to the `segments` table. 
If new segments are added, the function checks the percentage of automatic addition
(optionally specified in the segment creation request).
If the percentage is greater than zero, the function randomly selects users from the entire user pool
(based on the specified percentage) and then adds the created segment to the selected users.
It also records information about the performed operation in the `operations` table. 
This solution has its advantages, such as simplicity and implementation speed. 
However, there are downsides. The issue is that this approach only makes sense if we already have an existing user base.
If we don't, the entire purpose of the function disappears. It seemed somewhat logical to me to create new segments and test them on an existing user base, which is why I decided to stick with this solution.

* **What service to use to generate a link to a CSV file?**
> **Decision:** as soon as I read the task, I immediately understood that I would likely use an S3 storage,
since I already had experience with S3 in the project [**ImageBox**](https://github.com/zenorachi/image-box).
However, I can't deny that I was quite intrigued by the idea from the 2022 internship candidate of using
Google Drive as a service. So, I decided to try something new for myself and implemented 
the generation of a link to the CSV file using the Google Drive API integration.

* **Should an error be returned if a user (segment) has no active segments (users)?**
> **Decision:** it seemed to me that this would be superfluous and not entirely appropriate, because the request is processed correctly,
there is simply no active connection between users and segments.

* **Should I use *multistage build*?**
> **Decision:** I thought it might be unnecessary as it would significantly increase the build time.
