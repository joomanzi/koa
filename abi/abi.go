/*
 * Copyright 2018 De-labtory
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package abi

import (
	"encoding/json"
	"strings"

	"github.com/DE-labtory/koa/ast"
)

type ABI struct {
	Methods []Method
}

func New(abiJSON string) (ABI, error) {
	reader := strings.NewReader(abiJSON)
	dec := json.NewDecoder(reader)

	var abi ABI
	if err := dec.Decode(&abi); err != nil {
		return ABI{}, err
	}

	return abi, nil
}

// UnmarshalJSON is implementation of json.Decoder's UnmarshalJSON
func (abi *ABI) UnmarshalJSON(data []byte) error {
	var methods []Method

	if err := json.Unmarshal(data, &methods); err != nil {
		return err
	}

	for _, method := range methods {
		abi.Methods = append(abi.Methods, method)
	}

	return nil
}

// TODO: implement me w/ test cases :-)
// Extract ABI from function of ast.
func ExtractAbiFromFunction(f ast.FunctionLiteral) (Method, error) {
	return Method{}, nil
}
