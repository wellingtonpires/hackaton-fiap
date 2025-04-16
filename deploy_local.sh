#!/bin/bash

#Implementação automática da aplicação
sudo echo ""
go run vendas/handler.go &

#Realiza build e envia para o DockerHub
docker build -t wellingt0npires/hackaton:latest . && docker push wellingt0npires/hackaton:latest
docker build -f vendas/DockerfileVendas -t wellingt0npires/vendas-hackaton:latest . && docker push wellingt0npires/vendas-hackaton:latest

#Encerra e limpa o minikube
minikube stop
minikube delete
minikube cache delete

#Inicia uma nova instância do Minikube
minikube start --ports=8080:8080 --cpus 4 --memory 8192
sudo chmod 777 /var/run/docker.sock
minikube addons enable metrics-server

#Implementação das APIs e DB
kubectl apply -f hackaton-deployment.yaml,hackaton-service.yaml,postgres-initdb-config.yaml,postgres-claim0-persistentvolumeclaim.yaml,postgres-deployment.yaml,postgres-service.yaml,vendas-hackaton-deployment.yaml,vendas-hackaton-service.yaml

#Configura HPA
kubectl autoscale deployment vendas-hackaton --cpu-percent=80 --min=1 --max=10
kubectl autoscale deployment hackaton --cpu-percent=80 --min=1 --max=10
kubectl autoscale deployment postgres --cpu-percent=80 --min=1 --max=10

