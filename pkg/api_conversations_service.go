/*
 * Microsoft Bot Connector API - v3.0
 *
 * The Bot Connector REST API allows your bot to send and receive messages to channels configured in the  [Bot Framework Developer Portal](https://dev.botframework.com). The Connector service uses industry-standard REST  and JSON over HTTPS.    Client libraries for this REST API are available. See below for a list.    Many bots will use both the Bot Connector REST API and the associated [Bot State REST API](/en-us/restapi/state). The  Bot State REST API allows a bot to store and retrieve state associated with users and conversations.    Authentication for both the Bot Connector and Bot State REST APIs is accomplished with JWT Bearer tokens, and is  described in detail in the [Connector Authentication](/en-us/restapi/authentication) document.    # Client Libraries for the Bot Connector REST API    * [Bot Builder for C#](/en-us/csharp/builder/sdkreference/)  * [Bot Builder for Node.js](/en-us/node/builder/overview/)  * Generate your own from the [Connector API Swagger file](https://raw.githubusercontent.com/Microsoft/BotBuilder/master/CSharp/Library/Microsoft.Bot.Connector.Shared/Swagger/ConnectorAPI.json)    © 2016 Microsoft
 *
 * API version: v3
 * Contact: botframework@microsoft.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package botbuilder

import (
	"errors"
)

// ConversationsApiService is a service that implents the logic for the ConversationsApiServicer
// This service should implement the business logic for every endpoint for the ConversationsApi API.
// Include any external packages or services that will be required by this service.
type ConversationsApiService struct {
}

// NewConversationsApiService creates a default api service
func NewConversationsApiService() ConversationsApiServicer {
	return &ConversationsApiService{}
}

// ConversationsCreateConversation - CreateConversation
func (s *ConversationsApiService) ConversationsCreateConversation(parameters ConversationParameters) (interface{}, error) {
	// TODO - update ConversationsCreateConversation with the required logic for this service method.
	// Add api_conversations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ConversationsCreateConversation' not implemented")
}

// ConversationsDeleteActivity - DeleteActivity
func (s *ConversationsApiService) ConversationsDeleteActivity(conversationId string, activityId string) (interface{}, error) {
	// TODO - update ConversationsDeleteActivity with the required logic for this service method.
	// Add api_conversations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ConversationsDeleteActivity' not implemented")
}

// ConversationsDeleteConversationMember - DeleteConversationMember
func (s *ConversationsApiService) ConversationsDeleteConversationMember(conversationId string, memberId string) (interface{}, error) {
	// TODO - update ConversationsDeleteConversationMember with the required logic for this service method.
	// Add api_conversations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ConversationsDeleteConversationMember' not implemented")
}

// ConversationsGetActivityMembers - GetActivityMembers
func (s *ConversationsApiService) ConversationsGetActivityMembers(conversationId string, activityId string) (interface{}, error) {
	// TODO - update ConversationsGetActivityMembers with the required logic for this service method.
	// Add api_conversations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ConversationsGetActivityMembers' not implemented")
}

// ConversationsGetConversationMembers - GetConversationMembers
func (s *ConversationsApiService) ConversationsGetConversationMembers(conversationId string) (interface{}, error) {
	// TODO - update ConversationsGetConversationMembers with the required logic for this service method.
	// Add api_conversations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ConversationsGetConversationMembers' not implemented")
}

// ConversationsGetConversationPagedMembers - GetConversationPagedMembers
func (s *ConversationsApiService) ConversationsGetConversationPagedMembers(conversationId string, pageSize int32, continuationToken string) (interface{}, error) {
	// TODO - update ConversationsGetConversationPagedMembers with the required logic for this service method.
	// Add api_conversations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ConversationsGetConversationPagedMembers' not implemented")
}

// ConversationsGetConversations - GetConversations
func (s *ConversationsApiService) ConversationsGetConversations(continuationToken string) (interface{}, error) {
	// TODO - update ConversationsGetConversations with the required logic for this service method.
	// Add api_conversations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ConversationsGetConversations' not implemented")
}

// ConversationsReplyToActivity - ReplyToActivity
func (s *ConversationsApiService) ConversationsReplyToActivity(conversationId string, activityId string, activity Activity) (interface{}, error) {
	// TODO - update ConversationsReplyToActivity with the required logic for this service method.
	// Add api_conversations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ConversationsReplyToActivity' not implemented")
}

// ConversationsSendConversationHistory - SendConversationHistory
func (s *ConversationsApiService) ConversationsSendConversationHistory(conversationId string, history Transcript) (interface{}, error) {
	// TODO - update ConversationsSendConversationHistory with the required logic for this service method.
	// Add api_conversations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ConversationsSendConversationHistory' not implemented")
}

// ConversationsSendToConversation - SendToConversation
func (s *ConversationsApiService) ConversationsSendToConversation(conversationId string, activity Activity) (interface{}, error) {
	// TODO - update ConversationsSendToConversation with the required logic for this service method.
	// Add api_conversations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ConversationsSendToConversation' not implemented")
}

// ConversationsUpdateActivity - UpdateActivity
func (s *ConversationsApiService) ConversationsUpdateActivity(conversationId string, activityId string, activity Activity) (interface{}, error) {
	// TODO - update ConversationsUpdateActivity with the required logic for this service method.
	// Add api_conversations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ConversationsUpdateActivity' not implemented")
}

// ConversationsUploadAttachment - UploadAttachment
func (s *ConversationsApiService) ConversationsUploadAttachment(conversationId string, attachmentUpload AttachmentData) (interface{}, error) {
	// TODO - update ConversationsUploadAttachment with the required logic for this service method.
	// Add api_conversations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ConversationsUploadAttachment' not implemented")
}
