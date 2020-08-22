# pwgen

**you must change '[SECRET_KEY](https://github.com/zzztttkkk/pwgen/blob/master/secret.go#L4)', then recompile.**

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

    clean
        remove the 
```