.PHONY: update
update:
	go mod tidy
	go mod vendor

.PHONY: proto
proto:
	for f in ./* ; do \
		if [ -d "$${f}/proto" ]; then \
			./proto/bin/protoc.mac -I=$${f}/proto \
				--plugin=protoc-gen-go=./proto/bin/protoc-gen-go.mac \
				--go_out=$${f}/ \
				$${f}/proto/*.proto; \
			echo "mongo" $${f} ; \
			sed -i '' -r -E 's/json:"([^,]+),omitempty"`/json:"\1,omitempty" bson:"\1"`/g' $${f}/proto/*.go || true; \
			sed -i '' -r -E 's/bson:"_id"`/bson:"_id,omitempty"`/g' $${f}/proto/*.go || true; \
		fi \
	done

.PHONY: build
build: proto
	for f in ./* ; do \
		if [ -d "$${f}/proto" ]; then \
			echo "build" $${f} ; \
			CGO_ENABLED=0 GOOS=linux go build -mod vendor -o $${f}/app $${f}/*.go ; \
		fi \
	done

.PHONY: package
package: build
	docker compose build

.PHONY: docker
docker: package
	docker-compose up -d --remove-orphans && \
	docker-compose logs -f broker_grpc broker_http broker_tcp broker_web

.PHONY: down
down:
	docker-compose down

.PHONY: run
run: proto
	CGO_ENABLED=0 GOOS=linux go build -mod vendor -o ./$(APP)/app ./$(APP)/*.go
	docker-compose build $(APP) && \
	docker-compose up -d $(APP) && \
	docker-compose logs -f $(APP)