package go_binding

import (
	_ "embed"
	"github.com/go-clarum/go-binding/runtime/agent"
)

//go:embed builds/agent/clarum-agent-1.0.0-snapshot
var agentExecutable []byte
var agentService agent.AgentService

func init() {
	agentService = agent.NewAgentService()
	agentService.Initiate(agentExecutable, "clarum-agent-1.0.0-snapshot")
}

func Shutdown() {
	agentService.Shutdown()
}
