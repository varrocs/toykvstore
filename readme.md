

# Kubernetes
## To run minikube
### Install some prequisites (Maybe not needed)
 sudo apt install libvirt-bin
 sudo adduser $USER libvirt

### Virtual box based
- Install virtualbox apt install virtualbox virtualbox-dkms
- Dowload minikube binary from google
    curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
    chmod +x ./minikube

- minikube start --vm-driver=virtualbox
