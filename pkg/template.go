// Copyright (C) 2023  Shovra Nikita, Livitsky Andrey

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sshabu

import (
	"bytes"
	_ "embed"

	"io"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

//go:embed sshabu_example.yaml
var sshabu_example string
//go:embed *.gtpl
var ssh_template string

func ConfigExample() string{
  return sshabu_example
}

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
