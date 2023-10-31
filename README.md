## ascii-art by bbilisbe and ayessenb

### Objectives

Ascii-art FS is a program which consists in receiving a `string` as an argument and a banner format as an option, then outputting the `string` in a graphic representation using ASCII in a given format. 


- This project should handle an input with numbers, letters, spaces, special characters and `\n`.
- Usage: go run . [STRING] [BANNER]
- Ascii-art-output project implemented, the program accept the correctly formatted [OPTION] [STRING] [BANNER].
- The program must still able to run with a single [STRING] argument.


### Banner Format

* shadow
* standard
* thinkertoy


### Usage

```console
student$ go run . "" | cat -e
student$ go run . "\n" | cat -e
$
student$ go run . "Hello\n" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
student$ go run . "hello"  shadow | cat -e
 
                                 $
_|    _|          _| _|          $
_|    _|   _|_|   _| _|   _|_|   $
_|_|_|_| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $
$
student$ go run . "hello"  thinkertoy | cat -e

                 $
o  o     o o     $
|  |     | |     $
O--O o-o | | o-o $
|  | |-' | | | | $
o  o o-o o o o-o $
                 $
                 $
$

student$
```

### Allowed packages

- Only the [standard Go](https://golang.org/pkg/) packages are allowed

This project will help you learn about :

- The Go file system(**fs**) API
- Data manipulation