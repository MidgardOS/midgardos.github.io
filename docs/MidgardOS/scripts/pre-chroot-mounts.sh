#!/bin/bash

BRFS=/MidgardOS
for DIR in "dev" "proc" "sys" "run"; do
    install -d -v -m 755 -o root -g root $BRFS/$DIR
done

mount -vt devtmpfs devtmpfs $BRFS/dev
mount -vt devpts devpts -o gid=5,mode=620 $BRFS/dev/pts
mount -vt proc proc $BRFS/proc
mount -vt sysfs sysfs $BRFS/sys
mount -vt tmpfs tmpfs $BRFS/run
if [ -h $BRFS/dev/shm ]; then
  install -v -d -m 1777 ${BRFS}$(realpath /dev/shm)
else
  mount -vt tmpfs -o mode=1777,nosuid,nodev,inode64 tmpfs $BRFS/dev/shm
fi
