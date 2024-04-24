include Makefile.ansible

ssh:
	ssh -t root@70.34.242.4 "cd /var/go/shortener/; bash --login" # TODO_VPS_IP_ADDRESS

tls-certificates:
	mkdir -p /var/lib/u8views/go-url-shortener/tls-certificates/

app-build:
	go build -o /bin/shortener-server ./main.go

app-start:
	shortener-server &

app-stop:
	pkill shortener-server || echo "shortener-server already stopped"

app-restart: tls-certificates app-build app-stop app-start
