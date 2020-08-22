# pwgen [中文](https://github.com/zzztttkkk/pwgen/blob/master/README_CN.md)
- **you must change '[SECRET_KEY](https://github.com/zzztttkkk/pwgen/blob/master/secret.go#L4)', then recompile**
- **the '[SECRET_KEY](https://github.com/zzztttkkk/pwgen/blob/master/secret.go#L4)' does not participate in the generation of the password and is only used to encrypt files**
- **for more security, you can set the '[SECRET_KEY](https://github.com/zzztttkkk/pwgen/blob/master/secret.go#L4)' to empty, so you will be asked every time**
- **password not store in file and will be recalculated every time**
- **password = remap(sha512(`username`:`secret`@`hostname`))**
- **you can save many accounts, each containing a `username` and a `secret key`**
- **these accounts are just for generating passwords, and have nothing to do with real accounts on the real internet**

# help
```
pwgen: generate password by remap(sha512(username:secret@hostname)).
        -u : username
        -h : hostname

    add <username string> <-d bool> <-u string>
        save a account to the filesystem("~/.pwgen").
        -d : this account will be the default
	
    del <username string> <-u string>
        delete a account from the filesystem.

    ls
        list all saved accounts.

    gen <hostname string> <username string> <-l int> <-p bool> <-u string> <-h string> <-s bool> <-d bool>
        generate a password. if 'username' is empty, use the default account. this is the default sub-command.
        -l : length of the password, in range (1-64)
        -p : do not send to the clipboard, only print. 
        -s : auto save this account.
        -d : auto save this account as the default.
```