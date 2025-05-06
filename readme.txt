
This is not an official regify product! It is just an example!
No support, no warranty. Use at your own risk!

This is a simple web server example that listens for POST requests
on /post and hands over the received json "data" to the regiboxctl
command.

IMPORTANT: Before running
1) you need to define the config directory of your regiboxd instance!
2) you need to define listing interface and port (default localhost:8081)
3) the regiboxctl command must be on PATH

Other notes:
This is currently for Linux only! For Windows you might need to use -J option
and format your json different!
See https://manuals.regify.com/docs/regiboxd/current/win_reference_manual/#_windows_dos_syntax
This may also apply on the returned JSON values!

There is no complete error handling! For production use, please enhance
the code with additional error handling, logging and exit functionality.

There is no authentication! This API is open for everyone with access to
the interface you are listening to. If you plan to use this in production,
ensure that only the allowed systems are having access to this machine 
and/or port. Maybe by using iptables firewall.

If you plan to run this as a daemon (eg systemd), we recommend compiling
to a binary file.

If echo framework is missing, try running 

 go mod tidy

to install any missing dependencies.


In order to do a REST call, you can execute function calls like this (linux):

 curl -X POST http://localhost:8081/post -H "Content-Type: application/json" -d '{"data": "{\"op\": \"someFunction\"}"}'

