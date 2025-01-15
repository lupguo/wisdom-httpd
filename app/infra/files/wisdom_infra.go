package files

import (
	"encoding/json"
	"os"

	"github.com/lupguo/wisdom-httpd/app/domain/entity"
	"github.com/pkg/errors"
)

// ParseJsonWisdom 从json解析wisdom
func ParseJsonWisdom(filename string) (*entity.WisdomFileData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read file `wisdomHandler.json` got err")
	}

	var ws entity.WisdomFileData
	err = json.Unmarshal(data, &ws)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal wisdomHandler json data got err")
	}
	return &ws, nil
}
