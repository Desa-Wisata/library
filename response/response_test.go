package response

import "testing"

func Test_rs_IsError(t *testing.T) {
	type fields struct {
		Code         int
		Status       bool
		ErrorMessage string
		Data         interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "1. This is not errors",
			fields: fields{200, true, "", ""},
			want:   false,
		},

		{
			name:   "2. This is errors",
			fields: fields{500, false, "", ""},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rs{
				Code:         tt.fields.Code,
				Status:       tt.fields.Status,
				ErrorMessage: tt.fields.ErrorMessage,
				Data:         tt.fields.Data,
			}
			if got := r.IsError(); got != tt.want {
				t.Errorf("rs.IsError() = %v, want %v", got, tt.want)
			}
		})
	}
}
