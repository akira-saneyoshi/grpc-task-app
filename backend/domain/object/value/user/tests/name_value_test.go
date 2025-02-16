package value

import (
	"errors"
	"testing"

	value "github.com/akira-saneyoshi/task-app/domain/object/value/user"

	"github.com/stretchr/testify/require"
)

func TestName_Validate(tt *testing.T) {
	testcases := []struct {
		title string
		arg   *value.Name
		err   error
	}{
		{"正常系: 入力データが正しい場合", value.NewName("test-user-one"), nil},
		{"準正常系: 入力データが空の場合", value.NewName(""), errors.New("name is empty")},
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
