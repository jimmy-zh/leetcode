package symmetric_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/src/utils"
)

func TestIsSymmetric(t *testing.T) {
	vals := []interface{}{}
	tree := utils.BinaryTreeConstructor(vals)
	if !isSymmetricRecursion(tree) {
		t.Error("isSymmetricRecursion error")
	}
	vals = []interface{}{1}
	tree = utils.BinaryTreeConstructor(vals)
	if !isSymmetricRecursion(tree) {
		t.Error("isSymmetricRecursion error")
	}
	vals = []interface{}{1, 2, 2, 4, 3, 3, 4}
	tree = utils.BinaryTreeConstructor(vals)
	if !isSymmetricRecursion(tree) {
		t.Error("isSymmetricRecursion error")
	}
	vals = []interface{}{1, 2, 3, 4, 3, 3, 4}
	tree = utils.BinaryTreeConstructor(vals)
	if isSymmetricRecursion(tree) {
		t.Error("isSymmetricRecursion error")
	}
	vals = []interface{}{1, 2, 2, 4, 3, nil, 4}
	tree = utils.BinaryTreeConstructor(vals)
	tree = utils.BinaryTreeConstructor(vals)
	if isSymmetricRecursion(tree) {
		t.Error("isSymmetricRecursion error")
	}

	vals = []interface{}{}
	tree = utils.BinaryTreeConstructor(vals)
	if !isSymmetricIterationDFS(tree) {
		t.Error("isSymmetricIteration error")
	}
	vals = []interface{}{1}
	tree = utils.BinaryTreeConstructor(vals)
	if !isSymmetricIterationDFS(tree) {
		t.Error("isSymmetricIteration error")
	}
	vals = []interface{}{1, 2, 2, 4, 3, 3, 4}
	tree = utils.BinaryTreeConstructor(vals)
	if !isSymmetricIterationDFS(tree) {
		t.Error("isSymmetricIteration error")
	}
	vals = []interface{}{1, 2, 3, 4, 3, 3, 4}
	tree = utils.BinaryTreeConstructor(vals)
	if isSymmetricIterationDFS(tree) {
		t.Error("isSymmetricIteration error")
	}
	vals = []interface{}{1, 2, 2, 4, 3, nil, 4}
	tree = utils.BinaryTreeConstructor(vals)
	if isSymmetricIterationDFS(tree) {
		t.Error("isSymmetricIteration error")
	}

	vals = []interface{}{}
	tree = utils.BinaryTreeConstructor(vals)
	if !isSymmetricIterationBFS(tree) {
		t.Error("isSymmetricIteration error")
	}
	vals = []interface{}{1}
	tree = utils.BinaryTreeConstructor(vals)
	if !isSymmetricIterationBFS(tree) {
		t.Error("isSymmetricIteration error")
	}
	vals = []interface{}{1, 2, 2, 4, 3, 3, 4}
	tree = utils.BinaryTreeConstructor(vals)
	if !isSymmetricIterationBFS(tree) {
		t.Error("isSymmetricIteration error")
	}
	vals = []interface{}{1, 2, 3, 4, 3, 3, 4}
	tree = utils.BinaryTreeConstructor(vals)
	if isSymmetricIterationBFS(tree) {
		t.Error("isSymmetricIteration error")
	}
	vals = []interface{}{1, 2, 2, 4, 3, nil, 4}
	tree = utils.BinaryTreeConstructor(vals)
	if isSymmetricIterationBFS(tree) {
		t.Error("isSymmetricIteration error")
	}
}
