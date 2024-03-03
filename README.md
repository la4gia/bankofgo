# bankofgo
A simple, vulnerable bank web app for practicing appsec

**Currently in development...** 
**NOT and never intended for production use**

## prerequisites
* Docker, with a custom network
* self-hosted runner
* container running postgres

### creating a docker network
Used for setting our own static IPs for each container. <br>
```docker network create -d bridge --subnet 10.10.10.0/24 bog-network```

### deploy postgres container
Used for storing dummy accounts
```
docker pull postgres:alpine
docker run \
  --name bog-db \
  -e POSTGRES_PASSWORD=bankit \
  --network bog-network \
  --ip 10.10.10.30 \
  -p 5432:5432 -d postgres:alpine
```

### flow
A change to any of the backend files will trigger a rebuild of the image and container for the backend. Likewise for the frontend.
