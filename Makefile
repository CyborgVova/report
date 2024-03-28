.PHONY: all
all:
	docker build . --tag report
# 	docker run -it --name post -p 8080:8080 report sh
	docker run -d --name post -p 8080:8080 report

.PHONY: compose
compose:
	docker-compose up -d --wait

.PHONY: clean
clean:
	docker stop post
	docker rm post
	docker rmi report:latest

.PHONY: re
re: clean all