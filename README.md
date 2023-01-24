# goctail
Colorful configurable log output

![image](https://user-images.githubusercontent.com/221943/214445920-42343faa-2a34-4f99-bdb5-8f42a8d8074a.png)

## Usage

```
$ ./goctail /var/log/dpkg.log
2023-01-18 16:18:22 trigproc libc-bin:amd64 2.35-0ubuntu3.1 <none>
2023-01-18 16:18:22 status half-configured libc-bin:amd64 2.35-0ubuntu3.1
2023-01-18 16:18:22 status installed libc-bin:amd64 2.35-0ubuntu3.1
2023-01-18 16:18:22 trigproc man-db:amd64 2.10.2-1 <none>
2023-01-18 16:18:22 status half-configured man-db:amd64 2.10.2-1
2023-01-18 16:18:22 status installed man-db:amd64 2.10.2-1
2023-01-18 16:18:22 trigproc ca-certificates:all 20211016ubuntu0.22.04.1 <none>
2023-01-18 16:18:22 status half-configured ca-certificates:all 20211016ubuntu0.22.04.1
2023-01-18 16:18:22 status installed ca-certificates-java:all 20190909
2023-01-18 16:18:22 status installed ca-certificates:all 20211016ubuntu0.22.04.1
$ ./goctail -n 3 /var/log/dpkg.log
2023-01-18 16:18:22 status half-configured ca-certificates:all 20211016ubuntu0.22.04.1
2023-01-18 16:18:22 status installed ca-certificates-java:all 20190909
2023-01-18 16:18:22 status installed ca-certificates:all 20211016ubuntu0.22.04.1
$ ./goctail -c 8 /var/log/dpkg.log
22.04.1
```

## Features to implement

* Basic compatibility wih tail
* Colours
* Simply HTTP server
* Smart packing logs to zip for futher forensic
* Config files

## Other tail implementations

* [coreutils tail - c](https://github.com/coreutils/coreutils/blob/master/src/tail.c)
* [multitail - c](https://github.com/halturin/multitail)

## Colorizers

* [grc](https://github.com/garabik/grc  )
