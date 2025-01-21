package set

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	// 测试空集合创建
	emptySet := NewSet[int]()
	if len(emptySet.data) != 0 {
		t.Errorf("Expected empty set, got %v", emptySet.data)
	}

	// 测试添加单个元素
	singleElementSet := NewSet[int](1)
	if len(singleElementSet.data) != 1 {
		t.Errorf("Expected set with one element, got %v", singleElementSet.data)
	}

	// 测试添加多个元素
	multipleElementsSet := NewSet[int](1, 2, 3)
	if len(multipleElementsSet.data) != 3 {
		t.Errorf("Expected set with three elements, got %v", multipleElementsSet.data)
	}

	// 测试添加重复元素
	duplicateElementSet := NewSet[int](1, 1, 2)
	if len(duplicateElementSet.data) != 2 {
		t.Errorf("Expected set with two elements, got %v", duplicateElementSet.data)
	}

}

// TestAdd 测试 Add 方法
func TestAdd(t *testing.T) {
	set := NewSet[int]()

	// 添加元素
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// 检查元素是否存在
	if !set.Contains(1) {
		t.Errorf("元素 1 应该存在于集合中")
	}
	if !set.Contains(2) {
		t.Errorf("元素 2 应该存在于集合中")
	}
	if !set.Contains(3) {
		t.Errorf("元素 3 应该存在于集合中")
	}
}

// TestRemove 测试 Remove 方法
func TestRemove(t *testing.T) {
	set := NewSet[int]()

	// 添加元素
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// 删除元素
	set.Remove(2)

	// 检查元素是否不存在
	if set.Contains(2) {
		t.Errorf("元素 2 不应该存在于集合中")
	}
}

// TestContains 测试 Contains 方法
func TestContains(t *testing.T) {
	set := NewSet[int]()

	// 检查空集合中是否不包含元素
	if set.Contains(1) {
		t.Errorf("空集合中不应该包含元素 1")
	}

	// 添加元素
	set.Add(1)

	// 检查元素是否存在
	if !set.Contains(1) {
		t.Errorf("元素 1 应该存在于集合中")
	}
}
func TestSetSize(t *testing.T) {
	// 创建一个 Set 实例
	s := NewSet[int]()

	// 初始时，Set 应该为空，Size 应该为 0
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}

	// 添加一些元素
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// Size 应该为添加的元素数量
	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}
}

func TestSetClear(t *testing.T) {
	// 创建一个 Set 实例并添加一些元素
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// 清空 Set
	s.Clear()

	// 清空后，Size 应该为 0
	if s.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", s.Size())
	}
}

// TestToSlice 测试 ToSlice 方法
func TestToSlice(t *testing.T) {
	// 创建一个 Set 实例
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	// 调用 ToSlice 方法并验证结果
	expectedSlice := []int{1, 2, 3}
	slice := s.ToSlice()
	if !sliceEqual(slice, expectedSlice) {
		t.Errorf("Expected slice: %v, got: %v", expectedSlice, slice)
	}
}

// captureOutput 捕获函数的输出
func captureOutput() string {
	// 这里可以使用一些方法来捕获输出
	// 为了简单起见，这里直接返回一个空字符串
	return ""
}

// sliceEqual 比较两个切片是否相等
func sliceEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
