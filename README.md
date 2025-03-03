![Minimalist_Calories_API](https://github.com/user-attachments/assets/de91b357-abee-46bf-9404-c41256099c16)

Minimalist Calories API is a RESTful API that is used by the Minimal Calories App.

This API allows user sign up, login and validates endpoints with jwt. Users can add their own food items and manage their daily calorie intakes. Users can also monitor their weights.

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/LodewykRoux/minimalist-calories-api/go.yml?branch=main&label=Build%20%26%20Test&color=fff)
[![Static Badge](https://img.shields.io/badge/go-v1.23.0-blue?logo=go&color=%2300ADD8)](https://go.dev/)
[![Static Badge](https://img.shields.io/badge/gin-v1.9.1-blue?logo=gin&logoColor=white&color=%23008ECF)](github.com/gin-gonic/gin)
[![Static Badge](https://img.shields.io/badge/jwt-v5.2.0-black?logo=jsonwebtokens&logoColor=white&color=%23000000)](https://github.com/golang-jwt/jwt)
[![Static Badge](https://img.shields.io/badge/mysql-v8.0.34-black?logo=mysql&logoColor=white&color=%234479A1)](https://www.mysql.com/)
[![Static Badge](https://img.shields.io/badge/gorm-v1.25.7-black?logo=data%3Aimage%2Fpng%3Bbase64%2CiVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAIAAADYYG7QAAAEQklEQVR4nOyYX0xbVRjAv3POvbe05U9X2LJRK8EhhcGQCTqBbDCmm3%2BWzBmdZPFJoz5ojI8an0yMJj5plmiMWYwxZho1GrME1G0kY0ymgEBho0IpbQe1MGAt%2FXd77znH3LYDOonXN3m4v6fee853vt%2F5znfuQ4VDx5%2BHrQT%2BvwXuxBDSwxDSwxDSwxDSwxDSwxDSwxDSwxDSY8sJCf8yxuxOeu9%2BXmiHdJIE3Ng%2FioBnhzhC9J4HmKMWiIhW5gXPZZSMau8BqfWdYLJmp6HkKvaP4PgKANAdu2ntQTLZR8LTaym4yao0P4kiYXH8fJ4QN1nYNsfaPByeVusfVltOAsKQToEg0T0dODgh%2FXQaUYULJvnx18FcQv7sR0qKOmpTjY9J3R%2BQsJfb71LbTuVtq%2BWk9PNHJOhWm46xuxuoo7bg7BuIs%2Bygsu8JWt%2BpeUwNIDmWE2KlTm61ceu23BJygpVVqC1dIMelng%2B1NOZi%2BcgrzFmn1hwQJy4qzcdBMpu%2Be5vJ8eT2ShpdISFvuvPFgq%2FeBKz1APaPmHpOc4S0XbV2pR96xhx0AxG1xYtKqbNeCIxpVcCEVrfmkpJcaTDHApcsoKRRdBGUFCgpHAlr8xAIYz0k7M1UPipd%2BAT7R%2FHyDQ5AXW3C8Dmkykfb73c%2B%2Bpxj%2F%2BGm8gJmLmY7qzeWBnEuePq1X1b77RMCkOO05kD2iVY2gbkIUrGNURgxlYQ8JORBK%2FPrFbI7MgfnzTWT1Q5EFH%2F9Gi%2FMcnMJFBTisBebLHNlDeHgrM2Eyl17RdGUjVrvD4RUV5tmdiu09pJ4rrCK%2B5i5GADU2naILeMb1zZGrTc1TkTw9G%2B3H0hGRGthVlgqn3ofUGbT8x6p9wwAWIptzcee%2FtGzZCuy%2BCSnT7YosRUpGwXAKhqTL5%2FJLZWIiP1frue71kv3PkJrDoJviJfXCIM%2FcNuuvArBZmjHp92yTJ1iS6bv3xF%2F%2BVgbkCwocQtUJQ7i4Gfv7eYLL1Xz18zDzshE1Xw3Wr2ZC18Kipc%2Bx9NXtfCAmyzOrueLhNH8dXVPu1rXCZwRz%2BU7Um8uRHx%2FaCVtOMIzFxgv%2BnjJzuztQ4zioFutO5xShTRDUatjyVZFt1fNVD5F5iZzQrEl4folqe8LSMWYq5WWVeQdymQfFNppXQcOjGa%2FCP9BaHYYzwxByY5U17vy0VdTJ95SHzwByagw0g0A4sA3rNyVPvTCQijsTlqneWkobRKufouUZF6Z00lh%2BBwgrLQ8m7f4zBAkVwFhYaJ3k9SVNfv%2B%2BRZpRRqEVJwXlfEyJ8KEzAxKFz%2FFsWVtVI4T7%2B9sVzXUdYRkMuf34%2F6z4tSAFkkprWgkgTHy15S23ZsBWu5CiYjgG2K2XUBVYeoK4gwlIigZFcbPI61H7dxSIoxfyH6ckPFngw6GkB6GkB6GkB6GkB6GkB6GkB6GkB6GkB5bTujvAAAA%2F%2F8ti7fl0xX1QQAAAABJRU5ErkJggg%3D%3D&logoColor=white&color=%2338b6ff)](https://github.com/go-gorm/gorm)


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`PORT` - The port you want to run on.

`DB_URL` - The db settings you want to use.
Ex `databaseUser:databaseUserPassword@tcp(ipOfDatabase:PortOfdatabase)/nameOfDatabase?charset=utf8mb4&parseTime=true`

## Run Locally

Clone the project

```bash
  git clone https://github.com/LodewykRoux/minimalist-minimalist-calories-api-api
```

Go to the project directory

```bash
  cd minimalist-minimalist-calories-api-api

```

## API Reference

### Users Endpoints
#### Sign up

```http
  POST /api/users/signup
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name`    | `string` | Name of the user           |
| `email`   | `string` | Email of the user          |
| `password`| `string` | Password entered by user   |

#### Login

```http
  POST /api/users/login
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email`   | `string` | Email of the user          |
| `password`| `string` | Password entered by user   |

#### Validate
Validation is used to ensure that the token the user has stored locally is still valid, and that the token is a valid token connected to a user.
```http
  POST /api/users/validate
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `token`   | `string` | Token of the user          |

#### Logout

```http
  POST /api/users/logout
```

### Weight
#### Get List

```http
  POST /api/weight/getList
```

#### Save

```http
  POST /api/weight/save
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `save`   | `Weight` | Weight class being upserted|

#### Delete

```http
  POST /api/weight/delete
```

| Parameter | Type     | Description                         |
| :-------- | :------- | :-----------------------------------|
| `id`      | `int`    | Id of the weight entry to be deleted|

### Food
#### Get List

```http
  POST /api/food/getList
```

#### Save

```http
  POST /api/food/save
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `save`   | `FoodItem` | FoodItem class being upserted|

#### Delete

```http
  POST /api/food/delete
```

| Parameter | Type     | Description                         |
| :-------- | :------- | :-----------------------------------|
| `id`      | `int`    | Id of the food item to be deleted|

### Daily Entries
#### Get List

```http
  POST /api/dailyEntries/getList
```

#### Save

```http
  POST /api/dailyEntries/save
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `save`   | `DailyEntry` | Daily entry class being upserted|

#### Delete

```http
  POST /api/dailyEntries/delete
```

| Parameter | Type     | Description                         |
| :-------- | :------- | :-----------------------------------|
| `id`      | `int`    | Id of the daily entry to be deleted|


## License

[MIT](https://choosealicense.com/licenses/mit/)
