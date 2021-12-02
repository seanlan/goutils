package dingtalk

import (
	"reflect"
	"testing"
)

func TestClient_SendMessage(t *testing.T) {
	type fields struct {
		accessToken string
		secret      string
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp DingTalkResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		{args: args{
			msg: "123123123",
		}},
		{args: args{
			msg: "aaaaa",
		}},
	}
	token := "xxx"
	secret := "xxx"
	c := New(token, secret)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := c.SendMessage(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("SendMessage() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
