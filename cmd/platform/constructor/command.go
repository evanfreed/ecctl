// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cmdconstructor

import (
	"github.com/spf13/cobra"

	cmdutil "github.com/elastic/ecctl/cmd/util"
)

const (
	constructorShowMessage        = `Returns information about the constructor with given ID`
	constructorMaintenanceMessage = `Sets/un-sets a constructor's maintenance mode`
)

// Command represents the constructor command
var Command = &cobra.Command{
	Use:     "constructor",
	Short:   cmdutil.AdminReqDescription("Manages constructors"),
	PreRunE: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
