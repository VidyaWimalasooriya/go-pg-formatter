# go-pgfmt

[![Go](https://img.shields.io/badge/Go-1.23-blue)](https://golang.org/)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)

![MyLibrary](https://github.com/VidyaWimalasooriya/go-pgfmt/blob/main/logo.png)

`go-pgfmt` is a Go library and command-line tool for formatting SQL or PostgreSQL files or queries.

This project is a lightweight go lang wrapper around [pgFormatter](https://github.com/darold/pgFormatter) CLI tool. It requires Perl to be installed on the host system for execution.

## Installation

To install `go-pgfmt`, you can use `go get`:

```sh
go get github.com/VidyaWimalasooriya/go-pgfmt
```

or 

```sh
go install github.com/VidyaWimalasooriya/go-pgfmt@latest
```
you need to `Perl` install in your system.

## Usage
You can use `go-pgfmt` as a command-line tool to format SQL files or queries. Below are some examples of how to use the tool.

```sh
go-pgfmt --destination-path <path> [flags]
```

```sh
go-pgfmt --destination-path ./sql --spaces 2 --keyword-case 1
```

## Flags
The following flags are available for customizing the formatting:

```code
    -d, --destination-path string: Destination file path (required).
    -a, --anonymize: Useful to hide confidential data before formatting.
    -b, --comma-start: In a parameters list, start with the comma.
    -e, --comma-end: In a parameters list, end with the comma.
    -u, --keyword-case int: Change the case of the reserved keyword. Default is uppercase: 2. Values: 0=>unchanged, 1=>lowercase, 2=>uppercase, 3=>capitalize.
    -X, --no-rcfile: Don't read rc files automatically (./.pg_format or $HOME/.pg_format or $XDG_CONFIG_HOME/pg_format).
    -s, --spaces int: Change space indent, default 4 spaces.
    -n, --nocomment: Remove any comment from SQL code.
    -T, --tabs: Use tabs instead of space characters, when used spaces is set to 1 whatever is the value set to -s.
    -B, --comma-break: In insert statement, add a newline after each comma.
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request if you have any improvements or bug fixes.

## Acknowledgements
Special thanks to the contributors and the open-source community for their support and contributions.
