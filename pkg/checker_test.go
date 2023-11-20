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
	// "io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

var execCommand = exec.Command

func render(t *testing.T) string {
	currentDir, err := os.Getwd()
    if err != nil {
        t.Fatal(err)
    }

    // Create the path to sshabu_example.yaml
    yamlFilePath := filepath.Join(currentDir, "sshabu_example.yaml")
	configFilePath := filepath.Join(currentDir, "sshabu_example.config")
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
	os.WriteFile(configFilePath, buf.Bytes(), 0600)
	return configFilePath
}

// TODO: Do not use sshabu_example.yaml rendering
func TestOpensshCheck(t *testing.T) {
	tests := []struct {
		name    string
		cmd     *exec.Cmd
		wantErr bool
	}{
		{
			name: "Test with successful SSH check",
			cmd:  exec.Command("ssh", "-GTF", "test"),
			// For a real SSH check, you would set cmd: exec.Command("ssh", "-GTF", "fake_config", "test"),
			wantErr: false,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			execCommand = func(name string, arg ...string) *exec.Cmd {
				return tt.cmd
			}
			defer func() { execCommand = exec.Command }()
			path := render(t)
			if err := OpensshCheck(path); (err != nil) != tt.wantErr {
				t.Errorf("OpensshCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
			os.Remove(path)
		})
	}
}

// TODO: Do not use sshabu_example.yaml rendering
func TestDestinationHosts(t *testing.T) {
	// type args struct {
	// 	r io.Reader
	// }
	tests := []struct {
		name    string
		// args    args
		want    []string
		wantErr bool
	}{
		{
			name: "CRAP",
			want: []string{
				"project1-test",
				"project2-dev",
				"home-gitlab",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := render(t)

			file, _ := os.Open(path)
	
			defer file.Close()

			
			got, err := DestinationHosts(file)
			if (err != nil) != tt.wantErr {
				t.Errorf("DestinationHosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DestinationHosts() = %v, want %v", got, tt.want)
			}
			os.Remove(path)
		})
	}
}

func TestAskForConfirmation(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bool
	}{
		{
			name:  "Valid input 'yes'",
			input: "yes\n",
			want:  true,
		},
		{
			name:  "Valid input 'no'",
			input: "no\n",
			want:  false,
		},
		{
			name:    "Invalid input 'invalid'",
			input:   "invalid\n",
			want:  false,
		},
		{
			name:    "Invalid input 'invalid'",
			input:   "3\n",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stdin for testing
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			r, w, _ := os.Pipe()
			os.Stdin = r
			defer w.Close()

			// Write input to stdin
			_, _ = w.WriteString(tt.input)

			got := AskForConfirmation()

			if got != tt.want {
				t.Errorf("AskForConfirmation() = %v, want %v", got, tt.want)
			}
		})
	}
}
