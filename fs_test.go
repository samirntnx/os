package os

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestFs(t *testing.T) {
	ctrl := gomock.NewController(t)
	osMock := NewMockOsInterface(ctrl)
	osMock.EXPECT().Open(gomock.Any()).AnyTimes().Return(nil, nil)
	VarOsMock = osMock
	Open("/tmp/1")
}
