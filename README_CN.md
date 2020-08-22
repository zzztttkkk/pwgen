# pwgen
- **你必须修改'[SECRET_KEY](https://github.com/zzztttkkk/pwgen/blob/master/secret.go#L4)'，然后重新编译**
- **'[SECRET_KEY](https://github.com/zzztttkkk/pwgen/blob/master/secret.go#L4)'和密码生成无关，只用来加密文件**
- **为了更安全，你可以将'[SECRET_KEY](https://github.com/zzztttkkk/pwgen/blob/master/secret.go#L4)'改为空字符串，这样每次运行都会询问**
- **密码不会保存到文件中**
- **password = remap(sha512(`用户名`:`密钥`@`hostname`))**
- **你可以保存多个账户，每个账户都包含一个`用户名`和一个`密钥`**
- **保存的这些账户，只是用来生成密码的，和你真实网络中的真实账户没有任何关系**

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