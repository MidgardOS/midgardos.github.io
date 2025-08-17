#!/bin/bash

for DIR in "etc" "dev" "home" "mnt" "media" "opt" \
           "proc" "root" "run" "selinux" "srv" "sys"; do
    install -v -d -m 755 -o root -g root "/MidgardOS/$DIR"
done

for DIR in "usr/bin" "usr/include" "usr/lib" "usr/lib64" \
           "usr/sbin" "usr/share" "usr/src"; do
    install -v -d -m 755 -o root -g root "/MidgardOS/$DIR"
done

for DIR in "usr/local/bin" "usr/local/include" "usr/local/lib" \
           "usr/local/lib64" "usr/local/sbin" "usr/local/share" \
           "usr/local/src"; do
    install -v -d -m 755 -o root -g root "/MidgardOS/$DIR"
done

for DIR in "var/adm" "var/cache" "var/crash" "var/games" "var/lib" \
           "var/log" "var/opt" "var/spool" "var/yp"; do
    install -v -d -m 755 -o root -g root "/MidgardOS/$DIR"
done

for DIR in "tmp" "var/tmp" "var/spool/mail"; do
    install -v -d -m 1777 -o root -g root "/MidgardOS/$DIR"
done

# now build thhe symbolic links
pushd /MidgardOS
    for LINKED_DIR in "bin" "lib" "lib64" "sbin"; do
        ln -fsv "usr/$LINKED_DIR" "$LINKED_DIR"
    done
popd

# compat links
pushd /MidgardOS/var
    ln -fsv ../run run
    ln -fsv ../run/lock lock
    ln -fsv spool/mail mail
popd

# Tools and sources
install -v -d -m 755 -o root -g root /MidgardOS/sources
## link will be added in the next step
install -v -d -m 755 -o root -g root /MidgardOS/tools
ln -fsv /MidgardOS/tools /
