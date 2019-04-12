docker-compose --project-name unsealer down --rmi local
docker stop unsealer
# docker ps | awk '{print $1}' | grep -v CONTAINER | xargs docker stop
docker network rm safe
rm keys.txt
