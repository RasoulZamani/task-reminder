# Terminal App for Reminder Tasks 
A simple task reminder app built in Go that sends notifications and plays a beep sound when a scheduled task is due. Supports periodic tasks like daily and weekly reminders.


## Dependencies

- **Go**: [Download Go](https://golang.org/dl/)
- **Zenity**: Install with `sudo apt-get install zenity`
- **Beep**: Install with `sudo apt-get install beep`

### Optional (System Configuration)

- Enable PC speaker for beep sound:
  ```bash
  sudo modprobe pcspkr
## installation and set up
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
## usage
you could write your tasks in the tasks.json file like:

```
[
    {
        "title": "Retry Novin Jobs",
        "message": "Get up and stretch your legs.",
        "datetime": "2025-05-04T8:45:00+03:30",
        "repeat": "daily"
    }
]
```

## License
MIT License


## TODO:
add ui for adding tasks instead of json files
