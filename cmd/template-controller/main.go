package main

import (
    "fmt"
    "runtime"

    "github.com/dumpplane/template-controller/internal/configs"
    "github.com/dumpplane/template-controller/internal/configs/nginx"
    "github.com/dumpplane/template-controller/internal/configs/nginx/gateway"
    "github.com/golang/glog"
)

var version string

func main() {

    commitHash, commitTime, dirtyBuild := getBuildInfo()
    fmt.Printf("Template Controller Version=%v Commit=%v Date=%v DirtyState=%v Arch=%v/%v Go=%v\n", version, commitHash, commitTime, dirtyBuild, runtime.GOOS, runtime.GOARCH, runtime.Version())

    parseFlags()

    ngxTemplateExecutor, templateExecutor := createTemplateExecutors()

    nginxPlus := true
    cfgParams := configs.NewDefaultConfigParams(nginxPlus)
 
    fmt.Printf("Template Executor: %v, %v, %s\n", ngxTemplateExecutor, templateExecutor, cfgParams.UpstreamZoneSize)


    // The following is used to test the internal SDK
    //   ngxTemplateExecutor.ExecuteMainConfigTemplate - used to create/update main nginx conf
    //   templateExecutor.ExecuteVirtualServerTemplate
    //   templateExecutor.ExecuteTransportServerTemplate(&transportServerCfg)
    // ngxMain, appConf, tcpConf is testted only, will be delete if rest api is avaiable
    ngxMain, err := ngxTemplateExecutor.ExecuteMainConfigTemplate(&mainCfg)
    if err != nil {
        fmt.Errorf("Failed to execute template: %v", err)
    }
    fmt.Println(string(ngxMain))

    appConf, err := templateExecutor.ExecuteVirtualServerTemplate(&virtualServerCfg)
    if err != nil {
        fmt.Errorf("Failed to execute template: %v", err)
    }
    fmt.Println(string(appConf))

    tcpConf, err := templateExecutor.ExecuteTransportServerTemplate(&transportServerCfg)
    if err != nil {
        fmt.Errorf("Failed to execute template: %v", err)
    }
    fmt.Println(string(tcpConf))
}

func createTemplateExecutors() (*nginx.TemplateExecutor, *gateway.TemplateExecutor) {

    nginxConfTemplatePath := "tmpl/nginx-plus.tmpl"
    nginxVirtualServerTemplatePath := "tmpl/nginx-plus.virtualserver.tmpl"
    nginxTransportServerTemplatePath := "tmpl/nginx-plus.transportserver.tmpl"

    ngxTemplateExecutor, err := nginx.NewTemplateExecutor(nginxConfTemplatePath)
    if err != nil {
        glog.Fatalf("Error creating TemplateExecutor: %v", err)
    }

    templateExecutor, err := gateway.NewTemplateExecutor(nginxVirtualServerTemplatePath, nginxTransportServerTemplatePath)
    if err != nil {
        glog.Fatalf("Error creating TemplateExecutor: %v", err)
    }

    return ngxTemplateExecutor, templateExecutor
}

