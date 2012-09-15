# Zipper

Simple go package for downloading and zipping files.

## Install

Install zipper by running the following command

    go get github.com/federomero/zipper

## Example

    package main

    import (
      "zipper"
      "os"
    )

    func main() {
      // List of files we want to include in the resulting zip file
      urls := []string{
        "http://raw.github.com/federomero/zipper/master/zipper.go",
        "http://raw.github.com/federomero/zipper/master/README.md",
      }

      fo, err := os.Create("output.zip")
      defer fo.Close()
      if err != nil { panic(err) }
      zipper.CreateZip(urls, fo)
    }
