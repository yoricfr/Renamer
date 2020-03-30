# Renamer in Go

## Purpose

This CLI program run from the command line and will rename all the files in the search folder, following the renaming rules in the Excel file sitting in the root directory.

## Before running the script

    renamer/
    ├── LICENSE
    ├── README.md
    ├── rules.xlsx
    ├── renamer.go
    └── search/
        ├── 1
        |   └── file001.jpg
        └── 2
            └── file002.jpg

rules.xlsx:

    file001.jpg ; newname001.jpg
    file002.jpg ; anynewname.jpg
    
## After running the script, we get:

    renamer/
    ├── LICENSE
    ├── README.md
    ├── rules.xlsx
    ├── renamer.go
    └── search/
        ├── 1
        |   └── newname001.jpg
        └── 2
            └── anynewname.jpg

