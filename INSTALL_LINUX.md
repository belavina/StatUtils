**Linux Installation**

Steps to install `perfmonitor` release on a linux system;

Install binary:

`cp ./dist/perfmonitor /usr/bin/`

Install service:

```
cp perfmonitor.service /etc/systemd/system
systemctl start perfmonitor
systemctl enable perfmonitor
```

Data is accessible at `:9159/sysstats`