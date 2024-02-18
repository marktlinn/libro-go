chmod +xwr ./set_env.sh

. ./set_env.sh;


echo "Starting docker container..."
docker run --name mysql-db \
-e MYSQL_USER=$DB_USER \
-e MYSQL_ROOT_PASSWORD=$DB_PASSWORD \
-e MYSQL_PASSWORD=$DB_PASSWORD \
-e MYSQL_DATABASE=$DB_NAME \
-p 3306:3306 \
-d mysql:latest
