# doenter (Docker Enter)

**doenter** is an utility that allows you to **obtain a shell** inside the [Docker for Mac](https://docs.docker.com/engine/installation/mac/#/docker-for-mac) **xhyve** virtual machine.

A shell inside that Virtual Machine is useful because:

- Customize configurations or Docker Daemon options
- Use a custom way to share files with the host instead of osxfs
- Personalize networks
- Ideas to come


## Disclaimer

There are chances that for what you are trying to do **you don't need doenter at all!**

Yes, I know that I can access to the virtual machine resources by collapsing into the host namespace using a privileged container.
Doenter is intended to be used in those situations where you don't have the docker daemon running (e.g you are restarting it or substituting it with a brand new one)

So, if you just need to enter into the VM and poke around just do something like this:

```bash
docker run -it --pid=host --privileged debian:jessie nsenter -t 1 -m -p -n
```

# Installation

## Requirements

- A working and running [Docker for Mac](https://docs.docker.com/engine/installation/mac/)

```
curl -L https://github.com/fntlnz/doenter/releases/download/0.1.1/doenter > /usr/local/bin/doenter
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

