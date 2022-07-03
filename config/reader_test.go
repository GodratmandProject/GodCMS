package config

import "testing"

const testDir = "test_data/"

func TestReader(t *testing.T) {
	type args struct {
		cfg  *Config
		path string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool // check config values later (validate)
	}{
		{
			name: "yml file",
			args: args{
				cfg:  &Config{},
				path: testDir + "cfg.yml",
			},
			wantErr: false,
		},
		{
			//other tests
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if err := Parse(test.args.cfg, test.args.path); (err != nil) != test.wantErr {
				tt.Errorf("Parse(): wantErr = %v, err = %v", test.wantErr, err)
			}
		})
	}
}
