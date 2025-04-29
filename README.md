# gosubs

A simple tool for generating subtitles using [OpenAI Whisper](https://huggingface.co/openai/whisper-base) 📔

## Building for local development

Prerequisites:

1. Linux
2. Git
3. Docker

Firstly, clone the repository into your `$GOPATH`:

```shell
mkdir -p $GOPATH/src/github.com/nightnoryu
cd $GOPATH/src/github.com/nightnoryu

git clone --recurse-submodules git@github.com:nightnoryu/gosubs.git
cd gosubs
```

Then build the binary:

```shell
bin/gosubsbrewkit build
```

This script will download a [brewkit build system](https://github.com/ispringtech/brewkit) binary and put it in
the `bin` directory of the project. The build is entirely dockerized, so you don't have to install other dependencies.
First build may take a while, because we're building whisper.cpp bindings for Go, but subsequent builds will be much
faster, harnessing the power of brewkit image caching.

_STILL UNDER CONSTRUCTION_
