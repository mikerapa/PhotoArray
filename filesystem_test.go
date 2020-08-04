package PhotoArray

import "testing"

func TestIsValidFilePath(t *testing.T) {

	tests := []struct {
		name string
		path string
		want bool
	}{
		{"valid path from the test images", "./testimages/soccer1.jpg", true},
		{"valid folder, invalid filename", "./testimages/1.jpg", false},
		{"valid folder", "./testimages2/soccer1.jpg", false},
		{name: "empty path", path: "", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidFilePath(tt.path); got != tt.want {
				t.Errorf("IsValidFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
