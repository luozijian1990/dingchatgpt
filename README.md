# dingchatgpt

## supervisor 
``` text
[program:ding]
command=/opt/ding/new-chat-ding -conf-file /opt/ding/config.yaml 
autostart=true
autorestart=true
startsecs=1
user = root
stderr_logfile=/var/log/ding/ding.log
stdout_logfile=/var/log/ding/ding.log
redirect_stderr = true
stdout_logfile_maxbytes = 20MB
stdout_logfile_backups = 20
```
