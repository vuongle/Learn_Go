## Set up

https://github.com/Kong/docker-kong/tree/master/compose

## Run with db

```
KONG_DATABASE=postgres docker compose --profile database up -d
```

## Check setup

http://localhost:8001/

## another reference: include kong + konga + postgress

https://agitt.medium.com/kong-api-gateway-pantsel-ui-with-docker-209eb37d3eaf
