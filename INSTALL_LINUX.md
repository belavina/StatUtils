**Linux Installation**

Steps to install `perfmonitor` release on a linux system;

1) Install binary:

    `cp ./perfmonitor /usr/bin/`

2) Install service:

    ```
    cp perfmonitor.service /etc/systemd/system
    systemctl start perfmonitor
    systemctl enable perfmonitor
    ```

3) Data is accessible at `:9159/sysstats`
