package pkg

import (
	"os"
	"testing"
)

func TestParseSpecificFlagString(t *testing.T) {
	type args struct {
		flagName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		osArgs  []string
		wantErr bool
	}{
		{
			name: "parse with single - prefix",
			args: args{
				flagName: "test",
			},
			want:   "hello",
			osArgs: []string{"rootCommand", "subCommand1", "subcommand2", "-flag1", "fawekfpwe", "-flag2", "faowkefawe", "-test", "hello"},
		},
		{
			name: "parse with single - prefix and after some flags",
			args: args{
				flagName: "test",
			},
			want:   "hello",
			osArgs: []string{"rootCommand", "subCommand1", "subcommand2", "-flag2", "faowkefawe", "-test", "hello", "-flag1", "fawekfpwe"},
		},
		{
			name: "parse with double (--) prefix",
			args: args{
				flagName: "test-1",
			},
			want:   "hello",
			osArgs: []string{"rootCommand", "subCommand1", "subcommand2", "-flag1", "fawekfpwe", "-flag2", "faowkefawe", "--test-1", "hello"},
		},
		{
			name: "parse with double (--) prefix and after some flags",
			args: args{
				flagName: "test-1",
			},
			want:   "hello",
			osArgs: []string{"rootCommand", "subCommand1", "subcommand2", "-flag2", "faowkefawe", "--test-1", "hello", "-flag1", "fawekfpwe"},
		},
		{
			name: "parse with double (--) prefix and equal sign",
			args: args{
				flagName: "test-1",
			},
			want:   "hello",
			osArgs: []string{"rootCommand", "subCommand1", "subcommand2", "-flag2", "faowkefawe", "--test-1=hello", "-flag1", "fawekfpwe"},
		},
		{
			name: "parse with double (--) prefix and equal sign included in value",
			args: args{
				flagName: "test-1",
			},
			want:   "hello==",
			osArgs: []string{"rootCommand", "subCommand1", "subcommand2", "-flag2", "faowkefawe", "--test-1=hello==", "-flag1", "fawekfpwe"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.osArgs
			got, err := ParseSpecificFlagString(tt.args.flagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSpecificFlagString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseSpecificFlagString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
