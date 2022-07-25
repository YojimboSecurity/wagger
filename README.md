<h1 align="center">
  <a href="https://github.com/yojimbosecurity/wagger">
    <!-- Please provide path to your logo here -->
    <img src="docs/images/log.jpg" alt="Logo" width="300" height="500">
  </a>
</h1>

<div align="center">
  wagger
  <br />
  <a href="#about"><strong>Explore the docs »</strong></a>
  <br />
</div>

<div align="center">
<br />

![platform](https://img.shields.io/badge/platform-linux--64%20%7C%20win--32%20%7C%20osx--64%20%7C%20win--64-lightgrey)
![Go version](https://img.shields.io/badge/Go-v1.18-blue)

</div>

<details open="open">
<summary>Table of Contents</summary>

- [About](#about)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)

</details>

---

## About

wagger is a command line tool for tailing files. It is a
cross-platform tool that is meant to replace BareTail on Windows, and is a
drop-in replacement for Tail on Linux and Mac. Unlike the Unix tail it has
highlighting and filtering capabilities.

### Built With

- [Go](https://golang.org/)

## Getting Started

To get started, will need to create an executable for your platform.

```bash
make build
```

Or, if you need to build an executable for a specific platform, you can use one
of the following commands:

```bash
make build-windows
make build-linux
make build-darwin
```

You can also create builds for profiling with the following commands:

```bash
make build-profile-mem
make build-prfile-cpu
make build-repe
```

### Prerequisites

- [Go](https://golang.org/)
- [Make](https://man7.org/linux/man-pages/man1/make.1.html)

### Installation

wagger is an executable that is intended to be used with a config file. A basic
config file can be found [here](wagger.yml). This config file **must**
be located in the same directory as the executable.

## Usage

To use wagger, run the executable with the following command:

```bash
wagger
```

This command will read the file and apply highlighting.

If you only want to see highlighted lines, you can use the `--only-highlight` `-o` flag:

```bash
wagger --only-highlight
```

If you want to ignore lines that match, you can use the `--ignore` `-i` flag:

```bash
wagger --ignore 
```

To tail the log file, use the `tail` command:

```bash
wagger tail
```

The same filtering flags apply.
