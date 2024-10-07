# Food Recipe Website

## Prerequisites

Ensure you have the following installed on your system:
- **Go** (version 1.22 or later)
- **Node.js** (with `npm` or `yarn` for Nuxt.js/Vue.js setup)
- **Docker** (for Hasura)
- **PostgreSQL**

### Clone the Repository

To clone the project repository, run the following command:

```bash
git clone https://github.com/bamlakugetachew1/foodrecipe.git
cd foodrecipe
git checkout master
```

### `Running the Go Backend`

1.To run the Go backend, follow these steps:
Navigate to the Go backend directory:
```bash
cd golang-backend
```
2.Set up environment variables by creating a .env file with the following values in .sample.env

3.Ensure that PostgreSQL is running locally on port 5433. If not, adjust the DB_PORT in the .env file accordingly.

4.Run the backend locally:
```bash
go run cmd/main.go
```

### `Running Hasura in Docker`

1.Make sure Docker is running on your system.

2.Navigate to your Docker setup for Hasura (where the docker-compose-prod.yml file is located) and put appropirate env values in the fields

3.Use Docker Compose to bring up the Hasura container:
```bash
docker-compose up -d
```
4.Once started, Hasura will be available at http://localhost:8080. You can access the Hasura Console via your browser to manage database queries and triggers

### `Running the Nuxt.js Frontend`
1. Navigate to the Nuxt.js frontend directory:
```bash
cd vue-frontend
```
2. Install dependencies:.
```bash
npm install
```
3. Run the Nuxt.js development server:
```bash
npm run dev
```
4. The frontend should now be available at http://localhost:3000


