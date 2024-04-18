// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new OpsItem. You must have permission in Identity and Access
// Management (IAM) to create a new OpsItem. For more information, see Set up
// OpsCenter (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-setup.html)
// in the Amazon Web Services Systems Manager User Guide. Operations engineers and
// IT professionals use Amazon Web Services Systems Manager OpsCenter to view,
// investigate, and remediate operational issues impacting the performance and
// health of their Amazon Web Services resources. For more information, see Amazon
// Web Services Systems Manager OpsCenter (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html)
// in the Amazon Web Services Systems Manager User Guide.
func (c *Client) CreateOpsItem(ctx context.Context, params *CreateOpsItemInput, optFns ...func(*Options)) (*CreateOpsItemOutput, error) {
	if params == nil {
		params = &CreateOpsItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOpsItem", params, optFns, c.addOperationCreateOpsItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOpsItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOpsItemInput struct {

	// User-defined text that contains information about the OpsItem, in Markdown
	// format. Provide enough information so that users viewing this OpsItem for the
	// first time understand the issue.
	//
	// This member is required.
	Description *string

	// The origin of the OpsItem, such as Amazon EC2 or Systems Manager. The source
	// name can't contain the following strings: aws , amazon , and amzn .
	//
	// This member is required.
	Source *string

	// A short heading that describes the nature of the OpsItem and the impacted
	// resource.
	//
	// This member is required.
	Title *string

	// The target Amazon Web Services account where you want to create an OpsItem. To
	// make this call, your account must be configured to work with OpsItems across
	// accounts. For more information, see Set up OpsCenter (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-setup.html)
	// in the Amazon Web Services Systems Manager User Guide.
	AccountId *string

	// The time a runbook workflow ended. Currently reported only for the OpsItem type
	// /aws/changerequest .
	ActualEndTime *time.Time

	// The time a runbook workflow started. Currently reported only for the OpsItem
	// type /aws/changerequest .
	ActualStartTime *time.Time

	// Specify a category to assign to an OpsItem.
	Category *string

	// The Amazon Resource Name (ARN) of an SNS topic where notifications are sent
	// when this OpsItem is edited or changed.
	Notifications []types.OpsItemNotification

	// Operational data is custom data that provides useful reference details about
	// the OpsItem. For example, you can specify log files, error strings, license
	// keys, troubleshooting tips, or other relevant data. You enter operational data
	// as key-value pairs. The key has a maximum length of 128 characters. The value
	// has a maximum size of 20 KB. Operational data keys can't begin with the
	// following: amazon , aws , amzn , ssm , /amazon , /aws , /amzn , /ssm . You can
	// choose to make the data searchable by other users in the account or you can
	// restrict search access. Searchable data means that all users with access to the
	// OpsItem Overview page (as provided by the DescribeOpsItems API operation) can
	// view and search on the specified data. Operational data that isn't searchable is
	// only viewable by users who have access to the OpsItem (as provided by the
	// GetOpsItem API operation). Use the /aws/resources key in OperationalData to
	// specify a related resource in the request. Use the /aws/automations key in
	// OperationalData to associate an Automation runbook with the OpsItem. To view
	// Amazon Web Services CLI example commands that use these keys, see Create
	// OpsItems manually (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-manually-create-OpsItems.html)
	// in the Amazon Web Services Systems Manager User Guide.
	OperationalData map[string]types.OpsItemDataValue

	// The type of OpsItem to create. Systems Manager supports the following types of
	// OpsItems:
	//   - /aws/issue This type of OpsItem is used for default OpsItems created by
	//   OpsCenter.
	//   - /aws/changerequest This type of OpsItem is used by Change Manager for
	//   reviewing and approving or rejecting change requests.
	//   - /aws/insight This type of OpsItem is used by OpsCenter for aggregating and
	//   reporting on duplicate OpsItems.
	OpsItemType *string

	// The time specified in a change request for a runbook workflow to end. Currently
	// supported only for the OpsItem type /aws/changerequest .
	PlannedEndTime *time.Time

	// The time specified in a change request for a runbook workflow to start.
	// Currently supported only for the OpsItem type /aws/changerequest .
	PlannedStartTime *time.Time

	// The importance of this OpsItem in relation to other OpsItems in the system.
	Priority *int32

	// One or more OpsItems that share something in common with the current OpsItems.
	// For example, related OpsItems can include OpsItems with similar error messages,
	// impacted resources, or statuses for the impacted resource.
	RelatedOpsItems []types.RelatedOpsItem

	// Specify a severity to assign to an OpsItem.
	Severity *string

	// Optional metadata that you assign to a resource. Tags use a key-value pair. For
	// example: Key=Department,Value=Finance To add tags to a new OpsItem, a user must
	// have IAM permissions for both the ssm:CreateOpsItems operation and the
	// ssm:AddTagsToResource operation. To add tags to an existing OpsItem, use the
	// AddTagsToResource operation.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateOpsItemOutput struct {

	// The OpsItem Amazon Resource Name (ARN).
	OpsItemArn *string

	// The ID of the OpsItem.
	OpsItemId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOpsItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateOpsItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateOpsItem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateOpsItem"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateOpsItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOpsItem(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateOpsItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateOpsItem",
	}
}