package segment

import (
	"encoding/json"
	"testing"
	"github.com/hq-cml/FalconEngine/src/utils"
)

func init() {
	utils.GSegmenter = utils.NewSegmenter("/tmp/dic.txt")
}

func TestAddDoc(t *testing.T) {
	rIdx := newEmptyInvert(utils.IDX_TYPE_STRING_SEG, 0, "/tmp/xx", nil, nil)
	rIdx.addDocument(0, "我爱北京天安门")
	rIdx.addDocument(1, "天安门上太阳升")
	r, _ := json.Marshal(rIdx.tempHashTable)
	t.Log(string(r))
}