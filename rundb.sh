# Run Postgres using docker
# For running application without docker compose

set -x

docker run --rm --name postgres -d \
  -p "5432:5432" \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=admin \
  -e POSTGRES_DB=ruleengine postgres:alpine
