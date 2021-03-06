gen-go: server/internal/generated/openapicli/api_default.go

server/internal/generated/openapicli/api_default.go: api/task.yaml
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.1 generate --package-name=openapicli -Dapis,models,supportingFiles=client.go -i /local/api/task.yaml -g go -o /local/server/internal/generated/openapicli
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.1 generate --package-name=openapicli -DsupportingFiles=configuration.go -i /local/api/task.yaml -g go -o /local/server/internal/generated/openapicli


gen-ts: front/src/openapi_cli/index.ts

front/src/openapi_cli/index.ts: api/task.yaml
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.1 generate -i /local/api/task.yaml -g typescript-jquery -o /local/front/src/openapi_cli

build-js: gen-ts front/dist/bundle.js

front/dist/bundle.js: front/src/*.ts
	cd front && \
	npx webpack

SERVICES=proxy server front migrate live-reload db-backup

front: gen-ts

$(addprefix build-docker-tl-, $(SERVICES)): build-docker-tl-%: %
	DOCKER_BUILDKIT=1 docker build -t kulti/tl-$< ./$<

build-docker-tl-integration-tests:
	DOCKER_BUILDKIT=1 docker build -f server/tl-integration-tests.Dockerfile -t kulti/tl-integration-tests ./server

run-tl-integration-tests: build-docker-tl-integration-tests build-docker-tl-server build-docker-tl-migrate build-docker-tl-front build-docker-tl-proxy
	./scripts/run-tl-integration-tests.sh

push-docker-images: build-docker-tl-server build-docker-tl-migrate build-docker-tl-front build-docker-tl-proxy build-docker-tl-db-backup
	./scripts/push-docker-images.sh

db-restore:
	cd deployments && \
	source database.env && \
	docker-compose -p prod -f docker-compose.yaml -f docker-compose.prod.yaml run db_migrations up && \
	db_container=$(docker-compose -p prod -f docker-compose.yaml -f docker-compose.prod.yaml ps -q db) && \
	docker cp db.dump ${db_container}:/tmp/ && \
	docker-compose -p prod -f docker-compose.yaml -f docker-compose.prod.yaml exec -T db psql --username=$$POSTGRES_USER --dbname=$$POSTGRES_DB -f /tmp/db.dump

go-coverage:
	cd server && \
	./scripts/go_test.sh && \
	go tool cover -html=coverage.txt

include environments.mk
