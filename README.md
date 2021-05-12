# file2configmap

A tool to convert files into Kubernetes ConfigMap resource with `binaryData` field.

## Install

```shell
go install -u go.guoyk.net/file2configmap
```

## Usage

### Files in current directory

```shell
file2configmap -n myname -ns mynamespace file1.txt
```

### Files in subdirectories

`file2configmap` will use file base name as `ConfigMap` key.

```shell
file2configmap -n myname -ns mynamespace path/to/file2.txt
```

### Customize key name instead of using file name

```shell
file2configmap -n myname -ns mynamespace file3.json:file3.txt
```

### Multiple files

```shell
file2configmap -n myname -ns mynamespace file1.txt path/to/file2.txt files3.json:file3.txt
```

## Credits

Guo Y.K., MIT License
