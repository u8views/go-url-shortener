include Makefile.ansible

ssh:
	ssh -t root@70.34.244.46 "cd /var/go/shortener/; bash --login" # TODO_VPS_IP_ADDRESS

env-up:
	mkdir -p ./.docker/volumes/go/tls-certificates
	docker-compose -f docker-compose.yml --env-file .env up -d

logs:
	docker logs u8views_shortener_go_app

app:
	docker exec -it u8views_shortener_go_app sh

env-down:
	docker-compose -f docker-compose.yml --env-file .env down

env-down-with-clear:
	docker-compose -f docker-compose.yml --env-file .env down --remove-orphans -v # --rmi=all

app-build:
	docker exec u8views_shortener_go_app go build -o /bin/shortener-server ./main.go

app-start:
	docker exec u8views_shortener_go_app shortener-server

app-stop:
	docker exec u8views_shortener_go_app pkill shortener-server || echo "shortener-server already stopped"

app-restart: app-build app-stop app-start
