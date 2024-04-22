# URL Shortener
Pet-project for Alla's portfolio

### Requirements:
- Go 1.21
- Ansible

### Deploy:
```bash
# Local run
make ansible-ping
make ansible-install-docker-compose
make ansible-install-go
make ansible-git-pull

# Connect to the server
make ssh
# Run on the server
go mod vendor
make env-up
make app-restart
docker exec u8views_shortener_go_app ps aux | grep shortener-server
# Exit from the server

# Local run
browse https://shortener.u8views.com/
browse https://dev.shortener.u8views.com/
```
