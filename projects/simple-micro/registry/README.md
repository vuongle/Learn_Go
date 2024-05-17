## Set up

https://github.com/hashicorp/learn-consul-docker/blob/main/datacenter-deploy-service-discovery/README.md

## Check setup

http://localhost:8500

## Delete a service

curl -X PUT http://localhost:8500/v1/agent/service/deregister/user-service
