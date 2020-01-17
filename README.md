# GFTP - Go SFTP

GFTP is a simple SFTP (Simple File Transfer Protocol) client written in Go.

The goal of this application is purely educational and for me to challenge myself 
by using languages I am not as fluent in at the time of writing, or to learn 
concepts I haven't necessarily dove into in the past. 

That said, I am modeling this client loosely off of the BSD `sftp` command-line 
utility, and am developing on Kali Linux 5.3.0 x86_64.

---

## Install with Go

To build GFTP from source, run the following Go command:

```
go install github.com/zcking/gftp
```

This will install the `gftp` binary into your Go installation's `bin/` directory, 
which should already be in your system's `PATH` environment variable. 

After running the install command, simply run `gftp` to test that the command is found. 

