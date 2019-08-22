/*
Copyright © 2019 Portworx

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
package cmd

import (
	"os"

	"github.com/portworx/px/pkg/commander"
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	gendocCmd       *cobra.Command
	gendocOutputDir string
)

var _ = commander.RegisterCommandVar(func() {
	gendocCmd = &cobra.Command{
		Use:     "gendoc",
		Aliases: []string{"gendocs"},
		Short:   "Generate doc files in Markdown format",
		// Hide this command. Only used for generating docs by developers
		Hidden: true,
		RunE:   gendocExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	RootAddCommand(gendocCmd)

	gendocCmd.Flags().StringVar(&gendocOutputDir, "output-dir", "pxdocs", "Output directory")
})

func GenDocAddCommand(cmd *cobra.Command) {
	gendocCmd.AddCommand(cmd)
}

func gendocExec(cmd *cobra.Command, args []string) error {
	util.Printf("Creating docs in %s...\n", gendocOutputDir)

	os.MkdirAll(gendocOutputDir, 0755)
	return doc.GenMarkdownTree(rootCmd, gendocOutputDir)
}
