// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
)

func (classes *ClassCollection) generateProto(packageName string) string {
	code := CodeBuilder{}
	code.AddLine(LICENSE)
	code.AddLine("// THIS FILE IS AUTOMATICALLY GENERATED.")
	code.AddLine()

	code.AddLine("syntax = \"proto3\";")
	code.AddLine()
	code.AddLine("package " + packageName + ";")
	code.AddLine()
	code.AddLine("import \"google/protobuf/any.proto\";")
	code.AddLine()

	classNames := classes.sortedClassNames()
	for _, className := range classNames {
		code.AddLine("message %s {", className)
		classModel := classes.ClassModels[className]
		propertyNames := classModel.sortedPropertyNames()
		var fieldNumber = 0
		for _, propertyName := range propertyNames {
			propertyModel := classModel.Properties[propertyName]
			fieldNumber += 1
			propertyType := propertyModel.Type
			if propertyType == "int" {
				propertyType = "int64"
			}
			var displayName = propertyName
			if displayName == "$ref" {
				displayName = "_ref"
			}
			if displayName == "$schema" {
				displayName = "_schema"
			}
			displayName = camelCaseToSnakeCase(displayName)

			var line = fmt.Sprintf("%s %s = %d;", propertyType, displayName, fieldNumber)
			if propertyModel.Repeated {
				line = "repeated " + line
			}
			code.AddLine("  " + line)
		}
		code.AddLine("}")
		code.AddLine()
	}
	return code.Text()
}
