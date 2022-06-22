## Start



### Create a k3d cluster

Install k3d

`wget -q -O - https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash`

now create a cluster

`k3d cluster create`

and run

`kubectl apply -f k8s/deployment`