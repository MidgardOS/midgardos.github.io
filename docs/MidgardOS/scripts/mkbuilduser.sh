#!/bin/sh

groupadd -g 1001 -f build
useradd -u 1001 -s /bin/bash -g build -d /home/builder -c "MidgardOS build user" -M -N builder
install -v -d -m 755 -o builder -g build /home/builder

setfacl -m u:builder:rwx /MidgardOS/sources
setfacl -m u:builder:rwx /MidgardOS/tools
for dir in etc usr usr/bin usr/include usr/lib usr/lib64 usr/sbin usr/share var/lib; do
    setfacl -m u:builder:rwx /MidgardOS/$dir
done

ln -sv /MidgardOS/sources ~builder/sources
