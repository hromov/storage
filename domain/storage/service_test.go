package storage_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hromov/storage/domain/storage"
	"github.com/hromov/storage/mocks"
	"github.com/stretchr/testify/require"
)

func TestService_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fs := mocks.NewMockFileService(ctrl)
	ns := mocks.NewMockNotifierService(ctrl)

	s := storage.NewService(fs, ns)

	file := bytes.NewReader([]byte("test"))

	t.Run("success", func(t *testing.T) {
		fs.EXPECT().Save(storage.CategoryVideo, gomock.Any(), file).Return(nil)
		ns.EXPECT().Notify(gomock.Any()).Return(nil)

		err := s.Save(storage.CategoryVideo, file)
		require.NoError(t, err)
	})

	t.Run("fileService error", func(t *testing.T) {
		expectedError := errors.New("error")
		fs.EXPECT().Save(storage.CategoryVideo, gomock.Any(), file).Return(expectedError)

		err := s.Save(storage.CategoryVideo, file)
		require.ErrorIs(t, err, expectedError)
	})

	t.Run("notifService error", func(t *testing.T) {
		expectedError := errors.New("error")
		fs.EXPECT().Save(storage.CategoryVideo, gomock.Any(), file).Return(nil)
		ns.EXPECT().Notify(gomock.Any()).Return(expectedError)

		err := s.Save(storage.CategoryVideo, file)
		require.ErrorIs(t, err, expectedError)
	})

	t.Run("notified with not nil event", func(t *testing.T) {
		fs.EXPECT().Save(storage.CategoryVideo, gomock.Any(), file).Return(nil)
		ns.EXPECT().Notify(gomock.Not(nil)).Return(nil)

		err := s.Save(storage.CategoryVideo, file)
		require.NoError(t, err)
	})
}
