package go_binding

import (
	_ "embed"
	"github.com/go-clarum/go-binding/agent"
	clarumHttp "github.com/go-clarum/go-binding/http"
)

var agentService agent.AgentService

func init() {
	agentService = agent.NewAgentService()

	// while developing you can comment this line and start the agent separately
	//agentService.Initiate()
}

func Http() *clarumHttp.EndpointBuilder {
	return &clarumHttp.EndpointBuilder{}
}

// Shutdown signals the agent that it should stop running. At the moment this does not seem to be required because the
// spawned process is coupled with the one running the tests -> the agent is stopped when the tests stop.
// In the future though, when we implement agent keep alive, we will need this.
func Shutdown() {
	agentService.Shutdown()
}
