package main

import (
	"github.com/sms-api/mcp-server/config"
	"github.com/sms-api/mcp-server/models"
	tools_messages "github.com/sms-api/mcp-server/tools/messages"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_messages.CreateMessagesallTool(cfg),
		tools_messages.CreateMessagesaddTool(cfg),
		tools_messages.CreateMessagesdeleteTool(cfg),
		tools_messages.CreateMessagesoneTool(cfg),
		tools_messages.CreateMessagesupdateTool(cfg),
	}
}
