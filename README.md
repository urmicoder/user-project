# user-project



pc@hp:~/P/user-project/rbac-system-arch$ docker network create urmier
86968f034a95f618e25649439858d0606397925175f0795eb1b64d57dd6abd3c
pc@hp:~/P/user-project/rbac-system-arch$ docker run -d --name redis --network=urmier -p 6379:6379 redis
Unable to find image 'redis:latest' locally
latest: Pulling from library/redis
7cf63256a31a: Pull complete 
c8c62be273bb: Pull complete 
d6c5e428cfd7: Pull complete 
fc690ecf94f9: Pull complete 
a37a0a824c7e: Pull complete 
40836d0aa8f0: Pull complete 
4f4fb700ef54: Pull complete 
056ff5d77b71: Pull complete 
Digest: sha256:6aafb7f25fc93c4ff74e99cff8e85899f03901bc96e61ba12cd3c39e95503c73
Status: Downloaded newer image for redis:latest
1a79de80cdd7addd3ef4d03da6602f9fcd8f0df4ac5c2ad881b8415001e84058
pc@hp:~/P/user-project/rbac-system-arch$ docker ps -a
CONTAINER ID   IMAGE                                 COMMAND                  CREATED         STATUS                     PORTS                                       NAMES
1a79de80cdd7   redis                                 "docker-entrypoint.s…"   2 minutes ago   Up 2 minutes               0.0.0.0:6379->6379/tcp, :::6379->6379/tcp   redis
c253b69cc4be   golang:1.24.0                         "/bin/bash"              4 days ago      Exited (129) 4 days ago                                                agitated_ellis
891a6ef901aa   golang:1.24.0                         "/bin/bash"              4 days ago      Exited (129) 4 days ago                                                stupefied_galois
0599b542b08c   rabbitmq:3-management                 "docker-entrypoint.s…"   4 days ago      Exited (0) 4 days ago                                                  rabbitmq
f775289a42ab   golang:1.24.0                         "/bin/bash"              6 days ago      Exited (0) 6 days ago                                                  naughty_engelbart
7852090249c1   gcr.io/k8s-minikube/kicbase:v0.0.45   "/usr/local/bin/entr…"   7 weeks ago     Exited (137) 7 weeks ago                                               minikube
7494834aaac1   ubuntu:20.04                          "/bin/bash"              2 months ago    Exited (0) 4 weeks ago                                                 interview
pc@hp:~/P/user-project/rbac-system-arch$


-----------------------------------------------------------------------------WithDocker----------------------------------------------------------------------------


pc@hp:~/P/user-project/rbac-system-arch$ docker network create urmier
86968f034a95f618e25649439858d0606397925175f0795eb1b64d57dd6abd3c
pc@hp:~/P/user-project/rbac-system-arch$ docker run -d --name redis --network=urmier -p 6379:6379 redis
Unable to find image 'redis:latest' locally
latest: Pulling from library/redis
7cf63256a31a: Pull complete 
c8c62be273bb: Pull complete 
d6c5e428cfd7: Pull complete 
fc690ecf94f9: Pull complete 
a37a0a824c7e: Pull complete 
40836d0aa8f0: Pull complete 
4f4fb700ef54: Pull complete 
056ff5d77b71: Pull complete 
Digest: sha256:6aafb7f25fc93c4ff74e99cff8e85899f03901bc96e61ba12cd3c39e95503c73
Status: Downloaded newer image for redis:latest
1a79de80cdd7addd3ef4d03da6602f9fcd8f0df4ac5c2ad881b8415001e84058
pc@hp:~/P/user-project/rbac-system-arch$ docker ps -a
CONTAINER ID   IMAGE                                 COMMAND                  CREATED         STATUS                     PORTS                                       NAMES
1a79de80cdd7   redis                                 "docker-entrypoint.s…"   2 minutes ago   Up 2 minutes               0.0.0.0:6379->6379/tcp, :::6379->6379/tcp   redis
c253b69cc4be   golang:1.24.0                         "/bin/bash"              4 days ago      Exited (129) 4 days ago                                                agitated_ellis
891a6ef901aa   golang:1.24.0                         "/bin/bash"              4 days ago      Exited (129) 4 days ago                                                stupefied_galois
0599b542b08c   rabbitmq:3-management                 "docker-entrypoint.s…"   4 days ago      Exited (0) 4 days ago                                                  rabbitmq
f775289a42ab   golang:1.24.0                         "/bin/bash"              6 days ago      Exited (0) 6 days ago                                                  naughty_engelbart
7852090249c1   gcr.io/k8s-minikube/kicbase:v0.0.45   "/usr/local/bin/entr…"   7 weeks ago     Exited (137) 7 weeks ago                                               minikube
7494834aaac1   ubuntu:20.04                          "/bin/bash"              2 months ago    Exited (0) 4 weeks ago                                                 interview
pc@hp:~/P/user-project/rbac-system-arch$ docker ps -a
CONTAINER ID   IMAGE                                 COMMAND                  CREATED         STATUS                     PORTS                                       NAMES
1a79de80cdd7   redis                                 "docker-entrypoint.s…"   8 minutes ago   Up 8 minutes               0.0.0.0:6379->6379/tcp, :::6379->6379/tcp   redis
c253b69cc4be   golang:1.24.0                         "/bin/bash"              4 days ago      Exited (129) 4 days ago                                                agitated_ellis
891a6ef901aa   golang:1.24.0                         "/bin/bash"              4 days ago      Exited (129) 4 days ago                                                stupefied_galois
0599b542b08c   rabbitmq:3-management                 "docker-entrypoint.s…"   4 days ago      Exited (0) 4 days ago                                                  rabbitmq
f775289a42ab   golang:1.24.0                         "/bin/bash"              6 days ago      Exited (0) 6 days ago                                                  naughty_engelbart
7852090249c1   gcr.io/k8s-minikube/kicbase:v0.0.45   "/usr/local/bin/entr…"   7 weeks ago     Exited (137) 7 weeks ago                                               minikube
7494834aaac1   ubuntu:20.04                          "/bin/bash"              2 months ago    Exited (0) 4 weeks ago                                                 interview
pc@hp:~/P/user-project/rbac-system-arch$ docker build -t userimage -f dockerfile.user .
DEPRECATED: The legacy builder is deprecated and will be removed in a future release.
            Install the buildx component to build images with BuildKit:
            https://docs.docker.com/go/buildx/

