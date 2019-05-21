# Supervisord

Supervisor is an application to reload your server whenver the application gets killed. A process ID will be assigned to your server. To restart the app properly, we need to kill the existing instances and restart the application. We can write one such program in Go. But in order to not reinvent the wheel, we are using a popular program called supervisord.

## Install supervisord

```bash
brew install supervisor
```

## Create configuration file at: 

```bash
/etc/supervisor/conf.d/goproject.conf
```

## Launch `supervisorctl` with the command

```bash
supervisorctl
```