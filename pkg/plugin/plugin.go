package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/katallaxie/csync/pkg/proto"
	"github.com/katallaxie/csync/pkg/provider"
	"github.com/katallaxie/csync/pkg/spec"

	"github.com/hashicorp/go-hclog"
	p "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

var enablePluginAutoMTLS = os.Getenv("RUN_DISABLE_PLUGIN_TLS") == ""

// Meta are the meta information provided for the plugin.
// These are the arguments and the path to the plugin.
type Meta struct {
	// Path ...
	Path string
	// Arguments ...
	Arguments []string
}

// ExecutableFile ...
func (m *Meta) ExecutableFile() (string, error) {
	// TODO: make this check for the executable file
	return m.Path, nil
}

func (m *Meta) Factory(ctx context.Context) Factory {
	return pluginFactory(ctx, m)
}

// GRPCProviderPlugin ...
type GRPCProviderPlugin struct {
	p.Plugin
	GRPCPlugin func() proto.PluginServer
}

// GRPCClient ...
func (p *GRPCProviderPlugin) GRPCClient(_ context.Context, _ *p.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCPlugin{
		client: proto.NewPluginClient(c),
	}, nil
}

func (p *GRPCProviderPlugin) GRPCServer(_ *p.GRPCBroker, s *grpc.Server) error {
	proto.RegisterPluginServer(s, p.GRPCPlugin())
	return nil
}

// GRPCPlugin contains the configuration and the client connection
// for the provider Plugin.
type GRPCPlugin struct {
	PluginClient *p.Client
	Meta         *Meta

	client proto.PluginClient
}

// Close is closing the gRPC connection if a plugin is configured.
func (p *GRPCPlugin) Close() error {
	if p.PluginClient != nil {
		return nil
	}

	p.PluginClient.Kill()

	return nil
}

// Backup is sending a request to the plugin to backup the app.
func (p *GRPCPlugin) Backup(ctx context.Context, app *spec.App, opts *provider.Opts) error {
	r := new(proto.Backup_Request)
	r.Force = opts.Force
	r.Dry = opts.Dry
	r.Root = opts.Root

	r.Args = p.Meta.Arguments

	r.App = app.ToProto()

	_, err := p.client.Backup(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

// Restore is sending a request to the plugin to restore the app.
func (p *GRPCPlugin) Restore(ctx context.Context, app *spec.App, opts *provider.Opts) error {
	r := new(proto.Restore_Request)
	r.Force = opts.Force
	r.Dry = opts.Dry
	r.Root = opts.Root

	r.Args = p.Meta.Arguments

	r.App = app.ToProto()

	_, err := p.client.Restore(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

// Link is sending a request to link an app with the plugin provided.
func (p *GRPCPlugin) Link(ctx context.Context, app *spec.App, opts *provider.Opts) error {
	r := new(proto.Link_Request)
	r.Force = opts.Force
	r.Dry = opts.Dry
	r.Root = opts.Root

	r.Args = p.Meta.Arguments

	r.App = app.ToProto()

	_, err := p.client.Link(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

// Unlink is unlinking an app with the plugin provided.
// Restore is sending a request to the plugin to restore the app.
func (p *GRPCPlugin) Unlink(ctx context.Context, app *spec.App, opts *provider.Opts) error {
	r := new(proto.Unlink_Request)
	r.Force = opts.Force
	r.Dry = opts.Dry
	r.Root = opts.Root

	r.Args = p.Meta.Arguments

	r.App = app.ToProto()

	_, err := p.client.Unlink(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

// Factory is creatig a new instance of the plugin.
type Factory func() (Plugin, error)

var _ provider.Provider = (*GRPCPlugin)(nil)

// Plugin is defining the interface for a plugin.
// Which essentially implements the provider interface.
type Plugin interface {
	// Close ...
	Close() error

	provider.Provider
}

func pluginFactory(ctx context.Context, meta *Meta) Factory {
	return func() (Plugin, error) {
		f, err := meta.ExecutableFile()
		if err != nil {
			return nil, err
		}

		l := hclog.New(&hclog.LoggerOptions{
			Name:  meta.Path,
			Level: hclog.LevelFromString("DEBUG"),
		})

		cfg := &p.ClientConfig{
			Logger:           l,
			VersionedPlugins: VersionedPlugins,
			HandshakeConfig:  Handshake,
			AutoMTLS:         enablePluginAutoMTLS,
			Managed:          true,
			AllowedProtocols: []p.Protocol{p.ProtocolGRPC},
			Cmd:              exec.CommandContext(ctx, f, meta.Arguments...),
			SyncStderr:       l.StandardWriter(&hclog.StandardLoggerOptions{}),
			SyncStdout:       l.StandardWriter(&hclog.StandardLoggerOptions{}),
		}
		client := p.NewClient(cfg)

		rpc, err := client.Client()
		if err != nil {
			return nil, err
		}

		raw, err := rpc.Dispense(PluginName)
		if err != nil {
			return nil, err
		}

		p, ok := raw.(*GRPCPlugin)
		if !ok {
			return nil, fmt.Errorf("invalid plugin type %T", raw)
		}

		p.PluginClient = client
		p.Meta = meta

		return p, nil
	}
}
