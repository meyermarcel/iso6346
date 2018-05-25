// Copyright © 2018 Marcel Meyer meyermarcel@posteo.de
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"path/filepath"
)

type mappedType struct {
	Code     string
	TypeInfo string
}

const jsonTypesFileName = "types.json"

var loadedTypes map[string]string

func initCfgTypes(appDirPath string) {
	pathToTypes := initFile(filepath.Join(appDirPath, jsonTypesFileName), []byte(typesJSON))
	jsonUnmarshal(readFile(pathToTypes), &loadedTypes)
}

func getType(code string) (mappedType, bool) {
	typeInfo, exists := loadedTypes[code]

	return mappedType{code, typeInfo}, exists
}

func getRegexPartTypes() string {
	var regexString string
	for k := range loadedTypes {
		regexString += k + "|"
	}
	return regexString[:len(regexString)-1]
}

const typesJSON = `{
  "A0": "(unassigned)",
  "B0": "(unassigned)",
  "B1": "(unassigned)",
  "B2": "(unassigned)",
  "B3": "(unassigned)",
  "B4": "(unassigned)",
  "B5": "(unassigned)",
  "B6": "(unassigned)",
  "B7": "(unassigned)",
  "B8": "(unassigned)",
  "B9": "(unassigned)",
  "G1": "(unassigned)",
  "G2": "(unassigned)",
  "G3": "(unassigned)",
  "G4": "(unassigned)",
  "G5": "(unassigned)",
  "G6": "(unassigned)",
  "G7": "(unassigned)",
  "G8": "(unassigned)",
  "G9": "(unassigned)",
  "H0": "(unassigned)",
  "H1": "(unassigned)",
  "H2": "(unassigned)",
  "H3": "(unassigned)",
  "H4": "(unassigned)",
  "H5": "(unassigned)",
  "H6": "(unassigned)",
  "H7": "(unassigned)",
  "H8": "(unassigned)",
  "H9": "(unassigned)",
  "K0": "(unassigned)",
  "K1": "(unassigned)",
  "K2": "(unassigned)",
  "K3": "(unassigned)",
  "K4": "(unassigned)",
  "K5": "(unassigned)",
  "K6": "(unassigned)",
  "K7": "(unassigned)",
  "K8": "(unassigned)",
  "K9": "(unassigned)",
  "N0": "(unassigned)",
  "N1": "(unassigned)",
  "N2": "(unassigned)",
  "N3": "(unassigned)",
  "N4": "(unassigned)",
  "N5": "(unassigned)",
  "N6": "(unassigned)",
  "N7": "(unassigned)",
  "N8": "(unassigned)",
  "N9": "(unassigned)",
  "P0": "(unassigned)",
  "P1": "(unassigned)",
  "P2": "(unassigned)",
  "P3": "(unassigned)",
  "P4": "(unassigned)",
  "P5": "(unassigned)",
  "P6": "(unassigned)",
  "P7": "(unassigned)",
  "P8": "(unassigned)",
  "P9": "(unassigned)",
  "R0": "(unassigned)",
  "R1": "(unassigned)",
  "R2": "(unassigned)",
  "R3": "(unassigned)",
  "R4": "(unassigned)",
  "R5": "(unassigned)",
  "R6": "(unassigned)",
  "R7": "(unassigned)",
  "R8": "(unassigned)",
  "R9": "(unassigned)",
  "S0": "(unassigned)",
  "S1": "(unassigned)",
  "S2": "(unassigned)",
  "S3": "(unassigned)",
  "S4": "(unassigned)",
  "S5": "(unassigned)",
  "S6": "(unassigned)",
  "S7": "(unassigned)",
  "S8": "(unassigned)",
  "S9": "(unassigned)",
  "U0": "(unassigned)",
  "U1": "(unassigned)",
  "U2": "(unassigned)",
  "U3": "(unassigned)",
  "U4": "(unassigned)",
  "U5": "(unassigned)",
  "U6": "(unassigned)",
  "U7": "(unassigned)",
  "U8": "(unassigned)",
  "U9": "(unassigned)",
  "V0": "(unassigned)",
  "V1": "(unassigned)",
  "V2": "(unassigned)",
  "V3": "(unassigned)",
  "V4": "(unassigned)",
  "V5": "(unassigned)",
  "V6": "(unassigned)",
  "V7": "(unassigned)",
  "V8": "(unassigned)",
  "V9": "(unassigned)"
}
`
