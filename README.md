# Github Binary Manager

Manage installation, update and removal of binaries from Github releases.

## To-do
 - [X] Core
   - [X] get release info from Github API
   - [X] get release info from Github CLI
   - [X] check update
   - [X] download binary to path
   - [X] download archive and extract
 - [ ] CLI
   - [X] check update
   - [X] download binary to path
   - [X] download archive and extract
   - [ ] instruction set
 - [ ] Checksum validation

## Instructions

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

### print

- Print 1 or multiple values from stack
- Without param, a single value will be peaked and printed
- 

```
print
print 3
```