package value

import (
	"errors"
	"testing"

	value "github.com/akira-saneyoshi/task-app/domain/object/value/title"

	"github.com/stretchr/testify/require"
)

func TestTitle_Validate(tt *testing.T) {
	testcases := []struct {
		title string
		arg   *value.Title
		err   error
	}{
		{"正常系: 入力データが正しい場合", value.NewTitle("title-test"), nil},
		{"準正常系: 入力データが空の場合", value.NewTitle(""), errors.New("title is empty")},
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
