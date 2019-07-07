all:
	docker build -t pricecalculator/server server
	docker run -p 9096:9096 pricecalculator/server

build_server:
	docker build -t pricecalculator/server server

run_server:
	docker run -p 9096:9096 pricecalculator/server
