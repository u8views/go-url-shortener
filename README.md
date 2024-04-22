# URL Shortener
Pet-project for Alla's portfolio

### Requirements:
- Go 1.21
- Ansible

### Deploy:
```bash
make ansible-ping
make ansible-install-go
make ansible-git-pull
make ssh

# Inside the server
make app-restart
ps aux | grep shortener-server
# Exit from the server

browse https://shortener.u8views.com/
browse https://dev.shortener.u8views.com/
```
