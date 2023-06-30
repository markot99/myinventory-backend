# myInventory-backend

Everyone owns a variety of valuable items. In case of theft, fire or other misfortune, the insurance company needs proof of ownership of the items. Most of the time, the proofs are photos and invoices.

The backend will provide functions for inventorying multiple items with pictures and invoices. In addition, user management will be added to allow users to authenticate themselves.

myInventory was created as part of an examination assignment for the module Backend Development.

Further information on the project and the architecture can be found [here](./assets/myInventory.pdf).

## Deployment
The easiest way to deploy myInventory is with the help of Docker. Instructions for installing Docker can be found here.

To start myInventory, change to the root of the repository and enter the following:

```console
docker compose up -d --build
```

The services are then available under the following URLs:
- Frontend: http://localhost:8080/frontend
- Swagger: http://localhost:8080/api/swagger/index.html
- User Service: http://localhost:8080/api/v1/users
- Item Service: http://localhost:8080/api/v1/items

### Configuration
The individual services can be configured via environment variables. All environment variables are optional and have default values.

**Api Gateway**:
- __PORT__ can be used to set the internal port.
- __FORWARD_X__ can can be used to configure the Api Gateway. This can be used to set the host and the path for the forwarding. X must be incremented. Example: FORWARD_1=host;path, FORWARD_2=host;path, ...
- __LOG_LEVEL__ can be used to set the log level. Permissible values are *debug*, *info*, *warning* and *error*.

**Frontend Service**:
- __PORT__ can be used to set the internal port.
- __LOG_LEVEL__ can be used to set the log level. Permissible values are *debug*, *info*, *warning* and *error*.

**Item Service**:
- __PORT__ can be used to set the internal port.
- __LOG_LEVEL__ can be used to set the log level. Permissible values are *debug*, *info*, *warning* and *error*.
- __MONGO_DB_HOST__ host url of the mongo db instance
- __MONGO_DB_DATABASE__ name of the mongodb database
- __ITEM_COLLECTION__ name of the item collection
- __IMAGE_STORAGE_PATH__ storage path of the images
- __INVOICE_STORAGE_PATH__ stroage path of the invoices
- __JWT_SECRET__ Secret of the JWT tokens. Must match the user service

**Swagger Service**:
- __PORT__ can be used to set the internal port.
- __LOG_LEVEL__ can be used to set the log level. Permissible values are *debug*, *info*, *warning* and *error*.

**User Service**:
- __PORT__ can be used to set the internal port.
- __LOG_LEVEL__ can be used to set the log level. Permissible values are *debug*, *info*, *warning* and *error*.
- __MONGO_DB_HOST__ host url of the mongo db instance
- __MONGO_DB_DATABASE__ name of the mongodb database
- __USER_COLLECTION__ name of the user collection
- __JWT_SECRET__ Secret of the JWT tokens. Must match the item service

## Development

### Development environment
Visual Studio Code should be used as the development environment. Appropriate launch configurations and tasks for development are already available.

### Generate swagger documentation
To generate the Swagger documentation, swag must be installed. Swag can be installed with the following command:
```console
go install github.com/swaggo/swag/cmd/swag@latest
```
More information about swag can be found [here](https://github.com/swaggo/swag).

## Testing
Test files are located directly next to the Go files. MongoDB must be running for the tests to be successful. Tests can be started with the following command from the root of the repository:
```
go test -p 1 ./... -cover
```