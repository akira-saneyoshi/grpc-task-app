package entity

import (
	"errors"
	"testing"

	"github.com/akira-saneyoshi/task-app/domain/object/entity"
	"github.com/akira-saneyoshi/task-app/domain/object/value"

	"github.com/stretchr/testify/require"
)

func TestUserEntity_Validate(tt *testing.T) {
	testcases := []struct {
		title string
		arg   *entity.User
		err   error
	}{
		{"正常系: 入力データが正しい場合", &entity.User{ID: value.NewID("id"), Name: value.NewName("testuser"), Email: value.NewEmail("test@example.com"), Password: value.NewPassword("pass")}, nil},
		{"準正常系: IDが空の場合", &entity.User{ID: value.NewID(""), Name: value.NewName("testuser"), Email: value.NewEmail("test@example.com"), Password: value.NewPassword("pass")}, errors.New("id is empty")},
		{"準正常系: Nameが空の場合", &entity.User{ID: value.NewID("id"), Name: "", Email: "test@example.com", Password: "password"}, errors.New("name is empty")},
		{"準正常系: Emailが空の場合", &entity.User{ID: value.NewID("id"), Name: value.NewName("testuser"), Email: value.NewEmail(""), Password: value.NewPassword("pass")}, errors.New("email is empty")},
		{"準正常系: Passwordが空の場合", &entity.User{ID: value.NewID("id"), Name: value.NewName("testuser"), Email: value.NewEmail("test@example.com"), Password: value.NewPassword("")}, errors.New("password is empty")},
	}
	for _, v := range testcases {
		tt.Run(v.title, func(t *testing.T) {
			err := v.arg.Validate()

			if v.err == nil {
				require.NoError(t, err, "エラーが発生しないこと")
			} else {
				require.EqualError(t, err, v.err.Error(), "エラーが一致すること")
			}
		})
	}
}
