make ssh
# Run on the server
go mod vendor
make env-up
make app-restart
docker exec u8views_shortener_go_app ps aux | grep shortener-server
# Exit from the server