package sshabu

import (
	"reflect"
	"testing"
)

func Test_inheritOptions(t *testing.T) {
	type args struct {
		item     interface{}
		addition interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "No changes when src has nil fields",
			args: args{
				item: &Options{ // Your source object with some fields set to non-nil values
					AddressFamily: "ipv4",
					Port:          22,
				},
				addition: &Options{ // Your destination object with some fields set to nil
					AddressFamily: nil,
					Port:          nil,
				},
			},
		},
		{
			name: "Copy non-nil fields from src to dst",
			args: args{
				item: &Options{ // Your source object with all fields set to nil
					AddressFamily: nil,
					Port:          nil,
				},
				addition: &Options{ // Your destination object with some fields set to non-nil values
					AddressFamily: "ipv4",
					Port:          22,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			old_item := tt.args.item.(*Options)
			inheritOptions(tt.args.item, tt.args.addition)
			t.Log(old_item.AddressFamily)
			t.Log(tt.args.item.(*Options).AddressFamily)

			// Check if AddressFamily was copied correctly
			if old_item.AddressFamily == nil {
				// src.AddressFamily is nil, so dst.AddressFamily should remain nil
				if tt.args.addition.(*Options).AddressFamily != nil {
					t.Errorf("AddressFamily was not copied correctly")
				}
			} else if tt.args.item.(*Options).AddressFamily != old_item.AddressFamily {
				t.Errorf("AddressFamily was not copied correctly")
			}

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
						Options: Options{
							AddressFamily: "ipv4",
						},
						Hosts: []Host{
							{
								Name: "Host_in_group",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Name fields",
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
						Name: "Host1",
						Options: Options{
							AddressFamily: "ipv4",
						},
						Hosts: []Host{
							{
								Name: "Host_in_group",
							},
						},
					},
				},
			},
			wantErr: true,
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
			err := shabu.Boil()
			t.Log(shabu.Groups[0].Hosts[0].Options.AddressFamily)
			if (err != nil) != tt.wantErr {
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
		{
			name: "Test with valid options and groups",
			fields: fields{
				Options: Options{
					AddressFamily: "ipv4",
				},
				Name: "Host123",
			},
			args: args{
				groupOptions: Options{
					User: "lvtsky",
				},
			},
			wantErr: false,
		},
		// Add more test cases if needed
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

			// Add assertions to verify that options were inherited correctly
			if host.Options.User != "lvtsky" {
				t.Errorf("User field was not inherited correctly")
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
		{
			name: "Test with valid parent options",
			fields: fields{
				Options: Options{
					User: "lvtsky",
				},
				Name: "Group123",
				Subgroups: []Group{
					{
						Name: "Subgroup1",
						Options: Options{
							Port: 22,
						},
					},
				},
			},
			args: args{
				parentOptions: Options{
					AddressFamily: "ipv4",
				},
			},
			wantErr: false,
		},
		// Add more test cases if needed
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

			// Add assertions to verify that options were inherited correctly
			if group.Options.User != "lvtsky" {
				t.Errorf("User field was not inherited correctly")
			}

			// Add similar checks for other fields or nested structures
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
		{
			name: "Test with valid options and subgroups",
			fields: fields{
				Options: Options{
					User: "lvtsky",
				},
				Name: "Group123",
				Subgroups: []Group{
					{
						Name: "Subgroup1",
						Options: Options{
							Port: 22,
						},
					},
				},
			},
			args: args{
				parentOptions: Options{
					AddressFamily: "ipv4",
				},
			},
			wantErr: false,
		},
		// Add more test cases if needed
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

			// Add assertions to verify that options were inherited correctly
			if group.Options.User != "lvtsky" {
				t.Errorf("User field was not inherited correctly")
			}

			// Add similar checks for other fields or nested structures
		})
	}
}

func TestShabu_FindNamesInShabu(t *testing.T) {
	type fields struct {
		Options Options
		Hosts   []Host
		Groups  []Group
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Test with valid names in Shabu",
			fields: fields{
				Options: Options{
					User: "lvtsky",
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
					{
						Name: "Group2",
					},
				},
			},
			want: []string{"Host1", "Group1", "Group2"},
		},
		// Add more test cases if needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shabu := Shabu{
				Options: tt.fields.Options,
				Hosts:   tt.fields.Hosts,
				Groups:  tt.fields.Groups,
			}
			if got := shabu.FindNamesInShabu(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shabu.FindNamesInShabu() = %v, want %v", got, tt.want)
			}
		})
	}
}

