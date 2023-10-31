package nginx

import (
	"testing"
        "io/ioutil"
)

func newNGINXPlusMainTmpl(t *testing.T) *TemplateExecutor {

    t.Helper()

    nginxConfTemplatePath := "./../../../tmpl/nginx-plus.tmpl"
    tmpl, err := NewTemplateExecutor(nginxConfTemplatePath)
    if err != nil {
        t.Error(err)
    }

    return tmpl
}

func TestExecuteMainTemplateForNGINXPlus(t *testing.T) {

    t.Parallel()

    executor := newNGINXPlusMainTmpl(t)

    data, err := executor.ExecuteMainConfigTemplate(&mainCfg)
    if err != nil {
        t.Errorf("Failed to execute template: %v", err)
    }

    output := "/tmp/nginx.conf"
    ioutil.WriteFile(output, data, 0644)
    t.Log("output to file ", output)

}

func TestExecuteMainTemplateForNGINXPlusHTTP2On(t *testing.T) {

    t.Parallel()

    executor := newNGINXPlusMainTmpl(t)

    data, err := executor.ExecuteMainConfigTemplate(&mainCfgHTTP2On)
    if err != nil {
        t.Errorf("Failed to execute template: %v", err)
    }

    output := "/tmp/nginx-http2On.conf"
    ioutil.WriteFile(output, data, 0644)
    t.Log("output to file ", output)

}

func TestExecuteMainTemplateForNGINXPlusCustomTLSPassthroughPort(t *testing.T) {

    t.Parallel()

    executor := newNGINXPlusMainTmpl(t)

    data, err := executor.ExecuteMainConfigTemplate(&mainCfgCustomTLSPassthroughPort)
    if err != nil {
        t.Errorf("Failed to execute template: %v", err)
    }

    output := "/tmp/nginx-CustomTLSPassthroughPort.conf"
    ioutil.WriteFile(output, data, 0644)
    t.Log("output to file ", output)

}

func TestExecuteMainTemplateForNGINXPlusWithoutTLSPassthrough(t *testing.T) {

    t.Parallel()

    executor := newNGINXPlusMainTmpl(t)

    data, err := executor.ExecuteMainConfigTemplate(&mainCfgWithoutTLSPassthrough)
    if err != nil {
        t.Errorf("Failed to execute template: %v", err)
    }

    output := "/tmp/nginx-WithoutTLSPassthrough.conf"
    ioutil.WriteFile(output, data, 0644)
    t.Log("output to file ", output)

}

func TestExecuteMainTemplateForNGINXPlusDefaultTLSPassthroughPort(t *testing.T) {

    t.Parallel()

    executor := newNGINXPlusMainTmpl(t)

    data, err := executor.ExecuteMainConfigTemplate(&mainCfgDefaultTLSPassthroughPort)
    if err != nil {
        t.Errorf("Failed to execute template: %v", err)
    }

    output := "/tmp/nginx-DefaultTLSPassthroughPort.conf"
    ioutil.WriteFile(output, data, 0644)
    t.Log("output to file ", output)

}

func TestExecuteMainTemplateForNGINXPlusCustomDefaultHTTPAndHTTPSListenerPorts(t *testing.T) {

    t.Parallel()

    executor := newNGINXPlusMainTmpl(t)

    data, err := executor.ExecuteMainConfigTemplate(&mainCfgCustomDefaultHTTPAndHTTPSListenerPorts)
    if err != nil {
        t.Errorf("Failed to execute template: %v", err)
    }

    output := "/tmp/nginx-CustomDefaultHTTPAndHTTPSListenerPorts.conf"
    ioutil.WriteFile(output, data, 0644)
    t.Log("output to file ", output)

}


func TestExecuteMainTemplateForNGINXPlusCustomDefaultHTTPListenerPort(t *testing.T) {

    t.Parallel()

    executor := newNGINXPlusMainTmpl(t)

    data, err := executor.ExecuteMainConfigTemplate(&mainCfgCustomDefaultHTTPListenerPort)
    if err != nil {
        t.Errorf("Failed to execute template: %v", err)
    }

    output := "/tmp/nginx-CustomDefaultHTTPListenerPort.conf"
    ioutil.WriteFile(output, data, 0644)
    t.Log("output to file ", output)

}

