# Github Binary Manager

Install and update binaries from Github releases.

## Command Line

### Top-level Commands

| command | description |
|---|---|
| info | get release information |
| check | check for update |
| install | install binary from latest release |
| installx | install binary from latest release archive |
| code | run instruction set for multiple operation |

### Command Usage

- `info`

```
Usage of ghbin-mgr info:
  -json
        print output to json
  -p string
        pattern to filter asset (comma separated, case-insensitive)
  -r string
        github repository name
  -u string
        github user name
```

- `check`

```
Usage of ghbin-mgr info:
  -r string
        github repository name
  -t string
        release tag to compare with
  -u string
        github user name
```

- `install`

```
Usage of ghbin-mgr install:
  -d string
        installation directory
  -n string
        binary
  -p string
        pattern to filter asset (comma separated, case-insensitive)
  -r string
        github repository name
  -u string
        github user name
```

- `installx`

```
Usage of ghbin-mgr install:
  -d string
        installation directory
  -n string
        binary
  -p string
        pattern to filter asset (comma separated, case-insensitive)
  -px string
        pattern to filter binary file in archive (comma separated, case-insensitive)
  -r string
        github repository name
  -u string
        github user name
```

- `code`

```Usage of ghbin-mgr code:
  -f string
        instruction file path
```

### Patterns

- Patterns can be sub-string
- Multiple patterns can be added separated by comma (,)
- Special templates: these templates can be added as a pattern, which would be replaced by their corresponding values
  - `__os__` replaced with operating system name (windows, darwin, linux)
  - `__arch__` replace with architecture (x86_64, darwin, linux etc.)
- Exclusions can be added by adding `^` at the starting of a pattern (does not work with templates)

## Instructions

- Empty lines are permitted
- Lines starting with `##` are considered as comments

### push

- Push 1 or multiple values to stack.
- For multiple values, use space ` ` as separator
- Use double quote for string `"`
- When pushing multiple values, the params will be pushed from right to left
```
push "this is a test" 1 3.4 "another test"
push 1
```

### pop

- Pop 1 or multiple values to stack
- Without param, a single value will be poped
- 

```
pop
pop 3
```

### label

- Create a label which can be used to jump to that label
- Label name should start and end with `:`

```
label :do_this:
```

### goto

- Jump to a label
- Jump instruction will push the next instruction position in the return stack
- Label can be defined anywhere in the instruction list
- Label name should start and end with `:`

```
goto :do_this:
```

### return

- Pop from the return stack and return to that instruction position
- Does not take any parameters

```
return
```

### print

- Print 1 or multiple values from stack
- Without param, a single value will be peaked and printed

```
print
print 3
```

### exit

- Exits the instruction execution loop, effectivly exiting the vm
- Does not take any parameters

```
exit
```

### jumpeq

- Jump to a label if last 2 values in the stack are equal
- Jump instruction will push the next instruction position in the return stack
- Label can be defined anywhere in the instruction list
- Label name should start and end with `:`

```
jumpeq :do_this:
```

### jumpeqn

- Jump to a label if last 2 values in the stack are not equal
- Jump instruction will push the next instruction position in the return stack
- Label can be defined anywhere in the instruction list
- Label name should start and end with `:`

```
jumpeqn :do_this:
```

### kvload

- Loads a file as key-value storage
- Requires only one parameter
- Parameter must be a location for storage file
- Of file does not exist at the location, a new file will be created
- Throws error if the file is not a valid json file

```
kvload "/location/of/a/file.json"
```

### kvsave

- Saves the key-value storage
- No paramerters required

```
kvsave
```

### kvget

- Gets value of a key from storage and pushes it into the stack
- Requires key as parameter

```
kvget "some_key"
```

### kvset

- Sets value of a key from storage with the last value from the stack
- Requires key as parameter

```
kvset "some_key"
```

### kvdelete

- Deletes the value of a key from storage
- Requires key as parameter

```
kvdelete "some_key"
```

### ghcheck

- Check for latest updated release for a github repository
- Necessary params are picked from stack
- Params should match the following position in stack-
  - [2] tag
  - [1] github username
  - [0] repository name

```
# add paramas one-by-one
push "tag"
push "user"
push "repo"
ghckeck

# add paramas single instruction, values pushed from left-to-right
push "repo" "user" "tag"
ghckeck
```

### ghinstallx

- Download an archive file from github release and extract binary
- Necessary params are picked from stack
- Params should match the following position in stack-
  - [5] installation directory
  - [4] pattern to find binary in the archive
  - [3] pattern to find archive in release assets
  - [2] binary name
  - [1] github username
  - [0] repository name

```
# add paramas one-by-one
push "install_dir"
push "binary_in_archive_pattern"
push "asset_archive_pattern"
push "binary_name"
push "user"
push "repo"
ghinstallx

# add paramas single instruction, values pushed from left-to-right
push "repo" "user" "binary_name" "asset_archive_pattern" "binary_in_archive_pattern" "install_dir"
ghinstallx
```

### ghinstall

- Download a binary from github release
- Necessary params are picked from stack
- Params should match the following position in stack-
  - [4] installation directory
  - [3] pattern to find binary in release assets
  - [2] binary name
  - [1] github username
  - [0] repository name

```
# add paramas one-by-one
push "install_dir"
push "binary_file_pattern"
push "binary_name"
push "user"
push "repo"
ghinstall

# add paramas single instruction, values pushed from left-to-right
push "repo" "user" "binary_name" "binary_file_pattern" "install_dir"
ghinstall
```

## To-do
 - [X] Core
   - [X] get release info from Github API
   - [X] get release info from Github CLI
   - [X] check update
   - [X] download binary to path
   - [X] download archive and extract
 - [X] CLI
   - [X] check update
   - [X] download binary to path
   - [X] download archive and extract
   - [X] instruction set
 - [ ] Checksum validation
