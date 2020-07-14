package v1alpha3

import json "github.com/json-iterator/go"

// S3BucketPolicy defines the policy for the S3 Bucket being created.
type S3BucketPolicy struct {
	// This is the current IAM policy version
	PolicyVersion string `json:"Version"`

	// This is the policy's optional identifier
	PolicyID string `json:"Id,omitempty"`

	PolicyStatement []S3BucketPolicyStatement `json:"Statement"`
}

// S3BucketPolicyStatement defines an individual statement within the
// S3BucketPolicy
type S3BucketPolicyStatement struct {
	// Optional identifier for this statement, must be unique within the
	// policy if provided.
	StatementID string `json:"Sid"`

	// The effect is required and specifies whether the statement results
	// in an allow or an explicit deny. Valid values for Effect are Allow and Deny.
	Effect string `json:"Effect"`

	// Used with the S3 policy to specify the principal that is allowed
	// or denied access to a resource.
	Principal S3BucketPrincipal `json:"Principal"`

	// Each element of the PolicyAction array describes describes the specific
	// action or actions that will be allowed or denied with this PolicyStatement.
	PolicyAction []string `json:"Action"`

	// This flag indicates that this policy should apply to the IAMUsername
	// that was either passed in or created for this bucket.
	ApplyToIAMUser bool `json:"EffectIAMUser,omitempty"`

	// The paths on which this resource will apply
	ResourcePath []string `json:"Resource"`
}

// S3BucketPrincipal defines the principal users affected by
// the S3BucketPolicyStatement
type S3BucketPrincipal struct {
	// This flag indicates if the policy should be made available
	// to all anonymous users.
	AllowAnon bool `json:"AnonymousAccess,omitempty"`

	// This list contains the all of the AWS IAM users which are affected
	// by the policy statement
	AWSPrincipal []string `json:"AWS,omitempty"`
}

// MarshalJSON is the custom marshaller for the S3BucketPrincipal
func (p *S3BucketPrincipal) MarshalJSON() ([]byte, error) {
	if p.AllowAnon {
		return json.Marshal("*")
	}
	return json.Marshal(p)
}
