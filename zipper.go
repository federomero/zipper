package zipper

import (
  "archive/zip"
  "bytes"
  "net/http"
  "io"
  "strings"
)

type URLDownload struct {
  url  string
  body io.ReadCloser
  err  error
}

// Download all files from the given urls
// and write them on the output
func CreateZip(urls []string, output io.Writer) error{

  buf := new(bytes.Buffer)

  archive := zip.NewWriter(buf)

  c := make(chan URLDownload, len(urls))

  for _, url := range urls{
    go downloadURL(url, c)
  }

  for i := 0; i < len(urls); i++ {
    err := appendToArchive(archive, <-c)
    if err != nil { return err }
  }

  err := archive.Close()
  if err != nil { return err }

  _, err = buf.WriteTo(output)
  return err
}

func fileNameFromURL(url string) string{
  a := strings.Split(url, "/")
  return a[len(a)-1]
}

func downloadURL(url string, c chan URLDownload){
  resp, err := http.Get(url)
  if err != nil {
    c <- URLDownload{url, nil, err}
  }else{
    c <- URLDownload{url, resp.Body, err}
  }

}

func appendToArchive(archive *zip.Writer, download URLDownload) error{
  if download.err != nil { return download.err }
  fileName := fileNameFromURL(download.url)
  w, err := archive.Create(fileName)
  if err != nil { return err }
  _, err = io.Copy(w, download.body)
  defer download.body.Close()
  return nil
}
