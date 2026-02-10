## Para subir os containers
docker-compose -p serviceapi up -d

## Para derrubar os containers
docker-compose -p serviceapi down

## Para destruir os volumes
docker-compose -p serviceapi down -v

## Restart
docker-compose -p serviceapi restart

## Roda novamente e aplica novas migrations
docker compose -p serviceapi up flyway -d

## Drop tabelas
drop table logs;
drop table servicos;
drop table usuarios;
drop table flyway_schema_history;