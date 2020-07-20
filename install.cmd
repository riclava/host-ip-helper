@REM Install script

mkdir "C:\Program Files\host-ip-helper"

copy host-ip-helper.exe "C:\Program Files\host-ip-helper\"

sc create HostIPHelperService binpath= "C:\Program Files\host-ip-helper\host-ip-helper.exe" type=share start=auto displayname="Host IP Helper Services" depend=Tcpip