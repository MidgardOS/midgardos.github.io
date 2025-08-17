#!/bin/bash

groupadd -g 1001 -f build
useradd -u 1001 -s /bin/bash -g build -d /home/builder -c "MidgardOS build user" -M -N builder
install -v -d -m 755 -o builder -g build /home/builder

setfacl -m u:builder:rwx /MidgardOS/sources
setfacl -m u:builder:rwx /MidgardOS/tools
setfacl -m u:builder:rwx /MidgardOS/{etc,usr,usr/{bin,include,lib,lib64,sbin,share},var/lib}

ln -sv /MidgardOS/sources ~builder/sources

