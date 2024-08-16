# Resarch Go Elasticsearch library

Link:

```
https://developer.okta.com/blog/2021/04/23/elasticsearch-go-developers-guide
```

## Install the Go Elasticsearch library

```
go get github.com/elastic/go-elasticsearch/v8
```

#

#

# Resarching ELK stack

https://www.youtube.com/watch?v=Bs41dR_Kf-0&t=266s

## ELK stack?

Elasticsearch + Logstash + Kibana

## Install Elasticsearch in docker

https://hub.docker.com/_/elasticsearch

```
docker network create elastic

docker pull elasticsearch:8.15.0

docker volume create elasticsearch

docker run --name elasticsearch --net elastic --ip 172.24.100.230 -p 9200:9200 -it -e discovery.type=single-node -v elasticsearch:/usr/share/elasticsearch/data elasticsearch:8.15.0

change default password
docker exec -it elasticsearch /usr/share/elasticsearch/bin/elasticsearch-reset-password -u elastic

default user
elastic / 9VOU=p8CJJVIjQz*wYHr

verify
http://localhost:9200
http://172.24.100.230:9200

http://172.24.100.230:9200/stsc/_search
```

## Install Kibana in docker

https://hub.docker.com/_/kibana

```
docker pull kibana:8.15.0

docker volume create kibana

docker run --name kibana --net elastic -p 5601:5601 -it kibana:8.15.0
```
