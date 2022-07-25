docker rm $(docker ps -a) 
docker rmi $(docker images -a) 
go build -o . 
docker build -t ansible . 
docker exec -i ansible chmox +x automate-setup
docker run -it ansible bash
