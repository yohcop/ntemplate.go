package ntemplate

import (
  "testing"
)

var expected=`<html>
<title>Hello NTemplate</title>
<body>
<div><span>Hello
</span>
</div>

12.345679
a string
</body>
</html>
`

func TestNTemplate(t *testing.T) {
  conf := &Config{Basedir: "testdata", Cache: false}
  sub2 := conf.Template("sub/sub2.tpl")
  sub1 := conf.Template("sub/sub1.tpl")
  sub1.Vars["sub"] = sub2

  b := conf.Template("t2.tpl")
  b.Vars["var1"] = sub1

  a := conf.Template("t1.tpl")
  a.Vars["head"] = "<title>Hello NTemplate</title>"
  a.Vars["b"] = b
  a.Vars["f"] = 12.3456789
  a.Vars["s"] = "a string"
  s, err := a.Render()
  if s != expected {
    t.Error(s + ".\nexpected\n" + expected + ".")
  }
  if err != nil {
    t.Error("Error while rendering: " + err.String())
  }
}
