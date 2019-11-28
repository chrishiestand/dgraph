// +build oss

/*
 * Copyright 2018 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package enc

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// BadgerEncryptionKeyFlag is not exposed for OSS build
func BadgerEncryptionKeyFlag(flag *pflag.FlagSet) {
	return
}

// GetEncryptionKeyString return empty string for OSS build
func GetEncryptionKeyString(c *viper.Viper) string {
	return ""
}
