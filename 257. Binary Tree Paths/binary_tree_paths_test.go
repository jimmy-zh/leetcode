package binary_tree_paths

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/src/utils"
)

func TestBinaryTreePaths(t *testing.T) {

	/*
		s := make([]int, 0)
		fmt.Printf("%p\n", &s)
		fmt.Printf("%p\n", s)
		my(&s)
		fmt.Printf("%p\n", &s)
		fmt.Printf("%p\n", s)
		fmt.Println(s)
		return
	*/
	vals := []interface{}{1, 2, 3, 4, 5, 6, nil, nil, nil, 7, nil, nil, nil}
	tree := utils.BinaryTreeConstructor(vals)
	result := []string{"1->2->4", "1->2->5->7", "1->3->6"}
	res := binaryTreePaths(tree)

	if len(result) != len(res) {
		t.Error("binaryTreePaths error")
	}
	for i := range result {
		if result[i] != res[i] {
			t.Error("binaryTreePaths error")
		}
	}

}

func my(arr *[]int) {
	*arr = append(*arr, 1)
}
