/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package elasticsearch

import "errors"

var (
	stkIsEmpty  = errors.New("Stack is empty.")
	emptyString = ""
)

type stack struct {
	value []string
}

func NewStack(n uint) *stack {
	return &stack{value: make([]string, 0, n)}
}
func (s *stack) Len() int {
	return len(s.value)
}

func (s *stack) Push(value string) {
	s.value = append(s.value, value)
}

func (s *stack) Pop() (string, error) {
	if s.Len() > 0 {
		result := s.value[s.Len()-1]
		s.value = s.value[:s.Len()-1]
		return result, nil
	}
	return emptyString, stkIsEmpty
}

func (s *stack) Peek() (string, error) {
	if s.Len() > 0 {
		return s.value[s.Len()-1], nil
	}
	return emptyString, stkIsEmpty
}

func (s *stack) All() []string {
	return s.value
}
