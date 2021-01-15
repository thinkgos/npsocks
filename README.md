# npsocks

socks5 应用服务

默认端口: 10800  
默认日志位置: `临时目录/npsocks` 如linux为`/tmp/npsocks`,日志保留三天 

```bash
$ npsocks -h
npsocks

Usage:
  npsocks [flags]
  npsocks [command]

Available Commands:
  help        Help about any command
  install     Install the daemon server
  remove      Remove the daemon server
  server      Start API server
  start       Start the daemon server
  status      Get the daemon server status
  stop        Stop the daemon server
  version     Get version info

Flags:
  -h, --help   help for npsocks

Use "npsocks [command] --help" for more information about a command.
```