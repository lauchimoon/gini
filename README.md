# gini
Small INI library made in Go.

## Getting started
In your project directory, run
```sh
$ go get github.com/lauchimoon/gini
```

## Usage
```ini
; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.

[database]
; use IP address in case network name resolution is not working
server = 192.0.2.62     
port = 143
file = "payroll.dat"
```

### Loading
You can load an INI file into a map with `gini.LoadFromFile` as such:
```go
ini, err := gini.LoadFromFile("test.ini")
```

Alternatively, you can use `gini.LoadFromString`:
```go
ini, err := gini.LoadFromString("[owner]\nname = John Doe\norganization = Acme Widgets Inc.")
```

### Reading
Once the INI file is loaded, you can inspect the value of a key in a section in two ways:
```go
v, err := ini.Get("owner", "name")
```
or,

```go
v := ini["owner"]["name"]
```
The only difference is that `ini.Get` internally checks for both section and key existence.

### Writing
gini provides a `Dump` function which prints the given INI into a file, as such:
```go
ini["owner"]["name"] = "Jane Doe"
gini.Dump(ini, os.Stdout)
```

## License
This library is free software; you can redistribute it and/or modify it under
the terms of the MIT license. See [LICENSE](LICENSE) for details.
