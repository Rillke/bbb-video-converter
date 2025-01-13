docker-compose build bbb
docker-compose run bbb ash
# TODO: Get ID of latest run!
docker start 80c3675102ed
nohup docker exec -t 80c3675102ed ash -c 'cd /opt/input && ./convert.py' &!
tail -fn 10 nohup.out

