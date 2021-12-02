# Drive2 Images Downloader #

Allows to parse and download all the images that are in a post (logbook, community)

## Installation ##

```shell
git clone https://github.com/gravityperson/parser-drive2ru-images-downloader.git
```

## Building ##

### Unix ###

```shell
cd parser-drive2ru-images-downloader
make
```

## Running ##

### Required Arguments ###

> (string) "-post-url" an url of a post on drive2 platform

### Non-Required Arguments ###

> (string) (info) (debug, info, error, fatal) "-log-level" a log level for the application
> (string) (dist) "-dist" a destination folder of the images

### Unix ###

```shell
cd ./bin/unix
./downloader
```

## Example ##

The following snippet allows to download all the images from [the page](https://www.drive2.ru/c/600355065194100870/)
into `dist` folder

### Unix ###

```shell
./downloader -post-url https://www.drive2.ru/c/600355065194100870/
```

After the run completes you can see images within "dist" folder
