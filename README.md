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