Sending build context to Docker daemon  48.64kB
Step 1/13 : FROM golang:1.24.0-alpine as builder
 ---> 8f8e4dc60ac8
Step 2/13 : WORKDIR /app
 ---> Using cache
 ---> 56b9e4be4899
Step 3/13 : COPY go.mod ./
 ---> Using cache
 ---> 0108658e497a
Step 4/13 : COPY . /app/
 ---> b3dd6c4c167b
Step 5/13 : RUN go mod tidy
 ---> Running in 181bcbbb9b56
go: downloading golang.org/x/crypto v0.35.0
go: downloading github.com/go-redis/redis v6.15.9+incompatible
go: downloading github.com/gorilla/mux v1.8.1
go: downloading github.com/golang-jwt/jwt/v4 v4.5.1
go: downloading github.com/onsi/ginkgo v1.16.5
go: downloading github.com/onsi/gomega v1.36.2
go: downloading golang.org/x/sys v0.30.0
go: downloading github.com/nxadm/tail v1.4.8
go: downloading github.com/google/go-cmp v0.6.0
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading golang.org/x/net v0.33.0
go: downloading gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
go: downloading github.com/fsnotify/fsnotify v1.4.9
go: downloading golang.org/x/text v0.22.0
 ---> Removed intermediate container 181bcbbb9b56
 ---> e8e656ee5039
Step 6/13 : WORKDIR /app/user/cmd
 ---> Running in 22391c261833
 ---> Removed intermediate container 22391c261833
 ---> 4895b418fb0b
Step 7/13 : RUN go build -o /app/user/cmd/user ./user.go
 ---> Running in a9d31d8859e0
 ---> Removed intermediate container a9d31d8859e0
 ---> ce785803d0cd
Step 8/13 : FROM alpine:3.16
 ---> d49a5025be10
Step 9/13 : RUN apk --no-cache add ca-certificates
 ---> Running in dd291598bd6b
fetch https://dl-cdn.alpinelinux.org/alpine/v3.16/main/x86_64/APKINDEX.tar.gz
fetch https://dl-cdn.alpinelinux.org/alpine/v3.16/community/x86_64/APKINDEX.tar.gz
(1/1) Installing ca-certificates (20240226-r0)
Executing busybox-1.35.0-r17.trigger
Executing ca-certificates-20240226-r0.trigger
OK: 6 MiB in 15 packages
 ---> Removed intermediate container dd291598bd6b
 ---> ab550a6bb0ec
Step 10/13 : WORKDIR /app
 ---> Running in d049b561a444
 ---> Removed intermediate container d049b561a444
 ---> b46be72f19f5
Step 11/13 : COPY --from=builder /app/user/cmd/user /usr/local/bin/user
 ---> c68863cbe29a
Step 12/13 : EXPOSE 5002
 ---> Running in 53a9bafd18a0
 ---> Removed intermediate container 53a9bafd18a0
 ---> 19d6c024d78e
Step 13/13 : ENTRYPOINT ["user"]
 ---> Running in 7e6a6c585cc8
 ---> Removed intermediate container 7e6a6c585cc8
 ---> b68870972c70
Successfully built b68870972c70
Successfully tagged userimage:latest
pc@hp:~/P/user-project/rbac-system-arch$ docker run -it --network=urmier -p 5002:5002 --name usercontainer userimage /bin/sh
Redis Ping Response: PONG
&{0xc0000b2180}
APIKEY  false
2025/03/05 20:53:44 1
Query Key urmi_rate_limiting172.18.0.1
Rate Limiting Called: Count  0
2025/03/05 20:54:02 2 -is
Query Key urmi_rate_limiting172.18.0.1
Rate Limiting Called: Count  1
2025/03/05 20:54:02 2 -is
Query Key urmi_rate_limiting172.18.0.1
Rate Limiting Called: Count  2
2025/03/05 20:54:02 2 -is
Query Key urmi_rate_limiting172.18.0.1
Rate Limiting Called: Count  3
2025/03/05 20:54:03 2 -is
Query Key urmi_rate_limiting172.18.0.1
Rate Limiting Called: Count  4
2025/03/05 20:54:04 2 -is
Query Key urmi_rate_limiting172.18.0.1
Rate Limiting Called: Count  5
2025/03/05 20:54:15 2 -is
Query Key urmi_rate_limiting172.18.0.1
Rate Limiting Called: Count  6
2025/03/05 20:54:15 2 -is
Query Key urmi_rate_limiting172.18.0.1
Rate Limiting Called: Count  7
2025/03/05 20:54:15 2 -is
Query Key urmi_rate_limiting172.18.0.1
Rate Limiting Called: Count  8
2025/03/05 20:54:16 2 -is
Query Key urmi_rate_limiting172.18.0.1
Rate Limiting Called: Count  9
2025/03/05 20:54:16 2 -is
Query Key urmi_rate_limiting172.18.0.1