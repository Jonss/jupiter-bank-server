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
- [ ] - Create user entity
    - [x] - create migration
    - [x] - save in db
    - [ ] - hash password
    - [ ] - create function to handle json response
    - [ ] - test handlers
- [ ] - Implement validation
- [ ] - Create endpoint to get user profile 


### Authentication and Authorization
- [ ] - Implement Login mechanism [PASETO]
- [ ] - Create endpoint to create token
- [ ] - Create middleware to handle tokens
- [ ] - Test tokens
- [ ] - Improve get user profile, allowing only user to get its own profile

### Account
- [ ] - Create account entity
- [ ] - Implement validation
- [ ] - Funding user when account is created 

### Transfer
- [ ] - Create transfer feature
- [ ] - Save transfer

### Messaging
- [ ] - Integrate with email provider [sendgrid?]
- [ ] - Integrate with push notification provider (for web)

### Summary
- [ ] - Provide endpoint with balance and History (https://www.behance.net/gallery/53357679/UniBank)

** Test all features
- Database tests will use testcontainers, testing actual queries
- HTTP route tests use mocks 
