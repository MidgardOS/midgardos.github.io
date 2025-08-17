#!/bin/bash

BRFS=/MidgardOS
pushd $BRFS
sudo setfacl -bPR ./
sudo chown --from builder -hRv root:root ./
popd
