#### Dang follow: Run each container for each step

https://www.youtube.com/watch?v=wkWWzEMhiD0
https://docs.konghq.com/gateway/latest/install/docker/

docker network create kong-net
-> OK

docker run -d --name kong-database --network=kong-net -p 5432:5432 -e "POSTGRES_USER=kong" -e "POSTGRES_DB=kong" -e "POSTGRES_PASSWORD=kongpass" postgres:13
-> OK

docker run --rm --link kong-database:kong-database --network=kong-net -e "KONG_DATABASE=postgres" -e "KONG_PG_HOST=kong-database" -e "KONG_PG_PASSWORD=kongpass" kong/kong-gateway:3.6.1.3 kong migrations bootstrap
-> FAIL
Error: [PostgreSQL error] failed to retrieve PostgreSQL server_version_num: [cosocket] DNS resolution failed: failed to receive reply from UDP server 127.0.0.11:53: timeout. Tried: ["(short)kong-database:(na) - cache-miss","kong-database:33 - cache-miss/querying"]
-> Loi nay giong voi khi chay docker-compose.yml tren moi truong windows

#### TODO

- Run docker compose "registry" service in mac
- If OK -> continue the tutorial
- If FAIL -> try above steps in this file
