all:
	docker-compose up -d

all_down:
	docker-compose down

build_server:
	docker build -t pricecalculator/server server

run_server:
	docker run -p 9096:9096 pricecalculator/server

build_client:
	docker build -t pricecalculator/client client

run_client:
	docker run -p 5000:80 pricecalculator/client
