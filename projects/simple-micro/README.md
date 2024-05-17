This is a simple example of microservices with Kong API Gateway and Consul service registry.
Follow below steps to run the demostration

#### Set up Kong API getway

```
Inside "gateway" folder
```

#### Setup Consul

```
Inside "registry" folder
```

#### Set up services (user-service, order-service)

```
Run main.go of 2 services to register with Kong and Consul
```

#### Set up routers of services (user-service, order-service)

```
In Konga dashboard
Select each service (example: order-service)
Click Routes
Add a new route for the order-service
```

#### Test

The user-service: Run on port 3000 (host: 10.0.0.16)
The order-service: Run on port 3001 (host: 10.0.0.16)
The Kong API gateway: 10.0.0.16:8000
Client does not access the services directly. it must access via the API gateway

http://10.0.0.16:8000/user-service
http://10.0.0.16:8000/user-service/user/info

http://10.0.0.16:8000/order-service
http://10.0.0.16:8000/order-service/order/list/1
