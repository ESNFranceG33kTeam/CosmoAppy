# sAPI

## Table of Contents

- [Setup](#setup)
- [API Documentation](#api-documentation)
  - [HealthCheck](#healthcheck)
  - [Authentication](#authentication)
  - [Adherents endpoint](#adherents-endpoint)
- [Annexes](#annexes)
  - [Database](#database)

## Setup

The api is made in `Golang` aka `Go` and can be use from source or using `Docker`.

- **Helper**

```bash
Usage of /usr/local/bin/sAPI:
  -conf string
        path for the configuration file. (default "test/conf_local.yaml")
  -swagger string
        path for the swagger file. (default "/test/swagger.yaml")
```

- **Build**

```bash
docker build . -f docker/Dockerfile --tag sapi:latest
```

- **Run**

```bash
docker run -v $PWD/conf/folder:/etc/sAPI/conf -p 8080:8080 sapi:latest \
    -conf=/etc/sAPI/conf/conf_docker.yaml -swagger=/etc/sAPI/conf/swagger.yaml
```

## API Documentation

- A build-in interactive documentation is openly available at the `/docs` endpoint.
- The build-in swagger is available at the `/swagger` endpoint.

### HealthCheck

- The healthcheck is available at the `/health` endpoint.

### Authentication

A session token is need to acceed to each endpoint of the API and to use the swagger.

### Adherents endpoint

<details>

#### Get

- Get full list of adherents :

```bash
curl -X GET "https://${MYSERVER}/auth/adherents" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of adherents objects

- Get only one adherent :

```bash
curl -X  GET "https://${MYSERVER}/auth/adherents/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) adherent object

#### Post

- Create a new adherent :

```bash
curl -X POST "https://${MYSERVER}/auth/adherents" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "firstname": "Ash",
        "lastname": "Ketchum",
        "email": "dresseur@indigo.com",
        "dateofbirth": "1987-05-22",
        "student": true,
        "university": "League indigo",
        "homeland": "Kanto",
        "speakabout": "Twitter",
        "newsletter": false
    }'
```

Output : (`json`) adherent object

#### Put

- To update an adherent :

```bash
curl -X PUT "https://${MYSERVER}/auth/adherents/3" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "firstname": "Ash",
        "lastname": "Ketchum",
        "email": "dresseur@indigo.com",
        "dateofbirth": "1987-05-22",
        "student": false,
        "university": "League indigo",
        "homeland": "Kanto",
        "speakabout": "Twitter",
        "newsletter": false
    }'
```

Output : (`json`) adherent object

#### Delete

- To delete an adherent :

```bash
curl -X DELETE "https://${MYSERVER}/auth/adherents/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output

</details>

### ESNcards endpoint

#### Get

- Get full list of esncards :

```bash
curl -X GET "https://${MYSERVER}/auth/esncards" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of esncards objects

- Get only the esncard of a specific adherent :

```bash
curl -X  GET "https://${MYSERVER}/auth/esncards/id_adherent/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) esncard object

- Get only the esncard with the specific esncard code :

```bash
curl -X  GET "https://${MYSERVER}/auth/esncards/esncard/my-esncard-code" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) esncard object

#### Post

- To create a new esncard :

```bash
curl -X POST "https://${MYSERVER}/auth/esncards" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "id_adherent": 2,
        "esncard": "aVeryTooLongCode"
    }'
```

Output : (`json`) esncard object

#### Delete

- To delete an esncard :

```bash
curl -X DELETE "https://${MYSERVER}/auth/esncards/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output

## Annexes

### Database

The db schema is available in [config/sapi_db.mwb](config/sapi_db.mwb).
It can be edit with [MySQL Workbench](https://www.mysql.com/products/workbench/).

![DB schematique](./config/sapi_db.png)

:warning:
> Actually the db creation is coded by hand.
> But in a near future we would like to generated it automatically.
