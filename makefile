run:
	docker build -t bilab-backend .
	docker run -d --name bilab-backend  -p 9090:9090 bilab-backend 
stop:
	docker stop bilab-backend
delete:
	docker rm -f bilab-backend
update:
	docker stop bilab-backend
	docker rm -f bilab-backend
	docker build -t bilab-backend .
	docker run -p 9090:9090 bilab-backend -d