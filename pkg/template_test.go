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

// Import necessary packages for testing
import (
    "bytes"
    "os"
    "path/filepath"
    // "strings"
    "testing"

    "github.com/spf13/viper"
)
func readExpectedOutput(t *testing.T, filePath string) string {
    content, err := os.ReadFile(filePath)
    if err != nil {
        t.Fatal(err)
    }
    return string(content)
}
func TestConfigExample(t *testing.T) {
    tests := []struct {
        name string
        want string
    }{
        {
            name: "Example Config Test",
            want: readExpectedOutput(t, "sshabu_example.yaml"),
        },
        // Add more test cases if needed
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ConfigExample(); got != tt.want {
                t.Errorf("ConfigExample() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRenderTemplate(t *testing.T) {
    // Get the current directory
    currentDir, err := os.Getwd()
    if err != nil {
        t.Fatal(err)
    }

    // Create the path to sshabu_example.yaml
    yamlFilePath := filepath.Join(currentDir, "sshabu_example.yaml")

    // Set up Viper with the YAML file
    viper.SetConfigFile(yamlFilePath)
    err = viper.ReadInConfig()
    if err != nil {
        t.Fatal(err)
    }

    // Unmarshal YAML into Shabu
    var shabu Shabu
    err = viper.UnmarshalExact(&shabu)
    if err != nil {
        t.Fatal(err)
    }

    // Call Boil
    err = shabu.Boil()
    if err != nil {
        t.Fatal(err)
    }

    // Create a buffer for RenderTemplate
    buf := new(bytes.Buffer)

    // Render the template
    err = RenderTemplate(shabu, buf)
    if err != nil {
        t.Fatal(err)
    }

    // Optionally, you can add assertions based on the expected output of RenderTemplate

    // Example assertion:
    expectedOutput := `# -----------------------
# RATSKY SSHABU


Host project1-test
    Hostname 192.168.1.2
    IdentityFile ~/.ssh/id_rsa_work_p1
    Port 2222
    User user
    
Host project2-dev
    Hostname 192.168.11.3
    IdentityFile ~/.ssh/id_rsa_work
    User user
    
Host home-gitlab
    Hostname gitlab.ratsky.local
    `
    os.WriteFile("test_case", buf.Bytes(), 0600)
    if gotOutput := buf.String(); gotOutput != expectedOutput {
        t.Errorf("Rendered output does not match. \nGot: \n%s|\n Want: \n%s|", gotOutput, expectedOutput)
    }
}