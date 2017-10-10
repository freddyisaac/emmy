/*
 * Copyright 2017 XLAB d.o.o.
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
 *
 */

package cli

import (
	"github.com/urfave/cli"
	"github.com/xlab-si/emmy/server"
)

var ServerCmd = cli.Command{
	Name:  "server",
	Usage: "A server (verifier) that verifies clients (provers)",
	Subcommands: []cli.Command{
		{
			Name:  "start",
			Usage: "Starts emmy server",
			Flags: serverFlags,
			Action: func(ctx *cli.Context) error {
				server.SetLogLevel(ctx.String("loglevel"))
				err := startEmmyServer(ctx.Int("port"), ctx.String("cert"), ctx.String("key"))
				if err != nil {
					return cli.NewExitError(err, 1)
				}
				return nil
			},
		},
	},
}

// startEmmyServer configures and starts the gRPC server at the desired port
func startEmmyServer(port int, certPath, keyPath string) error {
	srv, err := server.NewProtocolServer(certPath, keyPath)
	if err != nil {
		return err
	}
	srv.EnableTracing()
	return srv.Start(port)
}
