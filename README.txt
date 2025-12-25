To run the code is using 
# go run cmd/http/main.go


To make postgres in database

docker run -d \
  --name postgres18 \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_DB=postgres \
  -p 5434:5432 \
  -v <local destination>:/var/lib/postgresql/18/main \
  postgres:12


To create DB and POSTGRES_USER

//project-local
create database dbname WITH ENCODING 'UTF8';

create user dbuser with password 'password';

grant all privileges on database dbname to dbuser;
//






