## Sidecar Hands-on Exercise #1: TOPZ
This exercise has the reader starting an arbitrary application container that exposes some HTTP API, and runs a sidecar container in the same namespace which provides an additional `/topz` endpoint which provides resource usage data for the container host.

<details>
<summary>Exercise definition</summary>

* Step1
```shell
docker run -d <my-app-image>
<container-hash-value>
```

* Step 2
```shell
docker run --pid container:${APP_ID} \
    -p 8080:8080 \
    brendanburns/topz:db0fa58 \
    /server --addr-0.0.0.0:8080
```
</details>

## Setup
I'm using the default installation of Rancher Desktop (v1.16.0) on Apple M1 Silicon, along with the provided `dockerd` container engine

I will do my best to translate the command-line inputs suggested by the book into Dockerfiles (with the aid of my AI assistant) that I can store in this repository

## Challenges
1. The author doesn't provide the main application for the exercise
    > In the absence of having a specific application to run, on the advice of my code-assistant AI, I've opted to use a standard NGINX container as my base application. This gives me some default behavior

2. The provided image expects a `Linux/AMD64` architecture.
    > My code assistant AI has suggested two alternatives:
    > 1. use an alternate sidecar application (cAdvisor)
    > 2. Build a Multi-Arch Image of Topz
    > Since this is all practice, I will try both approaches

---
## Solutions
* [Solution 01 | cadvisor sidecar][01]
* [Solution 02 | multi-arch topz image]

<!-- NAMED LINKS -->
[01]: 01/TOPZ01.md
