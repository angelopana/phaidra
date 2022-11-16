# ENV for local testing
URL ?= http://phaidra.ai
METHOD ?= GET

build-docker:
	docker build -t phaidra -f Dockerfile .

deploy-docker:
	docker-compose up -d

docker-down:
	docker-compose down

test-it:
	curl --header "Content-Type: application/json" --request POST --data '{"url": "http://phaidra.ai"}' localhost:8080;
	curl --header "Content-Type: application/json" --request GET --data '{"url": "http://phaidra.ai"}' localhost:8080;
	curl --header "Content-Type: application/json" --request GET --data '{"url": "http://google.com.ai"}' localhost:8080;
	curl --header "Content-Type: application/json" --request POST --data '{"url": "https://phaidra.ai/trackrecord"}' localhost:8080;

	curl http://localhost:9095/metrics

own-test:
	curl --header "Content-Type: application/json" --request $(METHOD) --data '{"url": "$(URL)"}' localhost:8080;
	curl http://localhost:9095/metrics
