docker
- ps  
  - ps all : `docker ps -a`
  - ps running : `docker ps`
- containers  
  - inspect  
    - all : `docker inspect $(docker ps -a -q)`
    - all running : `docker inspect $(docker ps -q)`
    - /docker_tomcat_1 (tomcat) : `docker inspect 7f9cfe812ce3`
    - /docker_registry_1 (registry) : `docker inspect 05d40e168d02`
    - /docker_simple-http-server_1 (visualjeff/simple-http-server:0.0.1) : `docker inspect ebcc6de91df7`
  - start  
    - all : `docker start $(docker ps -a -q) && docker inspect --format="{{.Name}} {{.State.Status}}" $(docker ps -a -q)`
    - all running : `docker start $(docker ps -q) && docker inspect --format="{{.Name}} {{.State.Status}}" $(docker ps -q)`
    - /docker_tomcat_1 (tomcat) : `docker start 7f9cfe812ce3 && docker inspect --format="{{.Name}} {{.State.Status}}" 7f9cfe812ce3`
    - /docker_registry_1 (registry) : `docker start 05d40e168d02 && docker inspect --format="{{.Name}} {{.State.Status}}" 05d40e168d02`
    - /docker_simple-http-server_1 (visualjeff/simple-http-server:0.0.1) : `docker start ebcc6de91df7 && docker inspect --format="{{.Name}} {{.State.Status}}" ebcc6de91df7`
  - exec -it  
    - /docker_tomcat_1 (tomcat) : `docker exec -it 7f9cfe812ce3 bash`
    - /docker_registry_1 (registry) : `docker exec -it 05d40e168d02 bash`
    - /docker_simple-http-server_1 (visualjeff/simple-http-server:0.0.1) : `docker exec -it ebcc6de91df7 bash`
  - stop  
    - all : `docker stop $(docker ps -a -q) && docker inspect --format="{{.Name}} {{.State.Status}}" $(docker ps -a -q)`
    - all running : `docker stop $(docker ps -q) && docker inspect --format="{{.Name}} {{.State.Status}}" $(docker ps -q)`
    - /docker_tomcat_1 (tomcat) : `docker stop 7f9cfe812ce3 && docker inspect --format="{{.Name}} {{.State.Status}}" 7f9cfe812ce3`
    - /docker_registry_1 (registry) : `docker stop 05d40e168d02 && docker inspect --format="{{.Name}} {{.State.Status}}" 05d40e168d02`
    - /docker_simple-http-server_1 (visualjeff/simple-http-server:0.0.1) : `docker stop ebcc6de91df7 && docker inspect --format="{{.Name}} {{.State.Status}}" ebcc6de91df7`
  - kill  
    - all : `docker kill $(docker ps -a -q) && docker inspect --format="{{.Name}} {{.State.Status}}" $(docker ps -a -q)`
    - all running : `docker kill $(docker ps -q) && docker inspect --format="{{.Name}} {{.State.Status}}" $(docker ps -q)`
    - /docker_tomcat_1 (tomcat) : `docker kill 7f9cfe812ce3 && docker inspect --format="{{.Name}} {{.State.Status}}" 7f9cfe812ce3`
    - /docker_registry_1 (registry) : `docker kill 05d40e168d02 && docker inspect --format="{{.Name}} {{.State.Status}}" 05d40e168d02`
    - /docker_simple-http-server_1 (visualjeff/simple-http-server:0.0.1) : `docker kill ebcc6de91df7 && docker inspect --format="{{.Name}} {{.State.Status}}" ebcc6de91df7`
  - restart  
    - all : `docker restart $(docker ps -a -q) && docker inspect --format="{{.Name}} {{.State.Status}}" $(docker ps -a -q)`
    - all running : `docker restart $(docker ps -q) && docker inspect --format="{{.Name}} {{.State.Status}}" $(docker ps -q)`
    - /docker_tomcat_1 (tomcat) : `docker restart 7f9cfe812ce3 && docker inspect --format="{{.Name}} {{.State.Status}}" 7f9cfe812ce3`
    - /docker_registry_1 (registry) : `docker restart 05d40e168d02 && docker inspect --format="{{.Name}} {{.State.Status}}" 05d40e168d02`
    - /docker_simple-http-server_1 (visualjeff/simple-http-server:0.0.1) : `docker restart ebcc6de91df7 && docker inspect --format="{{.Name}} {{.State.Status}}" ebcc6de91df7`
  - logs -f  
    - /docker_tomcat_1 (tomcat) : `docker logs -f 7f9cfe812ce3`
    - /docker_registry_1 (registry) : `docker logs -f 05d40e168d02`
    - /docker_simple-http-server_1 (visualjeff/simple-http-server:0.0.1) : `docker logs -f ebcc6de91df7`
  - top  
    - /docker_tomcat_1 (tomcat) : `docker top 7f9cfe812ce3`
    - /docker_registry_1 (registry) : `docker top 05d40e168d02`
    - /docker_simple-http-server_1 (visualjeff/simple-http-server:0.0.1) : `docker top ebcc6de91df7`
  - rm -f  
    - all : `docker rm -f $(docker ps -a -q)`
    - all running : `docker rm -f $(docker ps -q)`
    - /docker_tomcat_1 (tomcat) : `docker rm -f 7f9cfe812ce3`
    - /docker_registry_1 (registry) : `docker rm -f 05d40e168d02`
    - /docker_simple-http-server_1 (visualjeff/simple-http-server:0.0.1) : `docker rm -f ebcc6de91df7`
- images  
  - list : `docker images`
  - remove all : `docker rmi $(docker images -q)`
- volumes  
  - list : `docker volume ls`
  - remove all : `docker volume rm $(docker volume ls -q)`
- networks : `docker network ls`
- stats : `docker stats`

\* *generated using [hotshell](https://github.com/julienmoumne/hotshell)*