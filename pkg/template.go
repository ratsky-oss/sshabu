package sshabu

import (
	"bytes"
	_ "embed"

	"io"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

//go:embed *.gtpl
var ssh_template string

func RenderTemplate(vars interface{}, out io.Writer) error {
  t := template.New("openssh_config")
  
  var funcMap template.FuncMap = map[string]interface{}{}
  // copied from: https://github.com/helm/helm/blob/8648ccf5d35d682dcd5f7a9c2082f0aaf071e817/pkg/engine/engine.go#L147-L154
  funcMap["include"] = func(name string, data interface{}) (string, error) {
    buf := bytes.NewBuffer(nil)
    if err := t.ExecuteTemplate(buf, name, data); err != nil {
      return "", err
    }
    return buf.String(), nil
  }

  t, err := t.Funcs(sprig.TxtFuncMap()).Funcs(funcMap).Parse(ssh_template)
  if err != nil {
      return err
  }
  
  err = t.Execute(out, &vars)
  if err != nil {
      return err
  }
  return nil
}
