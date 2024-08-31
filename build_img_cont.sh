#!/bin/bash


#Build your image
sudo docker build -t forum .

#Watch all you images
sudo docker images

#Build a container 
sudo docker container run -d -p 8080:8000 forum

#Watch all you containers
sudo docker ps -a 