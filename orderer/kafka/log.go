/*
Copyright IBM Corp. 2016 All Rights Reserved.

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

package kafka

import (
	"os"
	"strings"

	logging "github.com/op/go-logging"
)

var logger *logging.Logger

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	logging.SetBackend(backend)
	formatter := logging.MustStringFormatter("[%{time:15:04:05}] %{shortfile:18s}: %{color}[%{level:-5s}]%{color:reset} %{message}")
	logging.SetFormatter(formatter)
	logger = logging.MustGetLogger("orderer/kafka")
	logging.SetLevel(logging.INFO, "") // Silence debug-level outputs when testing
}

// SetLogLevel sets the package logging level
func SetLogLevel(level string) {
	logLevel, _ := logging.LogLevel(strings.ToUpper(level)) // TODO Validate input
	logging.SetLevel(logLevel, logger.Module)
}
