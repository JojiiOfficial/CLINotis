# CLINotis
This is a simple tool to view new notifications from your desktopenvironment in you terminal.

# Requirements
- go1.13
- DBus

# Install
Run `build.sh` and enter the username of the desktop user<br>
If you want to get a notification in your terminal (like mail) then you need to do following:<br>
### Bash
\<tbc\><br>

### ZSH
Put this at the end of the .zshrc file:<br>
```bash
precmd() { 
	noti check
}
```
You need to start a new terminal to make this work.<br>

# Usage
Show new notifications:
```bash
noti
```

Check for new notifications:

```bash
noti check
```
