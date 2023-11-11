package sshabu

import (
	// "fmt"
	"testing"
)

func Test_inheritOptions(t *testing.T) {
	type args struct {
		src interface{}
		dst interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Copy non-nil fields from src to dst",
			args: args{
				src: &Options{ // Your source object with some fields set to non-nil values
					AddressFamily: "ipv4",
					Port:          22,
				},
				dst: &Options{ // Your destination object with some fields set to nil
					AddressFamily: nil,
					Port:          nil,
				},
			},
		},
		{
			name: "No changes when src has nil fields",
			args: args{
				src: &Options{ // Your source object with all fields set to nil
					AddressFamily: nil,
					Port:          nil,
				},
				dst: &Options{ // Your destination object with some fields set to non-nil values
					AddressFamily: "ipv4",
					Port:          22,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// old_src := tt.args.src
			// old_dst := tt.args.dst
			t.Log(tt.name)
			// t.Log(old_dst)
			// t.Log(old_dst)
			inheritOptions(tt.args.src, tt.args.dst)
			// Add assertions here to verify that non-nil fields were copied correctly
			// t.Log(old_dst)
			t.Log(tt.args.dst)
			// t.Log(tt.args.dst)
			// if old_src.(*Options).AddressFamily != nil && old_dst.(*Options).AddressFamily == nil && tt.args.src.(*Options).AddressFamily != tt.args.dst.(*Options).AddressFamily {
			t.Errorf("AddressFamily was not copied")
			// }
			// if tt.args.src.(*Options).AddressFamily != nil && tt.args.dst.(*Options).AddressFamily == nil {
			// 	t.Errorf("AddressFamily was not copied")
			// }

			// if tt.args.src.(*Options).Port != nil && tt.args.dst.(*Options).Port == nil {
			// 	t.Errorf("Port was not copied")
			// }
		})
	}
}

func TestShabu_Boil(t *testing.T) {
	type fields struct {
		Options Options
		Hosts   []Host
		Groups  []Group
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test with valid options and groups",
			fields: fields{
				Options: Options{
					AddressFamily: "ipv4",
					Port:          22,
				},
				Hosts: []Host{
					{
						Name: "Host1",
					},
				},
				Groups: []Group{
					{
						Name: "Group1",
					},
				},
			},
			wantErr: false,
		},
		// Other cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shabu := &Shabu{
				Options: tt.fields.Options,
				Hosts:   tt.fields.Hosts,
				Groups:  tt.fields.Groups,
			}
			if err := shabu.Boil(); (err != nil) != tt.wantErr {
				t.Errorf("Shabu.Boil() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHost_inheritOptions(t *testing.T) {
	type fields struct {
		Options Options
		Name    string
	}
	type args struct {
		groupOptions Options
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			host := &Host{
				Options: tt.fields.Options,
				Name:    tt.fields.Name,
			}
			if err := host.inheritOptions(tt.args.groupOptions); (err != nil) != tt.wantErr {
				t.Errorf("Host.inheritOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGroup_inheritOptions(t *testing.T) {
	type fields struct {
		Options   Options
		Hosts     []Host
		Name      string
		Subgroups []Group
	}
	type args struct {
		parentOptions Options
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group := &Group{
				Options:   tt.fields.Options,
				Hosts:     tt.fields.Hosts,
				Name:      tt.fields.Name,
				Subgroups: tt.fields.Subgroups,
			}
			if err := group.inheritOptions(tt.args.parentOptions); (err != nil) != tt.wantErr {
				t.Errorf("Group.inheritOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGroup_solveGroup(t *testing.T) {
	type fields struct {
		Options   Options
		Hosts     []Host
		Name      string
		Subgroups []Group
	}
	type args struct {
		parentOptions Options
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group := &Group{
				Options:   tt.fields.Options,
				Hosts:     tt.fields.Hosts,
				Name:      tt.fields.Name,
				Subgroups: tt.fields.Subgroups,
			}
			if err := group.solveGroup(tt.args.parentOptions); (err != nil) != tt.wantErr {
				t.Errorf("Group.solveGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
