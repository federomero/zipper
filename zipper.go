package zipper

import (
  "archive/zip"
  "bytes"
  "net/http"
  "io"
  "strings"
)

type download struct {
  url  string
  body io.ReadCloser
  err  error
}

func CreateZip(urls []string, fo io.Writer) error{

  buf := new(bytes.Buffer)

  archive := zip.NewWriter(buf)

  c := make(chan download, len(urls))

  for _, url := range urls{
    go downloadURL(url, c)
  }

  for i := 0; i < len(urls); i++ {
    d := <-c
    if d.err != nil { return d.err }
    fileName := fileNameFromURL(d.url)
    w, err := archive.Create(fileName)
    if err != nil { return err }
    _, err = io.Copy(w, d.body)
  }

  err := archive.Close()
  if err != nil { return err }

  _, err = buf.WriteTo(fo)
  return err
}

func fileNameFromURL(url string) string{
  a := strings.Split(url, "/")
  return a[len(a)-1]
}

func downloadURL(url string, c chan download){
  resp, err := http.Get(url)
  if err != nil {
    c <- download{url, nil, err}
  }else{
    c <- download{url, resp.Body, err}
  }

}
