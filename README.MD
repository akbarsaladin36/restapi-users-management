<h1 align="center">RestAPI User Management System - Backend</h1>

This API is created by me for build my personal project about users management system app. This backend api using Golang with Fiber and GORM package. [More about Golang](https://go.dev/)

## Built With

[![Go/Golang](https://img.shields.io/badge/Golang-1.23.2-cyan.svg?style=rounded-square)](https://go.dev/)
[![Fiber](https://img.shields.io/badge/Fiber-v.2.52.6-blue.svg?style=rounded-square)](https://github.com/gofiber/fiber)
[![GORM](https://img.shields.io/badge/Gorm-v.1.25-blue.svg?style=rounded-square)](https://gorm.io/)


## Requirements

1. <a href="https://go.dev/">Go</a>
2. <a href="https://github.com/gofiber/fiber">Fiber</a>
3. <a href="https://gorm.io/">GORM</a>
4. <a href="https://mariadb.org/">MariaDB</a>
5. <a href="https://github.com/air-verse/air">Air</a>
6. <a href="https://github.com/gofiber/swagger">Swagger For Golang</a>

## How to run the app ?

1. Clone this project.
2. Open app's directory in CMD or Terminal.
3. Turn on Web Server and MariaDB can using Third-party tool like xampp, etc.
4. Create a database with the name #nama_database, and Import file sql to **phpmyadmin**
5. Install Swaggo for Swagger CLI for Golang [here]() and set it globally on your OS system.
6. Choose HTTP Method and enter request url.(ex. localhost:3600/)
7. Type `swag init` to create a docs for Swagger app based on this Rest API.
8. Type `air server --port 8080` to activated the server (You can change 8080 to app port that you use in that project).
9. Open Swagger API documentation link. For example, that link will be like this: `http://localhost:8002/swagger/index.html` (You can change 8002 as port based on your port app that you used on this project).

## Set up project

1. Create a file ".env" then setup environment variable like code below:

```
APP_PORT=<FIBER LISTENING PORT>
DATABASE_HOSTNAME=<YOUR_DB_HOSTNAME>
DATABASE_PORT=<YOUR_DB_PORT>
DATABASE_USERNAME=<YOUR_DB_USERNAME>
DATABASE_PASSWORD=<YOUR_DB_PASSWORD>
DATABASE_NAME=<YOUR_DATABASE_NAME>
```

## Feature

1. Login and Register User
2. Admin
    1. Users API
        1. Get Users
        2. Get User
        3. Create User
        4. Update User
        5. Delete User
    2. Profile API
        1. Get Profile
        2. Update Profile
3. Users
    1. Profile API
        1. Get Profile
        2. Update Profile
    
## License

© [Muhammad Akbar Saladin Siregar](https://github.com/akbarsaladin36/)
