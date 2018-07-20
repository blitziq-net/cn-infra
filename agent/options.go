//  Copyright (c) 2018 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package agent

import (
	"context"
	"os"
	"reflect"
	"syscall"

	"github.com/ligato/cn-infra/infra"
)

// Options specifies option list for the Agent
type Options struct {
	QuitSignals []os.Signal
	QuitChan    chan struct{}
	Context     context.Context

	plugins     []infra.Plugin
	pluginMap   map[infra.Plugin]struct{}
	pluginNames map[string]struct{}
}

func newOptions(opts ...Option) Options {
	opt := Options{
		QuitSignals: []os.Signal{
			os.Interrupt,
			syscall.SIGTERM,
		},
		pluginMap:   make(map[infra.Plugin]struct{}),
		pluginNames: make(map[string]struct{}),
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Option is a function that operates on an Agent's Option
type Option func(*Options)

// Version returns an Option that sets the version of the Agent to the entered string
func Version(buildVer, buildDate, commitHash string) Option {
	return func(o *Options) {
		BuildVersion = buildVer
		BuildDate = buildDate
		CommitHash = commitHash
	}
}

// Context returns an Option that sets the context for the Agent
func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

// QuitSignals returns an Option that will set signals which stop Agent
func QuitSignals(sigs ...os.Signal) Option {
	return func(o *Options) {
		o.QuitSignals = sigs
	}
}

// QuitOnClose returns an Option that will set channel which stops Agent on close
func QuitOnClose(ch chan struct{}) Option {
	return func(o *Options) {
		o.QuitChan = ch
	}
}

// Plugins creates an Option that adds a list of Plugins to the Agent's Plugin list
func Plugins(plugins ...infra.Plugin) Option {
	return func(o *Options) {
		o.plugins = append(o.plugins, plugins...)
	}
}

// AllPlugins creates an Option that adds all of the nested
// plugins recursively to the Agent's plugin list.
func AllPlugins(plugins ...infra.Plugin) Option {
	return func(o *Options) {
		agentLogger.Debugf("AllPlugins with %d plugins", len(plugins))

		for _, plugin := range plugins {
			typ := reflect.TypeOf(plugin)
			agentLogger.Debugf("searching for all deps in: %v (type: %v)", plugin, typ)

			foundPlugins, err := findPlugins(reflect.ValueOf(plugin), o.pluginMap)
			if err != nil {
				panic(err)
			}

			agentLogger.Debugf("found %d plugins in: %v (type: %v)", len(foundPlugins), plugin, typ)
			for _, plug := range foundPlugins {
				agentLogger.Debugf(" - plugin: %v (%v)", plug, reflect.TypeOf(plug))

				if _, ok := o.pluginNames[plug.String()]; ok {
					agentLogger.Fatalf("plugin with name %q already registered", plug.String())
				}
				o.pluginNames[plug.String()] = struct{}{}
			}
			o.plugins = append(o.plugins, foundPlugins...)

			// TODO: perhaps set plugin name to typ.Strilng() if it's empty
			/*p, ok := plugin.(core.PluginNamed)
			if !ok {
				p = core.NamePlugin(typ.String(), plugin)
			}*/

			if _, ok := o.pluginNames[plugin.String()]; ok {
				agentLogger.Fatalf("plugin with name %q already registered, custom name should be used", plugin.String())
			}
			o.pluginNames[plugin.String()] = struct{}{}
			o.plugins = append(o.plugins, plugin)
		}
	}
}
