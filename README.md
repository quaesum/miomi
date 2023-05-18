ssh root@178.172.172.30
nano /etc/nginx/nginx.conf
/usr/local/miomi

scp  /home/s228620/GolandProjects/pet/miomi/cmd/server/miomi   root@178.172.172.30:/usr/local/miomi/
scp  -r  /home/s228620/GolandProjects/pet/miomi/config   root@178.172.172.30:/usr/local/miomi/

MINIO_ROOT_USER=admin MINIO_ROOT_PASSWORD=minio3000 minio server /mnt/data --console-address ":9001"

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o miomi
@miomi3000
mysql_dsn: "root:@Miomi3000@tcp(localhost:3306)/miomi?parseTime=true"

/etc/nginx/conf.d/*.conf;
miomi.conf
server {
listen 80;
server_name example.com;

    location /api {
        proxy_set_header   X-Forwarded-For $remote_addr;
        proxy_set_header   Host $http_host;
        proxy_pass         "http://127.0.0.1:8080";
    }
}

        include /etc/nginx/conf.d/*.conf;
        include /etc/nginx/sites-enabled/*;


* server {
*     listen 80;
*     server_name miomi.by www.miomi.by;
*      location /api {
*             proxy_pass  http://localhost:8080;
*         }
* }

/etc/systemd/system/miomi.service

[Unit]
Description=MyApp Go Service
ConditionPathExists=/usr/local/miomi/miomi
After=network.target[Service]
Type=simple
User=burn
Group=burnWorkingDirectory=/usr/local/miomi/miomi
ExecStart=/usr/local/go/bin/go run .Restart=on-failure
RestartSec=10StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=appgoservice[Install]
WantedBy=multi-user.target


[Unit]
Description=Dealingi Core Go Service
ConditionPathExists=/usr/local/miomi
After=network.target

[Service]
Type=simple
User=root
Group=root
LimitNOFILE=1024

Restart=on-failure
RestartSec=10
startLimitIntervalSec=60

WorkingDirectory=/usr/local/miomi
ExecStart=/usr/local/miomi/miomi

# make sure log directory exists and owned by syslog
PermissionsStartOnly=true
ExecStartPre=/bin/mkdir -p /var/log/miomi
ExecStartPre=/bin/chown syslog:adm /var/log/miomi
ExecStartPre=/bin/chmod 755 /var/log/miomi
SyslogIdentifier=miomi

[Install]
WantedBy=multi-user.target

MINIO_VOLUMES="/mnt/data"

MINIO_OPTS=" --console-address :9001"

MINIO_ROOT_USER=minioadmin

MINIO_ROOT_PASSWORD=minioadmin
