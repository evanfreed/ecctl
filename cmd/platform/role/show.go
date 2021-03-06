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

package cmdrole

import (
	"github.com/elastic/cloud-sdk-go/pkg/api/platformapi/roleapi"
	"github.com/spf13/cobra"

	"github.com/elastic/ecctl/pkg/ecctl"
)

var showCmd = &cobra.Command{
	Use:     "show <role>",
	Short:   "Shows the existing platform roles",
	PreRunE: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := roleapi.Show(roleapi.ShowParams{
			API:    ecctl.Get().API,
			ID:     args[0],
			Region: ecctl.Get().Config.Region,
		})
		if err != nil {
			return err
		}

		return ecctl.Get().Formatter.Format("roles/show", res)
	},
}

func init() {
	Command.AddCommand(showCmd)
}
