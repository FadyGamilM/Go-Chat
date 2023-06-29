DB_CONTAINER_NAME=go_chat_db_container
postgres:
	docker run -d --name ${DB_CONTAINER_NAME} -p 1234:5432 -e POSTGRES_USER=go_chat -e POSTGRES_PASSWORD=go_chat postgres:14

create_db:
	docker exec -it ${DB_CONTAINER_NAME} createdb --username=go_chat --owner=go_chat go_chat_db

access_shell:
	docker exec -it ${DB_CONTAINER_NAME} psql -U go_chat

start_db_container:
	docker start ${DB_CONTAINER_NAME}