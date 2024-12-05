docker run -d --name my-nginx \
  --network my-network \
  -v $(pwd)/nginx.conf:/etc/nginx/conf.d/default.conf \
  -p 8080:80 \
  nginx

