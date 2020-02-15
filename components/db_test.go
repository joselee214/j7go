package components

import (
	"testing"

	"github.com/joselee214/j7f/components/dao"
)

var DbConfig = &dao.DBConfig{
	Name: "",
	MaxIdleConns: 5,
	MaxConnNum: 10,
	Master: &dao.NodeConfig{
		Addr: "127.0.0.1:3307",
		User: "root",
		Password: "123456",
		Weight: 1,
	},
	Slave:[]*dao.NodeConfig{
		{
			Addr: "127.0.0.1:3307",
			User: "root",
			Password: "123456",
			Weight: 1,
		},
	},
}

func TestInitDB(t *testing.T) {
	type args struct {
		cfg *dao.DBConfig
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test_db", args: args{cfg:DbConfig}, wantErr: false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitDB(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("InitDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

