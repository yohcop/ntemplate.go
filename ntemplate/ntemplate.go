package ntemplate

import (
  "bytes"
  "io"
  "log"
  "os"
  "path"
  "template"
)

type NTemplate struct {
  File string // template file
  Vars map[string]interface{}
  conf *Config
}

type Config struct {
  Basedir string
  Cache   bool
  cache   map[string]*template.Template
}

func (conf *Config) Template(path string) *NTemplate {
  return &NTemplate{
    File: path,
    Vars: make(map[string]interface{}),
    conf: conf}
}

func (t *NTemplate) tpl() (*template.Template, os.Error) {
  s := t.conf
  path := path.Join(s.Basedir, t.File)
  if s.Cache && s.cache == nil {
    s.cache = make(map[string]*template.Template)
  }
  if s.Cache {
    tpl, ok := s.cache[path]
    if ok {
      return tpl, nil
    }
  }
  tpl, err := template.ParseFile(path, nil)
  if err != nil {
    return nil, err
  }
  if s.Cache {
    s.cache[path] = tpl
  }
  return tpl, nil
}

func (s *NTemplate) Execute(wr io.Writer) os.Error {
  str, err := s.Render()
  if err != nil {
    log.Println(err)
    return err
  }
  wr.Write([]byte(str))
  return nil
}

func (s *NTemplate) Render() (string, os.Error) {
  tpl, err := s.tpl()
  if err != nil {
    return "", err
  }
  tplMap := make(map[string]interface{})

  for sk, el := range s.Vars {
    tplMap[sk] = s.stringify(el)
  }
  var a bytes.Buffer
  tpl.Execute(&a, tplMap)
  return a.String(), nil
}

func (s *NTemplate) stringify(v interface{}) interface{} {
  switch vt := v.(type) {
  case *NTemplate:
    str, _ := vt.Render()
    return str
  }
  return v
}
