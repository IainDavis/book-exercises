docker run -d --name cadvisor \
  --platform linux/arm64 \
  --network my-network \
  --volume=/:/rootfs:ro \
  --volume=/var/run:/var/run:ro \
  --volume=/sys:/sys:ro \
  --volume=/var/lib/docker/:/var/lib/docker:ro \
  --volume=/cgroup:/cgroup:ro \
  --privileged \
  gcr.io/cadvisor/cadvisor \
