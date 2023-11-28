package counter

import (
	"Rebrain/module06/internal/app/services/post"
	postMock "Rebrain/module06/internal/app/services/post/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostCount(t *testing.T) {
	req := require.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := postMock.NewMockClient(ctrl)

	mockClient.EXPECT().GetList().Return([]post.Post{post.Post{ID: 1}}, nil).Times(1)

	count, err := PostCount(mockClient)
	req.NoError(err)
	req.Equal(1, count)
}