func TestExecuteMainTemplateForNGINXPlusCustomDefaultHTTPSListenerPort(t *testing.T) {

    t.Parallel()

    executor := newNGINXPlusMainTmpl(t)

    data, err := executor.ExecuteMainConfigTemplate(&mainCfgCustomDefaultHTTPSListenerPort)
    if err != nil {
        t.Errorf("Failed to execute template: %v", err)
    }

    output := "/tmp/nginx-CustomDefaultHTTPSListenerPort.conf"
    ioutil.WriteFile(output, data, 0644)
    t.Log("output to file ", output)

}


var (

	mainCfg = MainConfig{
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


        mainCfgHTTP2On = MainConfig{
		DefaultHTTPListenerPort:  80,
		DefaultHTTPSListenerPort: 443,
		HTTP2:                    true,
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

	mainCfgCustomTLSPassthroughPort = MainConfig{
		ServerNamesHashMaxSize:  "512",
		ServerTokens:            "off",
		WorkerProcesses:         "auto",
		WorkerCPUAffinity:       "auto",
		WorkerShutdownTimeout:   "1m",
		WorkerConnections:       "1024",
		WorkerRlimitNofile:      "65536",
		LogFormat:               []string{"$remote_addr", "$remote_user"},
		LogFormatEscaping:       "default",
		StreamSnippets:          []string{"# comment"},
		StreamLogFormat:         []string{"$remote_addr", "$remote_user"},
		StreamLogFormatEscaping: "none",
		ResolverAddresses:       []string{"example.com", "127.0.0.1"},
		ResolverIPV6:            false,
		ResolverValid:           "10s",
		ResolverTimeout:         "15s",
		KeepaliveTimeout:        "65s",
		KeepaliveRequests:       100,
		VariablesHashBucketSize: 256,
		VariablesHashMaxSize:    1024,
		TLSPassthrough:          true,
		TLSPassthroughPort:      8443,
	}

	mainCfgWithoutTLSPassthrough = MainConfig{
		ServerNamesHashMaxSize:  "512",
		ServerTokens:            "off",
		WorkerProcesses:         "auto",
		WorkerCPUAffinity:       "auto",
		WorkerShutdownTimeout:   "1m",
		WorkerConnections:       "1024",
		WorkerRlimitNofile:      "65536",
		LogFormat:               []string{"$remote_addr", "$remote_user"},
		LogFormatEscaping:       "default",
		StreamSnippets:          []string{"# comment"},
		StreamLogFormat:         []string{"$remote_addr", "$remote_user"},
		StreamLogFormatEscaping: "none",
		ResolverAddresses:       []string{"example.com", "127.0.0.1"},
		ResolverIPV6:            false,
		ResolverValid:           "10s",
		ResolverTimeout:         "15s",
		KeepaliveTimeout:        "65s",
		KeepaliveRequests:       100,
		VariablesHashBucketSize: 256,
		VariablesHashMaxSize:    1024,
		TLSPassthrough:          false,
		TLSPassthroughPort:      8443,
	}

	mainCfgDefaultTLSPassthroughPort = MainConfig{
		ServerNamesHashMaxSize:  "512",
		ServerTokens:            "off",
		WorkerProcesses:         "auto",
		WorkerCPUAffinity:       "auto",
		WorkerShutdownTimeout:   "1m",
		WorkerConnections:       "1024",
		WorkerRlimitNofile:      "65536",
		LogFormat:               []string{"$remote_addr", "$remote_user"},
		LogFormatEscaping:       "default",
		StreamSnippets:          []string{"# comment"},
		StreamLogFormat:         []string{"$remote_addr", "$remote_user"},
		StreamLogFormatEscaping: "none",
		ResolverAddresses:       []string{"example.com", "127.0.0.1"},
		ResolverIPV6:            false,
		ResolverValid:           "10s",
		ResolverTimeout:         "15s",
		KeepaliveTimeout:        "65s",
		KeepaliveRequests:       100,
		VariablesHashBucketSize: 256,
		VariablesHashMaxSize:    1024,
		TLSPassthrough:          true,
		TLSPassthroughPort:      443,
	}

	mainCfgCustomDefaultHTTPAndHTTPSListenerPorts = MainConfig{
		DefaultHTTPListenerPort:  8083,
		DefaultHTTPSListenerPort: 8443,
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

	mainCfgCustomDefaultHTTPListenerPort = MainConfig{
		DefaultHTTPListenerPort:  8083,
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

	mainCfgCustomDefaultHTTPSListenerPort = MainConfig{
		DefaultHTTPListenerPort:  80,
		DefaultHTTPSListenerPort: 8443,
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

)
