#!/bin/bash

BRFS=/MidgardOS
echo "Unmounting virtual filesystems from $BRFS"
for mp in "dev/shm" "dev/pts" "sys" "proc" "run" "dev"; do
  if mountpoint -q "$BRFS/$mp"; then
    umount -v "$BRFS/$mp"
  fi
done
