// Copyright 2019 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestReverseSimple(t *testing.T) {
	err := reverseFromConfig("../example/simple.yml")
	assert.NoError(t, err)

	outputBytes, err := ioutil.ReadFile("../models/simple/models.go")
	assert.NoError(t, err)

	// expected output
	expectedOutputBytes, err := ioutil.ReadFile("../models/simple/simple-models.go")
	assert.NoError(t, err)

	assert.EqualValues(t, string(expectedOutputBytes), string(outputBytes))

	t.Cleanup(func() {
		if !t.Failed() {
			os.Remove("../models/simple/models.go")
		}
	})
}

func TestReverseMultiple(t *testing.T) {
	err := reverseFromConfig("../example/multiple.yml")
	assert.NoError(t, err)

	oneOutputBytes, err := ioutil.ReadFile("../models/multiple/one/models.go")
	assert.NoError(t, err)

	// expected output
	expectedOneOutputBytes, err := ioutil.ReadFile("../models/multiple/one/expected-models.go")
	assert.NoError(t, err)

	assert.EqualValues(t, string(expectedOneOutputBytes), string(oneOutputBytes))

	twoOutputBytes, err := ioutil.ReadFile("../models/multiple/two/models.go")
	assert.NoError(t, err)

	// expected output
	expectedTwoOutputBytes, err := ioutil.ReadFile("../models/multiple/two/expected-models.go")
	assert.NoError(t, err)

	assert.EqualValues(t, string(expectedTwoOutputBytes), string(twoOutputBytes))

	t.Cleanup(func() {
		if !t.Failed() {
			os.Remove("../models/multiple/one/models.go")
			os.Remove("../models/multiple/two/models.go")
		}
	})
}

func TestReverseMultipleTarget(t *testing.T) {
	err := reverseFromConfig("../example/multiple-target.yml")
	assert.NoError(t, err)

	oneOutputBytes, err := ioutil.ReadFile("../models/multiple-target/one/models.go")
	assert.NoError(t, err)

	// expected output
	expectedOneOutputBytes, err := ioutil.ReadFile("../models/multiple-target/one/expected-models.go")
	assert.NoError(t, err)

	assert.EqualValues(t, string(expectedOneOutputBytes), string(oneOutputBytes))

	twoOutputBytes, err := ioutil.ReadFile("../models/multiple-target/two/models.go")
	assert.NoError(t, err)

	// expected output
	expectedTwoOutputBytes, err := ioutil.ReadFile("../models/multiple-target/two/expected-models.go")
	assert.NoError(t, err)

	assert.EqualValues(t, string(expectedTwoOutputBytes), string(twoOutputBytes))

	t.Cleanup(func() {
		if !t.Failed() {
			os.Remove("../models/multiple-target/one/models.go")
			os.Remove("../models/multiple-target/two/models.go")
		}
	})
}

func TestReverseIncludeExclude(t *testing.T) {
	err := reverseFromConfig("../example/include-exclude.yml")
	assert.NoError(t, err)

	incOnlyOutputBytes, err := ioutil.ReadFile("../models/include-exclude/include-only/models.go")
	assert.NoError(t, err)

	// expected output
	expectedIncOnlyOutputBytes, err := ioutil.ReadFile("../models/include-exclude/include-only/expected-models.go")
	assert.NoError(t, err)

	assert.EqualValues(t, string(expectedIncOnlyOutputBytes), string(incOnlyOutputBytes))

	excOnlyOutputBytes, err := ioutil.ReadFile("../models/include-exclude/exclude-only/models.go")
	assert.NoError(t, err)

	// expected output
	expectedExcOnlyOutputBytes, err := ioutil.ReadFile("../models/include-exclude/exclude-only/expected-models.go")
	assert.NoError(t, err)

	assert.EqualValues(t, string(expectedExcOnlyOutputBytes), string(excOnlyOutputBytes))

	bothOutputBytes, err := ioutil.ReadFile("../models/include-exclude/both/models.go")
	assert.NoError(t, err)

	// expected output
	expectedBothOutputBytes, err := ioutil.ReadFile("../models/include-exclude/both/expected-models.go")
	assert.NoError(t, err)

	assert.EqualValues(t, string(expectedBothOutputBytes), string(bothOutputBytes))

	t.Cleanup(func() {
		os.Remove("../models/include-exclude/include-only/models.go")
		os.Remove("../models/include-exclude/exclude-only/models.go")
		os.Remove("../models/include-exclude/both/models.go")
	})
}

func TestReverseExplicitTableName(t *testing.T) {
	err := reverseFromConfig("../example/explicit-table-name.yml")
	assert.NoError(t, err)

	outputBytes, err := ioutil.ReadFile("../models/explicit-table-name/models.go")
	assert.NoError(t, err)

	// expected output
	expectedOutputBytes, err := ioutil.ReadFile("../models/explicit-table-name/expected-models.go")
	assert.NoError(t, err)

	assert.EqualValues(t, string(expectedOutputBytes), string(outputBytes))

	t.Cleanup(func() {
		if !t.Failed() {
			os.Remove("../models/explicit-table-name/models.go")
		}
	})
}

func TestReverseExplicitColumnName(t *testing.T) {
	err := reverseFromConfig("../example/explicit-column-name.yml")
	assert.NoError(t, err)

	outputBytes, err := ioutil.ReadFile("../models/explicit-column-name/models.go")
	assert.NoError(t, err)

	// expected output
	expectedOutputBytes, err := ioutil.ReadFile("../models/explicit-column-name/expected-models.go")
	assert.NoError(t, err)

	assert.EqualValues(t, string(expectedOutputBytes), string(outputBytes))

	t.Cleanup(func() {
		if !t.Failed() {
			os.Remove("../models/explicit-column-name/models.go")
		}
	})
}
