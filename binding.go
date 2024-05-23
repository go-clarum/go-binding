package go_binding

import (
	_ "embed"
	"github.com/go-clarum/go-binding/core/logging"
	clarumHttp "github.com/go-clarum/go-binding/http"
	"github.com/go-clarum/go-binding/runtime/agent"
)

//go:embed builds/agent/clarum-agent-1.0.0-snapshot
var agentExecutable []byte
var agentService agent.AgentService

func init() {
	agentService = agent.NewAgentService()

	// when testing the agent from your IDE, just comment this line
	agentService.Initiate(agentExecutable, "clarum-agent-1.0.0-snapshot")
}

func Http() *clarumHttp.EndpointBuilder {
	return &clarumHttp.EndpointBuilder{}
}

func Shutdown() {
	logging.Info("sending shutdown signal to agent")
	agentService.Shutdown()
}
