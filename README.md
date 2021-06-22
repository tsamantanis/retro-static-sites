# 📻 Retro Static Sites

[![Go Report Card](https://goreportcard.com/badge/github.com/tsamantanis/retro-static-sites)](https://goreportcard.com/report/github.com/tsamantanis/retro-static-sites)

### 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Getting Started](#getting-started)

## Project Structure

```bash
📂 retro-static-sites
├── README.md
├── examples
    ├── txt-example
        ├── cosmos-first-post.txt
        ├── cosmos-second-post.txt
        ├── cosmos-third-post.txt
    ├── md-example
        ├── first-post.md
        ├── second-post.md
        ├── readme-example.md
├── retro-static-sites.go
├── styles.css
└── template.tmpl
```

## Getting Started

1. Visit [https://github.com/tsamantanis/retro-static-sites](https://github.com/tsamantanis/retro-static-sites) and create a new repository named `retro-static-sites`.
2. Run each command line-by-line in your terminal to set up the project:

```bash
$ git clone git@github.com:tsamantanis/retro-static-sites.git
$ cd retro-static-sites
```

### Single File 📄
To generate a retro styled static site from a single file use:

```bash
$ go run retro-static-sites.go --file=PATH_TO_FILE 
```

replacing `PATH_TO_FILE` with the relative path to a `.md` or `.txt` file of your choosing

### Folder 📂
To generate a retro styled static site from a single file use:

```bash
$ go run retro-static-sites.go --dir=PATH_TO_DIRECTORY 
```

replacing `PATH_TO_DIRECTORY` with the relative path to a directory which contains `.md` or `.txt` files. The directory can contain sub-directories.