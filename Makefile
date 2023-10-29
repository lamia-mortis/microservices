docker_start:
	docker compose up -d

docker_stop:
	docker compose down

k8s_start:
	kubectl apply -f k8s/microservices
	kubectl apply -f k8s/ingress.yml
	minikube tunnel