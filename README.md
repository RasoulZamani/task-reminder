# Terminal App for Reminder Tasks 
Simple terminal app for remind your tasks in linux level

## installation and usage
after cloning code you could run it manually (!) or create syestemd for it:
```
nano ~/.config/systemd/user/task-reminder.service
fill it with:
[Unit]
Description=Task Reminder
After=graphical.target

[Service]
ExecStart=/home/yourusername/bin/task-reminder
Restart=always
RestartSec=5
Environment=DISPLAY=:0
Environment=XAUTHORITY=%h/.Xauthority

[Install]
WantedBy=default.target
```
and then run these:
```
systemctl --user daemon-reexec
systemctl --user daemon-reload
systemctl --user enable task-reminder.service
systemctl --user start task-reminder.service

```

## dependencies
golang 

### TODO:
add ui for adding tasks
