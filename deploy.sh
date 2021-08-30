docker build -t thanhdhnt/multi-reactk8s:latest -t thanhdhnt/multi-reactk8s:$SHA -f ./multi-reactk8s/Dockerfile ./multi-reactk8s
docker build -t thanhdhnt/multi-apik8s:latest -t thanhdhnt/multi-apik8s:$SHA -f ./multi-apik8s/Dockerfile ./multi-apik8s
docker push thanhdhnt/multi-reactk8s:latest
docker push thanhdhnt/multi-apik8s:latest
docker push thanhdhnt/multi-reactk8s:$SHA
docker push thanhdhnt/multi-apik8s:$SHA

kubectl apply -f k8s

kubectl set image deployments/client-deployment client=thanhdhnt/multi-reactk8s:$SHA
kubectl set image deployments/server-deployment server=thanhdhnt/multi-apik8s:$SHA