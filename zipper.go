package zipper

import (
  "archive/zip"
  "bytes"
  "net/http"
  "io"
  "strings"
)

func CreateZip(urls []string, fo io.Writer){

  buf := new(bytes.Buffer)

  w := zip.NewWriter(buf)

  for _, url := range urls{
    fileName := fileNameFromURL(url)
    f, err := w.Create(fileName)
    if err != nil { panic(err) }
    appendFile(f, url)
  }

  err := w.Close()
  if err != nil { panic(err) }

  _, err = buf.WriteTo(fo)
  if err != nil { panic(err) }
}

func fileNameFromURL(url string) string{
  a := strings.Split(url, "/")
  return a[len(a)-1]
}

func appendFile(archive io.Writer, url string){
  resp, err := http.Get(url)
  defer resp.Body.Close()
  if err != nil { panic(err) }
  _, err = io.Copy(archive, resp.Body)
  if err != nil { panic(err) }
}
