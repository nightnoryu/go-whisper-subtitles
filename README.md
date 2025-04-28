# gosubs

A simple tool for generating subtitles using [OpenAI Whisper](https://huggingface.co/openai/whisper-base) 📔

## Building for local development

Prerequisites:

1. Linux
2. Git

Firstly, clone the repository:

```shell
git clone git@github.com:nightnoryu/gosubs.git
```

Then build the binary:

```shell
./bin/gosubsbrewkit build
```

This script will download a [brewkit build system](https://github.com/ispringtech/brewkit) binary and put it in the `bin` directory of the project.

## Tasks to do

- [ ] Add model loading from the web
- [ ] Provide `libwhisper.a` and header files
- [ ] Add Makefile with proper linking
- [ ] Remove hardcoded filenames, use temporary directory with randomly generated filenames and remove leftovers
- [ ] Add CLI parameters
  - to mux the subtitles file or not + SRT output filename
  - specifying output container format (mkv/mp4/webm)
- [ ] Make fancy CLI output & logging
