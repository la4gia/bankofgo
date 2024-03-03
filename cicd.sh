#!/bin/bash

if [[ $(docker ps | grep $2) ]] 
    then
        if [[ $(docker ps | grep $3) ]] 
            then 
                docker container stop $3
            else 
                echo "container not running"
        fi
        docker rmi $2
    else
        echo "no image detected"
fi
cd $1 && docker build -t $2 .
docker run -d --rm --name $3 --network bog-network --ip $5 -p $4:$4 $2
