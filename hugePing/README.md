# hugeping

this program could send a flood of icmp packets in a short time. also, it could send icmp packets concurrently by options.

download the source code 'hugeping.go'

use 'go build hugeping.go' to compile the program.

in shell, running 'hugeping' will use 2 threads to send approximate 16.7 million packets separately.

also could add an option to control the number of threads like 'hugeping x'.

x means the number of threads you want to set.

each thread will send approximate 16.7 million packets.

the range of ip addr the thread send is from x.0.0.1 to x.255.255.254 .



sephrouse. Mar 21, 2016.

transfer this reponsitory into my gadgets reponsitory.