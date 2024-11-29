// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// PolicyFulcioSubjectApplyConfiguration represents a declarative configuration of the PolicyFulcioSubject type for use
// with apply.
type PolicyFulcioSubjectApplyConfiguration struct {
	OIDCIssuer  *string `json:"oidcIssuer,omitempty"`
	SignedEmail *string `json:"signedEmail,omitempty"`
}

// PolicyFulcioSubjectApplyConfiguration constructs a declarative configuration of the PolicyFulcioSubject type for use with
// apply.
func PolicyFulcioSubject() *PolicyFulcioSubjectApplyConfiguration {
	return &PolicyFulcioSubjectApplyConfiguration{}
}

// WithOIDCIssuer sets the OIDCIssuer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OIDCIssuer field is set to the value of the last call.
func (b *PolicyFulcioSubjectApplyConfiguration) WithOIDCIssuer(value string) *PolicyFulcioSubjectApplyConfiguration {
	b.OIDCIssuer = &value
	return b
}

// WithSignedEmail sets the SignedEmail field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SignedEmail field is set to the value of the last call.
func (b *PolicyFulcioSubjectApplyConfiguration) WithSignedEmail(value string) *PolicyFulcioSubjectApplyConfiguration {
	b.SignedEmail = &value
	return b
}