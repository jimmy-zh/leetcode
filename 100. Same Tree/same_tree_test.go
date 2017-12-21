package same_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/src/utils"
)

func TestIsSameTree(t *testing.T) {
	vals1 := []interface{}{nil}
	t1 := utils.BinaryTreeConstructor(vals1)
	vals2 := []interface{}{nil}
	t2 := utils.BinaryTreeConstructor(vals2)
	if !isSameTree(t1, t2) {
		t.Error("isSameTsree error")
	}

	vals1 = []interface{}{1, 2, 3, nil, 5, 4, nil}
	t1 = utils.BinaryTreeConstructor(vals1)
	vals2 = []interface{}{1, 2, 3, nil, 5, 4, nil}
	t2 = utils.BinaryTreeConstructor(vals2)
	if !isSameTree(t1, t2) {
		t.Error("isSameTsree error")
	}

	vals2 = []interface{}{1, 2, 3, nil, 5, 4, 7}
	t2 = utils.BinaryTreeConstructor(vals2)
	if isSameTree(t1, t2) {
		t.Error("isSameTree error")
	}

	vals2 = []interface{}{1, 2, 3, nil, 5, nil, nil}
	t2 = utils.BinaryTreeConstructor(vals2)
	if isSameTree(t1, t2) {
		t.Error("isSameTree error")
	}

}
