package strategy

import "testing"

func TestVIPCenter_ServiceUser(t *testing.T) {
	type fields struct {
		providers map[UserType]ServiceProvider
	}
	type args struct {
		user []*User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: struct{ providers map[UserType]ServiceProvider }{providers: map[UserType]ServiceProvider{
				UserTypeNormal: &normalUserServiceProvider{},
				UserTypeVIP:    &vipUserServiceProvider{},
				UserTypeExtra:  &extraUserServiceProvider{},
			}},
			args: args{user: []*User{
				&User{
					Name: "zhang",
					Age:  12,
					Type: UserTypeNormal,
				},
				&User{
					Name: "li",
					Age:  1,
					Type: UserTypeVIP,
				},
				&User{
					Name: "wang",
					Age:  3,
					Type: UserTypeExtra,
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VIPCenter{
				providers: tt.fields.providers,
			}
			for _, user := range tt.args.user {
				v.ServiceUser(user)
			}
		})
	}
}

func TestVIPCenter_ServiceUser1(t *testing.T) {
	type fields struct {
		providers map[UserType]ServiceProvider
	}
	type args struct {
		user *User
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
			v := &VIPCenter{
				providers: tt.fields.providers,
			}
			if err := v.ServiceUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("ServiceUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}