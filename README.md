# ASCII-ART Project

## Overview

The ASCII-ART project is a Go-based utility designed to convert text strings into graphical representations using ASCII characters. This project provides a versatile tool for handling numbers, letters, spaces, special characters, and newline characters.

## Features

- ASCII representation for numbers, letters, spaces, and special characters.
- Support for newline characters.
- Integration with various optional projects.

## Optional Projects

### Reverse

The reverse feature complements the main ASCII-ART project. It converts graphical representations back into their original text form. The user provides a file containing the ASCII template, and the program generates the corresponding text.

#### Usage

```sh
$ go run . --reverse=<fileName>
```

# Color
The color feature enhances the ASCII-ART output by allowing users to add color to specific letters or the entire string. Users can specify the desired color using various notation systems (e.g., RGB, HSL, ANSI).

#### Usage
```sh
Copy code
$ go run . --color=<color> <letters to be colored> [STRING]
```
# Output
The output feature saves the ASCII-ART result to a file. Users can specify the desired file name using the --output flag.

#### Usage
```sh
Copy code
$ go run . --output=<fileName.txt> [STRING] [BANNER]
```
# Fs
The fs feature extends the functionality of the ASCII-ART project by allowing users to select from a range of predefined graphical templates. Users provide a string and choose a template to generate the corresponding ASCII representation.

#### Usage
```sh
Copy code
$ go run . [STRING] [BANNER]
```

# Justify
The justify feature allows users to change the alignment of the ASCII-ART output. Users can choose from options like center, left, right, and justify to customize the display.

#### Usage
```sh
Copy code
$ go run . --align=<type> [STRING] [BANNER]
```
