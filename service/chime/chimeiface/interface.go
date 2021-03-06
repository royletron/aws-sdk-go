// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package chimeiface provides an interface to enable mocking the Amazon Chime service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package chimeiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/chime"
)

// ChimeAPI provides an interface to enable mocking the
// chime.Chime service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Chime.
//    func myFunc(svc chimeiface.ChimeAPI) bool {
//        // Make svc.AssociatePhoneNumberWithUser request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := chime.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockChimeClient struct {
//        chimeiface.ChimeAPI
//    }
//    func (m *mockChimeClient) AssociatePhoneNumberWithUser(input *chime.AssociatePhoneNumberWithUserInput) (*chime.AssociatePhoneNumberWithUserOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockChimeClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ChimeAPI interface {
	AssociatePhoneNumberWithUser(*chime.AssociatePhoneNumberWithUserInput) (*chime.AssociatePhoneNumberWithUserOutput, error)
	AssociatePhoneNumberWithUserWithContext(aws.Context, *chime.AssociatePhoneNumberWithUserInput, ...request.Option) (*chime.AssociatePhoneNumberWithUserOutput, error)
	AssociatePhoneNumberWithUserRequest(*chime.AssociatePhoneNumberWithUserInput) (*request.Request, *chime.AssociatePhoneNumberWithUserOutput)

	AssociatePhoneNumbersWithVoiceConnector(*chime.AssociatePhoneNumbersWithVoiceConnectorInput) (*chime.AssociatePhoneNumbersWithVoiceConnectorOutput, error)
	AssociatePhoneNumbersWithVoiceConnectorWithContext(aws.Context, *chime.AssociatePhoneNumbersWithVoiceConnectorInput, ...request.Option) (*chime.AssociatePhoneNumbersWithVoiceConnectorOutput, error)
	AssociatePhoneNumbersWithVoiceConnectorRequest(*chime.AssociatePhoneNumbersWithVoiceConnectorInput) (*request.Request, *chime.AssociatePhoneNumbersWithVoiceConnectorOutput)

	BatchDeletePhoneNumber(*chime.BatchDeletePhoneNumberInput) (*chime.BatchDeletePhoneNumberOutput, error)
	BatchDeletePhoneNumberWithContext(aws.Context, *chime.BatchDeletePhoneNumberInput, ...request.Option) (*chime.BatchDeletePhoneNumberOutput, error)
	BatchDeletePhoneNumberRequest(*chime.BatchDeletePhoneNumberInput) (*request.Request, *chime.BatchDeletePhoneNumberOutput)

	BatchSuspendUser(*chime.BatchSuspendUserInput) (*chime.BatchSuspendUserOutput, error)
	BatchSuspendUserWithContext(aws.Context, *chime.BatchSuspendUserInput, ...request.Option) (*chime.BatchSuspendUserOutput, error)
	BatchSuspendUserRequest(*chime.BatchSuspendUserInput) (*request.Request, *chime.BatchSuspendUserOutput)

	BatchUnsuspendUser(*chime.BatchUnsuspendUserInput) (*chime.BatchUnsuspendUserOutput, error)
	BatchUnsuspendUserWithContext(aws.Context, *chime.BatchUnsuspendUserInput, ...request.Option) (*chime.BatchUnsuspendUserOutput, error)
	BatchUnsuspendUserRequest(*chime.BatchUnsuspendUserInput) (*request.Request, *chime.BatchUnsuspendUserOutput)

	BatchUpdatePhoneNumber(*chime.BatchUpdatePhoneNumberInput) (*chime.BatchUpdatePhoneNumberOutput, error)
	BatchUpdatePhoneNumberWithContext(aws.Context, *chime.BatchUpdatePhoneNumberInput, ...request.Option) (*chime.BatchUpdatePhoneNumberOutput, error)
	BatchUpdatePhoneNumberRequest(*chime.BatchUpdatePhoneNumberInput) (*request.Request, *chime.BatchUpdatePhoneNumberOutput)

	BatchUpdateUser(*chime.BatchUpdateUserInput) (*chime.BatchUpdateUserOutput, error)
	BatchUpdateUserWithContext(aws.Context, *chime.BatchUpdateUserInput, ...request.Option) (*chime.BatchUpdateUserOutput, error)
	BatchUpdateUserRequest(*chime.BatchUpdateUserInput) (*request.Request, *chime.BatchUpdateUserOutput)

	CreateAccount(*chime.CreateAccountInput) (*chime.CreateAccountOutput, error)
	CreateAccountWithContext(aws.Context, *chime.CreateAccountInput, ...request.Option) (*chime.CreateAccountOutput, error)
	CreateAccountRequest(*chime.CreateAccountInput) (*request.Request, *chime.CreateAccountOutput)

	CreatePhoneNumberOrder(*chime.CreatePhoneNumberOrderInput) (*chime.CreatePhoneNumberOrderOutput, error)
	CreatePhoneNumberOrderWithContext(aws.Context, *chime.CreatePhoneNumberOrderInput, ...request.Option) (*chime.CreatePhoneNumberOrderOutput, error)
	CreatePhoneNumberOrderRequest(*chime.CreatePhoneNumberOrderInput) (*request.Request, *chime.CreatePhoneNumberOrderOutput)

	CreateVoiceConnector(*chime.CreateVoiceConnectorInput) (*chime.CreateVoiceConnectorOutput, error)
	CreateVoiceConnectorWithContext(aws.Context, *chime.CreateVoiceConnectorInput, ...request.Option) (*chime.CreateVoiceConnectorOutput, error)
	CreateVoiceConnectorRequest(*chime.CreateVoiceConnectorInput) (*request.Request, *chime.CreateVoiceConnectorOutput)

	DeleteAccount(*chime.DeleteAccountInput) (*chime.DeleteAccountOutput, error)
	DeleteAccountWithContext(aws.Context, *chime.DeleteAccountInput, ...request.Option) (*chime.DeleteAccountOutput, error)
	DeleteAccountRequest(*chime.DeleteAccountInput) (*request.Request, *chime.DeleteAccountOutput)

	DeletePhoneNumber(*chime.DeletePhoneNumberInput) (*chime.DeletePhoneNumberOutput, error)
	DeletePhoneNumberWithContext(aws.Context, *chime.DeletePhoneNumberInput, ...request.Option) (*chime.DeletePhoneNumberOutput, error)
	DeletePhoneNumberRequest(*chime.DeletePhoneNumberInput) (*request.Request, *chime.DeletePhoneNumberOutput)

	DeleteVoiceConnector(*chime.DeleteVoiceConnectorInput) (*chime.DeleteVoiceConnectorOutput, error)
	DeleteVoiceConnectorWithContext(aws.Context, *chime.DeleteVoiceConnectorInput, ...request.Option) (*chime.DeleteVoiceConnectorOutput, error)
	DeleteVoiceConnectorRequest(*chime.DeleteVoiceConnectorInput) (*request.Request, *chime.DeleteVoiceConnectorOutput)

	DeleteVoiceConnectorOrigination(*chime.DeleteVoiceConnectorOriginationInput) (*chime.DeleteVoiceConnectorOriginationOutput, error)
	DeleteVoiceConnectorOriginationWithContext(aws.Context, *chime.DeleteVoiceConnectorOriginationInput, ...request.Option) (*chime.DeleteVoiceConnectorOriginationOutput, error)
	DeleteVoiceConnectorOriginationRequest(*chime.DeleteVoiceConnectorOriginationInput) (*request.Request, *chime.DeleteVoiceConnectorOriginationOutput)

	DeleteVoiceConnectorTermination(*chime.DeleteVoiceConnectorTerminationInput) (*chime.DeleteVoiceConnectorTerminationOutput, error)
	DeleteVoiceConnectorTerminationWithContext(aws.Context, *chime.DeleteVoiceConnectorTerminationInput, ...request.Option) (*chime.DeleteVoiceConnectorTerminationOutput, error)
	DeleteVoiceConnectorTerminationRequest(*chime.DeleteVoiceConnectorTerminationInput) (*request.Request, *chime.DeleteVoiceConnectorTerminationOutput)

	DeleteVoiceConnectorTerminationCredentials(*chime.DeleteVoiceConnectorTerminationCredentialsInput) (*chime.DeleteVoiceConnectorTerminationCredentialsOutput, error)
	DeleteVoiceConnectorTerminationCredentialsWithContext(aws.Context, *chime.DeleteVoiceConnectorTerminationCredentialsInput, ...request.Option) (*chime.DeleteVoiceConnectorTerminationCredentialsOutput, error)
	DeleteVoiceConnectorTerminationCredentialsRequest(*chime.DeleteVoiceConnectorTerminationCredentialsInput) (*request.Request, *chime.DeleteVoiceConnectorTerminationCredentialsOutput)

	DisassociatePhoneNumberFromUser(*chime.DisassociatePhoneNumberFromUserInput) (*chime.DisassociatePhoneNumberFromUserOutput, error)
	DisassociatePhoneNumberFromUserWithContext(aws.Context, *chime.DisassociatePhoneNumberFromUserInput, ...request.Option) (*chime.DisassociatePhoneNumberFromUserOutput, error)
	DisassociatePhoneNumberFromUserRequest(*chime.DisassociatePhoneNumberFromUserInput) (*request.Request, *chime.DisassociatePhoneNumberFromUserOutput)

	DisassociatePhoneNumbersFromVoiceConnector(*chime.DisassociatePhoneNumbersFromVoiceConnectorInput) (*chime.DisassociatePhoneNumbersFromVoiceConnectorOutput, error)
	DisassociatePhoneNumbersFromVoiceConnectorWithContext(aws.Context, *chime.DisassociatePhoneNumbersFromVoiceConnectorInput, ...request.Option) (*chime.DisassociatePhoneNumbersFromVoiceConnectorOutput, error)
	DisassociatePhoneNumbersFromVoiceConnectorRequest(*chime.DisassociatePhoneNumbersFromVoiceConnectorInput) (*request.Request, *chime.DisassociatePhoneNumbersFromVoiceConnectorOutput)

	GetAccount(*chime.GetAccountInput) (*chime.GetAccountOutput, error)
	GetAccountWithContext(aws.Context, *chime.GetAccountInput, ...request.Option) (*chime.GetAccountOutput, error)
	GetAccountRequest(*chime.GetAccountInput) (*request.Request, *chime.GetAccountOutput)

	GetAccountSettings(*chime.GetAccountSettingsInput) (*chime.GetAccountSettingsOutput, error)
	GetAccountSettingsWithContext(aws.Context, *chime.GetAccountSettingsInput, ...request.Option) (*chime.GetAccountSettingsOutput, error)
	GetAccountSettingsRequest(*chime.GetAccountSettingsInput) (*request.Request, *chime.GetAccountSettingsOutput)

	GetGlobalSettings(*chime.GetGlobalSettingsInput) (*chime.GetGlobalSettingsOutput, error)
	GetGlobalSettingsWithContext(aws.Context, *chime.GetGlobalSettingsInput, ...request.Option) (*chime.GetGlobalSettingsOutput, error)
	GetGlobalSettingsRequest(*chime.GetGlobalSettingsInput) (*request.Request, *chime.GetGlobalSettingsOutput)

	GetPhoneNumber(*chime.GetPhoneNumberInput) (*chime.GetPhoneNumberOutput, error)
	GetPhoneNumberWithContext(aws.Context, *chime.GetPhoneNumberInput, ...request.Option) (*chime.GetPhoneNumberOutput, error)
	GetPhoneNumberRequest(*chime.GetPhoneNumberInput) (*request.Request, *chime.GetPhoneNumberOutput)

	GetPhoneNumberOrder(*chime.GetPhoneNumberOrderInput) (*chime.GetPhoneNumberOrderOutput, error)
	GetPhoneNumberOrderWithContext(aws.Context, *chime.GetPhoneNumberOrderInput, ...request.Option) (*chime.GetPhoneNumberOrderOutput, error)
	GetPhoneNumberOrderRequest(*chime.GetPhoneNumberOrderInput) (*request.Request, *chime.GetPhoneNumberOrderOutput)

	GetUser(*chime.GetUserInput) (*chime.GetUserOutput, error)
	GetUserWithContext(aws.Context, *chime.GetUserInput, ...request.Option) (*chime.GetUserOutput, error)
	GetUserRequest(*chime.GetUserInput) (*request.Request, *chime.GetUserOutput)

	GetUserSettings(*chime.GetUserSettingsInput) (*chime.GetUserSettingsOutput, error)
	GetUserSettingsWithContext(aws.Context, *chime.GetUserSettingsInput, ...request.Option) (*chime.GetUserSettingsOutput, error)
	GetUserSettingsRequest(*chime.GetUserSettingsInput) (*request.Request, *chime.GetUserSettingsOutput)

	GetVoiceConnector(*chime.GetVoiceConnectorInput) (*chime.GetVoiceConnectorOutput, error)
	GetVoiceConnectorWithContext(aws.Context, *chime.GetVoiceConnectorInput, ...request.Option) (*chime.GetVoiceConnectorOutput, error)
	GetVoiceConnectorRequest(*chime.GetVoiceConnectorInput) (*request.Request, *chime.GetVoiceConnectorOutput)

	GetVoiceConnectorOrigination(*chime.GetVoiceConnectorOriginationInput) (*chime.GetVoiceConnectorOriginationOutput, error)
	GetVoiceConnectorOriginationWithContext(aws.Context, *chime.GetVoiceConnectorOriginationInput, ...request.Option) (*chime.GetVoiceConnectorOriginationOutput, error)
	GetVoiceConnectorOriginationRequest(*chime.GetVoiceConnectorOriginationInput) (*request.Request, *chime.GetVoiceConnectorOriginationOutput)

	GetVoiceConnectorTermination(*chime.GetVoiceConnectorTerminationInput) (*chime.GetVoiceConnectorTerminationOutput, error)
	GetVoiceConnectorTerminationWithContext(aws.Context, *chime.GetVoiceConnectorTerminationInput, ...request.Option) (*chime.GetVoiceConnectorTerminationOutput, error)
	GetVoiceConnectorTerminationRequest(*chime.GetVoiceConnectorTerminationInput) (*request.Request, *chime.GetVoiceConnectorTerminationOutput)

	GetVoiceConnectorTerminationHealth(*chime.GetVoiceConnectorTerminationHealthInput) (*chime.GetVoiceConnectorTerminationHealthOutput, error)
	GetVoiceConnectorTerminationHealthWithContext(aws.Context, *chime.GetVoiceConnectorTerminationHealthInput, ...request.Option) (*chime.GetVoiceConnectorTerminationHealthOutput, error)
	GetVoiceConnectorTerminationHealthRequest(*chime.GetVoiceConnectorTerminationHealthInput) (*request.Request, *chime.GetVoiceConnectorTerminationHealthOutput)

	InviteUsers(*chime.InviteUsersInput) (*chime.InviteUsersOutput, error)
	InviteUsersWithContext(aws.Context, *chime.InviteUsersInput, ...request.Option) (*chime.InviteUsersOutput, error)
	InviteUsersRequest(*chime.InviteUsersInput) (*request.Request, *chime.InviteUsersOutput)

	ListAccounts(*chime.ListAccountsInput) (*chime.ListAccountsOutput, error)
	ListAccountsWithContext(aws.Context, *chime.ListAccountsInput, ...request.Option) (*chime.ListAccountsOutput, error)
	ListAccountsRequest(*chime.ListAccountsInput) (*request.Request, *chime.ListAccountsOutput)

	ListAccountsPages(*chime.ListAccountsInput, func(*chime.ListAccountsOutput, bool) bool) error
	ListAccountsPagesWithContext(aws.Context, *chime.ListAccountsInput, func(*chime.ListAccountsOutput, bool) bool, ...request.Option) error

	ListPhoneNumberOrders(*chime.ListPhoneNumberOrdersInput) (*chime.ListPhoneNumberOrdersOutput, error)
	ListPhoneNumberOrdersWithContext(aws.Context, *chime.ListPhoneNumberOrdersInput, ...request.Option) (*chime.ListPhoneNumberOrdersOutput, error)
	ListPhoneNumberOrdersRequest(*chime.ListPhoneNumberOrdersInput) (*request.Request, *chime.ListPhoneNumberOrdersOutput)

	ListPhoneNumberOrdersPages(*chime.ListPhoneNumberOrdersInput, func(*chime.ListPhoneNumberOrdersOutput, bool) bool) error
	ListPhoneNumberOrdersPagesWithContext(aws.Context, *chime.ListPhoneNumberOrdersInput, func(*chime.ListPhoneNumberOrdersOutput, bool) bool, ...request.Option) error

	ListPhoneNumbers(*chime.ListPhoneNumbersInput) (*chime.ListPhoneNumbersOutput, error)
	ListPhoneNumbersWithContext(aws.Context, *chime.ListPhoneNumbersInput, ...request.Option) (*chime.ListPhoneNumbersOutput, error)
	ListPhoneNumbersRequest(*chime.ListPhoneNumbersInput) (*request.Request, *chime.ListPhoneNumbersOutput)

	ListPhoneNumbersPages(*chime.ListPhoneNumbersInput, func(*chime.ListPhoneNumbersOutput, bool) bool) error
	ListPhoneNumbersPagesWithContext(aws.Context, *chime.ListPhoneNumbersInput, func(*chime.ListPhoneNumbersOutput, bool) bool, ...request.Option) error

	ListUsers(*chime.ListUsersInput) (*chime.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *chime.ListUsersInput, ...request.Option) (*chime.ListUsersOutput, error)
	ListUsersRequest(*chime.ListUsersInput) (*request.Request, *chime.ListUsersOutput)

	ListUsersPages(*chime.ListUsersInput, func(*chime.ListUsersOutput, bool) bool) error
	ListUsersPagesWithContext(aws.Context, *chime.ListUsersInput, func(*chime.ListUsersOutput, bool) bool, ...request.Option) error

	ListVoiceConnectorTerminationCredentials(*chime.ListVoiceConnectorTerminationCredentialsInput) (*chime.ListVoiceConnectorTerminationCredentialsOutput, error)
	ListVoiceConnectorTerminationCredentialsWithContext(aws.Context, *chime.ListVoiceConnectorTerminationCredentialsInput, ...request.Option) (*chime.ListVoiceConnectorTerminationCredentialsOutput, error)
	ListVoiceConnectorTerminationCredentialsRequest(*chime.ListVoiceConnectorTerminationCredentialsInput) (*request.Request, *chime.ListVoiceConnectorTerminationCredentialsOutput)

	ListVoiceConnectors(*chime.ListVoiceConnectorsInput) (*chime.ListVoiceConnectorsOutput, error)
	ListVoiceConnectorsWithContext(aws.Context, *chime.ListVoiceConnectorsInput, ...request.Option) (*chime.ListVoiceConnectorsOutput, error)
	ListVoiceConnectorsRequest(*chime.ListVoiceConnectorsInput) (*request.Request, *chime.ListVoiceConnectorsOutput)

	ListVoiceConnectorsPages(*chime.ListVoiceConnectorsInput, func(*chime.ListVoiceConnectorsOutput, bool) bool) error
	ListVoiceConnectorsPagesWithContext(aws.Context, *chime.ListVoiceConnectorsInput, func(*chime.ListVoiceConnectorsOutput, bool) bool, ...request.Option) error

	LogoutUser(*chime.LogoutUserInput) (*chime.LogoutUserOutput, error)
	LogoutUserWithContext(aws.Context, *chime.LogoutUserInput, ...request.Option) (*chime.LogoutUserOutput, error)
	LogoutUserRequest(*chime.LogoutUserInput) (*request.Request, *chime.LogoutUserOutput)

	PutVoiceConnectorOrigination(*chime.PutVoiceConnectorOriginationInput) (*chime.PutVoiceConnectorOriginationOutput, error)
	PutVoiceConnectorOriginationWithContext(aws.Context, *chime.PutVoiceConnectorOriginationInput, ...request.Option) (*chime.PutVoiceConnectorOriginationOutput, error)
	PutVoiceConnectorOriginationRequest(*chime.PutVoiceConnectorOriginationInput) (*request.Request, *chime.PutVoiceConnectorOriginationOutput)

	PutVoiceConnectorTermination(*chime.PutVoiceConnectorTerminationInput) (*chime.PutVoiceConnectorTerminationOutput, error)
	PutVoiceConnectorTerminationWithContext(aws.Context, *chime.PutVoiceConnectorTerminationInput, ...request.Option) (*chime.PutVoiceConnectorTerminationOutput, error)
	PutVoiceConnectorTerminationRequest(*chime.PutVoiceConnectorTerminationInput) (*request.Request, *chime.PutVoiceConnectorTerminationOutput)

	PutVoiceConnectorTerminationCredentials(*chime.PutVoiceConnectorTerminationCredentialsInput) (*chime.PutVoiceConnectorTerminationCredentialsOutput, error)
	PutVoiceConnectorTerminationCredentialsWithContext(aws.Context, *chime.PutVoiceConnectorTerminationCredentialsInput, ...request.Option) (*chime.PutVoiceConnectorTerminationCredentialsOutput, error)
	PutVoiceConnectorTerminationCredentialsRequest(*chime.PutVoiceConnectorTerminationCredentialsInput) (*request.Request, *chime.PutVoiceConnectorTerminationCredentialsOutput)

	ResetPersonalPIN(*chime.ResetPersonalPINInput) (*chime.ResetPersonalPINOutput, error)
	ResetPersonalPINWithContext(aws.Context, *chime.ResetPersonalPINInput, ...request.Option) (*chime.ResetPersonalPINOutput, error)
	ResetPersonalPINRequest(*chime.ResetPersonalPINInput) (*request.Request, *chime.ResetPersonalPINOutput)

	RestorePhoneNumber(*chime.RestorePhoneNumberInput) (*chime.RestorePhoneNumberOutput, error)
	RestorePhoneNumberWithContext(aws.Context, *chime.RestorePhoneNumberInput, ...request.Option) (*chime.RestorePhoneNumberOutput, error)
	RestorePhoneNumberRequest(*chime.RestorePhoneNumberInput) (*request.Request, *chime.RestorePhoneNumberOutput)

	SearchAvailablePhoneNumbers(*chime.SearchAvailablePhoneNumbersInput) (*chime.SearchAvailablePhoneNumbersOutput, error)
	SearchAvailablePhoneNumbersWithContext(aws.Context, *chime.SearchAvailablePhoneNumbersInput, ...request.Option) (*chime.SearchAvailablePhoneNumbersOutput, error)
	SearchAvailablePhoneNumbersRequest(*chime.SearchAvailablePhoneNumbersInput) (*request.Request, *chime.SearchAvailablePhoneNumbersOutput)

	UpdateAccount(*chime.UpdateAccountInput) (*chime.UpdateAccountOutput, error)
	UpdateAccountWithContext(aws.Context, *chime.UpdateAccountInput, ...request.Option) (*chime.UpdateAccountOutput, error)
	UpdateAccountRequest(*chime.UpdateAccountInput) (*request.Request, *chime.UpdateAccountOutput)

	UpdateAccountSettings(*chime.UpdateAccountSettingsInput) (*chime.UpdateAccountSettingsOutput, error)
	UpdateAccountSettingsWithContext(aws.Context, *chime.UpdateAccountSettingsInput, ...request.Option) (*chime.UpdateAccountSettingsOutput, error)
	UpdateAccountSettingsRequest(*chime.UpdateAccountSettingsInput) (*request.Request, *chime.UpdateAccountSettingsOutput)

	UpdateGlobalSettings(*chime.UpdateGlobalSettingsInput) (*chime.UpdateGlobalSettingsOutput, error)
	UpdateGlobalSettingsWithContext(aws.Context, *chime.UpdateGlobalSettingsInput, ...request.Option) (*chime.UpdateGlobalSettingsOutput, error)
	UpdateGlobalSettingsRequest(*chime.UpdateGlobalSettingsInput) (*request.Request, *chime.UpdateGlobalSettingsOutput)

	UpdatePhoneNumber(*chime.UpdatePhoneNumberInput) (*chime.UpdatePhoneNumberOutput, error)
	UpdatePhoneNumberWithContext(aws.Context, *chime.UpdatePhoneNumberInput, ...request.Option) (*chime.UpdatePhoneNumberOutput, error)
	UpdatePhoneNumberRequest(*chime.UpdatePhoneNumberInput) (*request.Request, *chime.UpdatePhoneNumberOutput)

	UpdateUser(*chime.UpdateUserInput) (*chime.UpdateUserOutput, error)
	UpdateUserWithContext(aws.Context, *chime.UpdateUserInput, ...request.Option) (*chime.UpdateUserOutput, error)
	UpdateUserRequest(*chime.UpdateUserInput) (*request.Request, *chime.UpdateUserOutput)

	UpdateUserSettings(*chime.UpdateUserSettingsInput) (*chime.UpdateUserSettingsOutput, error)
	UpdateUserSettingsWithContext(aws.Context, *chime.UpdateUserSettingsInput, ...request.Option) (*chime.UpdateUserSettingsOutput, error)
	UpdateUserSettingsRequest(*chime.UpdateUserSettingsInput) (*request.Request, *chime.UpdateUserSettingsOutput)

	UpdateVoiceConnector(*chime.UpdateVoiceConnectorInput) (*chime.UpdateVoiceConnectorOutput, error)
	UpdateVoiceConnectorWithContext(aws.Context, *chime.UpdateVoiceConnectorInput, ...request.Option) (*chime.UpdateVoiceConnectorOutput, error)
	UpdateVoiceConnectorRequest(*chime.UpdateVoiceConnectorInput) (*request.Request, *chime.UpdateVoiceConnectorOutput)
}

var _ ChimeAPI = (*chime.Chime)(nil)
