build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose stop

share-proto:
	sh share-proto.sh