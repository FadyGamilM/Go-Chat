DB_CONTAINER_NAME=go_chat_db_container
postgres:
	docker run -d --name ${DB_CONTAINER_NAME} -p 1234:5432 -e POSTGRES_USER=go_chat -e POSTGRES_PASSWORD=go_chat postgres:14

create_db:
	docker exec -it ${DB_CONTAINER_NAME} createdb --username=go_chat --owner=go_chat go_chat_db

access_shell:
	docker exec -it ${DB_CONTAINER_NAME} psql -U go_chat

start_db_container:
	docker start ${DB_CONTAINER_NAME}



migrate_up:
	docker run -i -v "H:\1- freelancing path\Courses\golang stack\projects\Go-Chat\db\migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "postgresql://go_chat:go_chat@127.0.0.1:1234/go_chat_db?sslmode=disable" up 1

migrate_down:
	docker run -i -v "H:\1- freelancing path\Courses\golang stack\projects\Go-Chat\db\migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "postgresql://go_chat:go_chat@127.0.0.1:1234/go_chat_db?sslmode=disable" down 1
