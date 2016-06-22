# doenter (Docker Enter)

**doenter** is an utility that allows you to **obtain a shell** inside the [Docker for Mac Beta](https://blog.docker.com/2016/03/docker-for-mac-windows-beta/) **xhyve** virtual machine.

A shell inside that Virtual Machine is useful because:

- Customize configurations or Docker Daemon options
- Use a custom way to share files with the host instead of osxfs
- Personalize networks
- Ideas to come


# Installation

## Requirements

- A working and running [Docker for Mac](https://docs.docker.com/engine/installation/mac/)

```
curl -L https://github.com/fntlnz/doenter/releases/download/0.1.0/doenter > /usr/local/bin/doenter
chmod +x /usr/local/bin/doenter
```

# Usage

Once installed, `doenter` is ready to be used.

Just run `doenter` in your terminal and you will be in the virtual machine's shell.

*IMPORTANT*: you will be asked for username and password. Username is `root` and password is empty.

```
Welcome to Moby alpha
Kernel 4.4.13-moby on an x86_64 (/dev/ttyS0)

                        ##         .
                  ## ## ##        ==
               ## ## ## ## ##    ===
           /"""""""""""""""""___/ ===
      ~~~ {~~ ~~~~ ~~~ ~~~~ ~~~ ~ /  ===- ~~~
           \______ o           __/
             \    \         __/
              \____\_______/

moby login:
```

# Plans

- [ ] Tests
- [ ] Support the Windows version
- [ ] Inject scripts at boot or change the virtual machine image


# Credits

- [ftorn](https://github.com/ftorn) : he suggested the tty approach and helped a lot in general as always :D
- [atosatto](https://github.com/atosatto) : he's always very helpful

