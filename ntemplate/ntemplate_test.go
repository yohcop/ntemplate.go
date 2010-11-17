package ntemplate

import "testing"

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
  sub2 := MakeTemplate("sub/sub2.tpl", conf)
  sub1 := MakeTemplate("sub/sub1.tpl", conf)
  sub1.Vars["sub"] = sub2

  b := MakeTemplate("t2.tpl", conf)
  b.Vars["var1"] = sub1

  a := MakeTemplate("t1.tpl", conf)
  a.Vars["head"] = "<title>Hello NTemplate</title>"
  a.Vars["b"] = b
  a.Vars["f"] = 12.3456789
  a.Vars["s"] = "a string"
  s, _ := a.Render()
  if s != expected {
    t.Error(s + ".\nexpected\n" + expected + ".")
  }
}

