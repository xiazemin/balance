往192.168.1.27的机器的8080端口通过UDP协议发送数据，使用以下命令：
注意：半角英文输入：
echo “hello world!” > /dev/udp/192.168.1.27/8080

表示发送数据包内容”hello world!”到192.168.1.27机器。
通过这种方式，我们可以在服务器上将监控脚本的告警数据内容推送给某个中间服务(192.168.1.27)，这个中间服务可以利用websocket技术将数据推送给前端展示（电子白板预警等形式）


如果往遠程UDP端口發送數據，那麼可以使用以下命令：echo “hello” | socat - udp4-datagram:192.168.1.80:5060

意思是往遠程192.168.1.80的5060端口發送數據包hello


/dev/tcp

$exec 5<>/dev/tcp/time.nist.gov/13; cat <&5 & cat >&5; exec 5>&-
[1] 909

58753 19-09-27 03:59:13 50 0 0   0.0 UTC(NIST) *


$exec 5<>/dev/udp/127.0.0.1/8082; cat <&5 & cat >&5; exec 5>&-

$exec 5<>/dev/udp/127.0.0.1/8082; cat <&5 & cat >&5; exec 5>&-
[1] 2274
-bash: 5: Bad file descriptor
-bash: 5: Bad file descriptor
[1]+  Exit 1                  cat 0<&5



$nc -u 127.0.0.1 8082

$exec 3<>/dev/udp/127.0.0.1/8082

$echo $(read < /dev/udp/127.0.0.1/8082)









https://n0where.net/bash-open-tcpudp-sockets


Bash shell has a built-in feature that allows to open TCP/UDP sockets using a simple syntax. This is very useful when tools like netcat are not installed or we don’t have the permission to use it.

The syntax is

$ exec {file-descriptor}<>/dev/{protocol}/{host}/{port}


{file-descriptor} – 0, 1 and 2 are seserved for stdin, stout and stderr respectively. At least 3 must be used. The Bash manual suggest to be careful in using descriptors above 9 since there could be conflict with descriptors used internally by the shell.
<> – the file is open for both reading and writing
{protocol} – TCP or UDP
{host} – ip address or domain name of the host
{port} – logic port


Sockets can be closed using

$ exec {file-descriptor}<>&-
To send a message through the socket

echo -e -n "$MSG_OUT" >&3
or

printf "$MSG_OUT" >&3
To read a message from the socket

read -r -u -n $MSG_IN <&3
Output can be printed recursively

while read LINE <&3
do
    echo $LINE >&1
done
Or read entirely in one variable

OUTPUT=$(dd bs=$BYTES count=1 <&3 2> /dev/null)


Example:
$ exec 3<>/dev/tcp/127.0.0.1/1234
We are opening a socket for reading and writing to the 1234 port in the loopback interface.

The /dev/tcp and /dev/udp files aren’t real devices but are keywords interpreted by the Bash shell. Being a “bashism” this solution is not portable even if seems that ksh and zsh shells have the same feature enabled.

In this example we fetch the Google main page:

$ exec 3<>/dev/tcp/www.google.com/80
$ echo -e "GET / HTTP/1.1\n\n" >&3
$ cat <&3
It’s good practice to always close file descriptors

$ exec 3<&-
$ exec 3>&-


Finally,  IRC server example:

#!/bin/bash

##########################################################
# Config

NICK="CyberPunk"
SERVER="irc.n0where.net"
PORT=6667
CHANNEL="#CyberPunk"

##########################################################
# Main

exec 3<>/dev/tcp/${SERVER}/${PORT}
echo "NICK ${NICK}" >&3
echo "USER ${NICK} 8 * : ${NICK}" >&3
echo "JOIN ${CHANNEL}" >&3
cat <&3

exit $?


Enable/disable net redirections
More the feature must be enabled in Bash at compile time. To enable it if you want to compile the Bash yourself include the flag

--enable-net-redirections
while to disable it explicitly use

--disable-net-redirections
Each distribution may or not have the feature enabled in their precompiled Bash.




