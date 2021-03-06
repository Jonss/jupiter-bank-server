# Jupiter Bank Server

## TODO:

### Hello world!

- [x] - Hello world endpoint
- [x] - CD on Heroku using github actions
- [x] - Setup database
- [x] - Setub migrations
- [x] - Setup test containers
- [x] - Setup env variables wth Viper

## Features

### Users

- [x] - Create user entity
- [x] - create migration
- [x] - save in db
- [x] - hash password
- [x] - create function to handle json response
- [x] - test handlers
- [x] - Implement validation
- [x] - Create migration to add indexes on users table

### Security
- [x] - Create "tenants" table, for check if request is allowed to perform changes on server.
- [x] - Create middleware to handle base64 token using secret, key from tenant on open endpoints.

### Authentication and Authorization

- [x] - Implement Login mechanism [PASETO]
- [x] - Create endpoint to create token
- [x] - Create middleware to handle tokens
- [x] - Create endpoint to get user profile protected by authentication
- [x] - Test tokens

### Account

- [ ] - Create account entity [TODO - add steps]
- [ ] - Implement validation
- [ ] - Funding user when account is created

### Transfer
- [ ] - Create transfer feature [TODO - add steps]
- [ ] - Save transfer

### Summary

- [ ] - Provide endpoint with balance and History (https://www.behance.net/gallery/53357679/UniBank)

### Messaging
- [ ] - Integrate with email provider [sendgrid?]
- [ ] - Integrate with push notification provider (for web)

** Test all features

- Database tests will use testcontainers, testing actual queries
- HTTP route tests use mocks 
- services tests will use mocks