# Kong/Konga/Postgres in Docker Compose

This is the official Docker Compose template for [Kong][kong-site-url]. This implementation has Konga (Kong GUI).

# How to use this template

This Docker Compose template provisions a Kong container with a Postgres database, plus a nginx load-balancer. After running the template, the `nginx-lb` load-balancer will be the entrypoint to Kong.

To run this template execute:

```shell
$ docker-compose up
```

To scale Kong (ie, to three instances) execute:

```shell
$ docker-compose scale kong=3
```

In case of error with COMPOSE_HTTP_TIMEOUT: run on this way:

```shell
$ COMPOSE_HTTP_TIMEOUT=3600 docker-compose up
```

Run Konga

```
http://localhost:1337/register

admin account
    username: vuong
    email: vuonglg@gmail.com
    pwd: vuong1986
```

Kong gateway URL (Client access services via this URL)

```
http://localhost:8000/
```

Kong admin URL

```
http://localhost:8001/
```

Create a Konga Connection to Kong

```shell
Name: Kong (or any name)
Kong Admin URL: http://kong:8001/
```

Kong will be available through the `nginx-lb` instance on port `8000`, and `8001`. You can customize the template with your own environment variables or datastore configuration.

Kong's documentation can be found at [https://docs.konghq.com/][kong-docs-url].

## Official kong github

https://github.com/Kong/docker-kong/tree/master/compose

## docker-compose.yml in this example is the following

https://github.com/abrahamjoc/docker-compose-kong-konga/tree/master
