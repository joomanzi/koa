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
	"fmt"
)

type ParamType string

const (
	Integer64 ParamType = "int64"
	Boolean   ParamType = "bool"
	String    ParamType = "string"
)

type Type struct {
	Type ParamType
}

func NewType(paramType string) (Type, error) {
	typ := Type{}

	switch paramType {
	case "int64":
		typ.Type = Integer64
	case "bool":
		typ.Type = Boolean
	case "string":
		typ.Type = String
	default:
		return Type{}, fmt.Errorf("unsupported arg type: %s", paramType)
	}

	return typ, nil
}