var (
    mainCfg = nginx.MainConfig{
                DefaultHTTPListenerPort:  80,
                DefaultHTTPSListenerPort: 443,
                ServerNamesHashMaxSize:   "512",
                ServerTokens:             "off",
                WorkerProcesses:          "auto",
                WorkerCPUAffinity:        "auto",
                WorkerShutdownTimeout:    "1m",
                WorkerConnections:        "1024",
                WorkerRlimitNofile:       "65536",
                LogFormat:                []string{"$remote_addr", "$remote_user"},
                LogFormatEscaping:        "default",
                StreamSnippets:           []string{"# comment"},
                StreamLogFormat:          []string{"$remote_addr", "$remote_user"},
                StreamLogFormatEscaping:  "none",
                ResolverAddresses:        []string{"example.com", "127.0.0.1"},
                ResolverIPV6:             false,
                ResolverValid:            "10s",
                ResolverTimeout:          "15s",
                KeepaliveTimeout:         "65s",
                KeepaliveRequests:        100,
                VariablesHashBucketSize:  256,
                VariablesHashMaxSize:     1024,
        }

    virtualServerCfg = gateway.VirtualServerConfig{
		Upstreams: []gateway.Upstream{
			{
				Name: "test-upstream",
				Servers: []gateway.UpstreamServer{
					{
						Address: "10.0.0.20:8001",
					},
				},
				LBMethod:         "random",
				Keepalive:        32,
				MaxFails:         4,
				FailTimeout:      "10s",
				MaxConns:         31,
				SlowStart:        "10s",
				UpstreamZoneSize: "256k",
				NTLM:             true,
			},
		},
		Server: gateway.Server{
			ServerName:    "example.com",
			StatusZone:    "example.com",
			ServerTokens:    "off",
			SetRealIPFrom:   []string{"0.0.0.0/0"},
			RealIPHeader:    "X-Real-IP",
			RealIPRecursive: true,
			Allow:           []string{"127.0.0.1"},
			Deny:            []string{"127.0.0.1"},
			HealthChecks: []gateway.HealthCheck{
				{
					Name:       "coffee",
					URI:        "/",
					Interval:   "5s",
					Jitter:     "0s",
					Fails:      1,
					Passes:     1,
					Port:       50,
					ProxyPass:  "http://test-upstream",
					Mandatory:  true,
					Persistent: true,
				},
			},
			Locations: []gateway.Location{
				{
					Path:     "/",
					Snippets: []string{"# location snippet"},
					Allow:    []string{"127.0.0.1"},
					Deny:     []string{"127.0.0.1"},
					ProxyConnectTimeout:      "30s",
					ProxyReadTimeout:         "31s",
					ProxySendTimeout:         "32s",
					ClientMaxBodySize:        "1m",
					ProxyBuffering:           true,
					ProxyBuffers:             "8 4k",
					ProxyBufferSize:          "4k",
					ProxyMaxTempFileSize:     "1024m",
					ProxyPass:                "http://test-upstream",
					ProxyNextUpstream:        "error timeout",
					ProxyNextUpstreamTimeout: "5s",
					Internal:                 true,
					ProxyPassRequestHeaders:  false,
					ProxyPassHeaders:         []string{"Host"},
					ProxyPassRewrite:         "$request_uri",
					ProxyHideHeaders:         []string{"Header"},
					ProxyIgnoreHeaders:       "Cache",
					Rewrites:                 []string{"$request_uri $request_uri", "$request_uri $request_uri"},
					AddHeaders: []gateway.AddHeader{
						{
							Header: gateway.Header{
								Name:  "Header-Name",
								Value: "Header Value",
							},
							Always: true,
						},
					},
				},
			},
		},
	}

    transportServerCfg = gateway.TransportServerConfig{
                Upstreams: []gateway.StreamUpstream{
                        {
                                Name: "udp-upstream",
                                Servers: []gateway.StreamUpstreamServer{
                                        {
                                                Address: "10.0.0.20:5001",
                                        },
                                },
                        },
                },
                Match: &gateway.Match{
                        Name:                "match_udp-upstream",
                        Send:                `GET / HTTP/1.0\r\nHost: localhost\r\n\r\n`,
                        ExpectRegexModifier: "~*",
                        Expect:              "200 OK",
                },
                Server: gateway.StreamServer{
                        Port:                     1234,
                        UDP:                      true,
                        StatusZone:               "udp-app",
                        ProxyPass:                "udp-upstream",
                        ProxyTimeout:             "10s",
                        ProxyConnectTimeout:      "10s",
                        ProxyNextUpstream:        true,
                        ProxyNextUpstreamTimeout: "10s",
                        ProxyNextUpstreamTries:   5,
                        HealthCheck: &gateway.StreamHealthCheck{
                                Enabled:  false,
                                Timeout:  "5s",
                                Jitter:   "0",
                                Port:     8080,
                                Interval: "5s",
                                Passes:   1,
                                Fails:    1,
                                Match:    "match_udp-upstream",
                        },
                },
        }
)
