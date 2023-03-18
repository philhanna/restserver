package webserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfiguration(t *testing.T) {
	type args struct {
		configfile string
	}
	tests := []struct {
		name     string
		args     args
		wantHost string
		wantPort int
	}{
		{"basic", args{configfile: "./cmd/config.yaml"}, "localhost", 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := NewConfiguration(tt.args.configfile)
			if err != nil {
				t.Errorf("Unexpected error from load: %s\n", err)
			}
			haveHost := config.HOST
			havePort := config.PORT
			assert.Equal(t, tt.wantHost, haveHost)
			assert.Equal(t, tt.wantPort, havePort)
		})
	}
}
