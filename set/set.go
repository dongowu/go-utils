package set

import (
	"fmt"
)

// Set 泛型结构体
type Set[T comparable] struct {
	data map[T]struct{}
}

// NewSet 创建一个新的泛型 Set
func NewSet[T comparable](elements ...T) *Set[T] {
	set := &Set[T]{
		data: make(map[T]struct{}),
	}
	// 将传入的元素添加到集合中
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

// Add 向 Set 中添加元素
func (s *Set[T]) Add(element T) {
	s.data[element] = struct{}{}
}

// Remove 从 Set 中删除元素
func (s *Set[T]) Remove(element T) {
	delete(s.data, element)
}

// Contains 检查 Set 中是否包含某个元素
func (s *Set[T]) Contains(element T) bool {
	_, found := s.data[element]
	return found
}

// Size 返回 Set 中元素的数量
func (s *Set[T]) Size() int {
	return len(s.data)
}

// Clear 清空 Set
func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}

// Display 打印 Set 中的元素
func (s *Set[T]) Display() {
	for key := range s.data {
		fmt.Println(key)
	}
}

// ToSlice 返回 Set 中的元素作为一个切片
func (s *Set[T]) ToSlice() []T {
	var slice []T
	for key := range s.data {
		slice = append(slice, key)
	}
	return slice
}

// Union 计算并集
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range s.data {
		result.Add(element)
	}
	for element := range other.data {
		result.Add(element)
	}
	return result
}

// Intersection 计算交集
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range s.data {
		if other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

// Difference 计算差集
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range s.data {
		if !other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

// SymmetricDifference 计算对称差集
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range s.data {
		if !other.Contains(element) {
			result.Add(element)
		}
	}
	for element := range other.data {
		if !s.Contains(element) {
			result.Add(element)
		}
	}
	return result
}
