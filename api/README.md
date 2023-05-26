# Go Server

## Tech Used

- docker
- postgres
- httprouter

## Running the API

**sudo privledges may be required to run the commands**

- Install postgres via docker

`make installpostgres`

- Run postgres

`make postgres`

- Create the database

`make createdb`

- Run migration

`make migrateup`

- Run the api

`make runapi`

## Endpoints

| Type | Endpoint          | Description                                                             |
|------|-------------------|-------------------------------------------------------------------------|
| GET  | /v1/healthcheck   | Provides information regarding the state of the API                     |
| GET  | /v1/employees     | Returns a list of all employees                                         |
| POST | /v1/employees     | Creates an employee assuming the data posted is a valid employee object |
| POST | /v1/employees/:id | Returns an employee that matches the provided ID                        |


## Data Objects

## Employee

```json
{
	"name": "Brandon",
	"performance": 15,
	"date": "2022-09-02"
}
```

## Health check

```json

{
	"status": "available",
	"system_info": {
		"environment": "development",
		"version": "1.0.0"
	}
}
```

Flags can be passed in to change the environment and some of the other options

## Troubleshooting

- Drop migrations

`make migratedown`

- Delete database

`make dropdb`
