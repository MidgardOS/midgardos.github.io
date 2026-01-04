#!/bin/sh

BRFS=/MidgardOS
cd $BRFS
sudo setfacl -bPR ./
sudo chown --from builder -hRv root:root ./
cd -
