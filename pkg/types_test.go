package sshabu

import "testing"

func Test_inheritOptions(t *testing.T) {
	type args struct {
		src interface{}
		dst interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inheritOptions(tt.args.src, tt.args.dst)
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
		// TODO: Add test cases.
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
