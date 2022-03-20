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
- [ ] - Implement validation
- [ ] - Get user profile
- [ ] - Login (okta? Firebase? JWT? Paseto)

### Login

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
