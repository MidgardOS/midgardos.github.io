| Navigation |||
| --- | --- | ---: |
| [<<](./ManageDisk.md) | [HOME](./README.md) | [>>](./CreateBuildUser.md) |

# Creating Directories and Symbolic Links

As said earlier, MidgardOS is a traditional Linux server distribution that follows the File Hierarchial Standard. With that said, it also follows the current practice in most distributions of merging many of the common directories from the root of the filesystem of the OS to the `/usr` hierarchy, such as `/bin`, `/lib`, `/lib64`, and `/abin`. To assist in ensuring the builds install correctly with this arrangement, we'll build the basic directory structure on the newly mounted filesystems.

To do so, run the following commands:

```bash
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
           "var/log" "var/opt" "var/spool" "yp"; do
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
```

| Navigation |||
| --- | --- | ---: |
| [<<](./ManageDisk.md) | [HOME](./README.md) | [>>](./CreateBuildUser.md) |