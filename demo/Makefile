IMAGENAME=andpaddating:1.0

buil-image:
	docker build -t $(IMAGENAME) .

docker-run:
	docker run \
  -p 8080:8080 \
  -e MYSQL_ADDRS="root:12345678@tcp(host.docker.internal:3306)/kumin?charset=utf8&parseTime=True&loc=Local&multiStatements=true" \
  andpaddating:1.0

upload-image:
	minikube image load $(IMAGENAME)

create-deployment: 
	kubectl apply -f ../deploymentes/app-deployment.yaml

delete-deployment:
	kubectl delete deployment andpad-deployment
