# Run Postrgres using docker
# For running application without docker compose

docker run --rm --name postgres \
  -p 5432:5432 \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=admin \
  -e POSTGRES_DB=rules postgres:alpine