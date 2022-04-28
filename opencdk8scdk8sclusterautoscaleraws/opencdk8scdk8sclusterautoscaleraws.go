// @opencdk8s/cdk8s-cluster-autoscaler-aws
package opencdk8scdk8sclusterautoscaleraws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/opencdk8s/cdk8s-cluster-autoscaler-aws-go/opencdk8scdk8sclusterautoscaleraws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/opencdk8s/cdk8s-cluster-autoscaler-aws-go/opencdk8scdk8sclusterautoscaleraws/internal"
)

// Aws External Dns Policy class ,help you add policy to your Iam Role for service account.
// Experimental.
type AwsClusterAutoScalerPolicyHelper interface {
}

// The jsii proxy struct for AwsClusterAutoScalerPolicyHelper
type jsiiProxy_AwsClusterAutoScalerPolicyHelper struct {
	_ byte // padding
}

// Experimental.
func NewAwsClusterAutoScalerPolicyHelper() AwsClusterAutoScalerPolicyHelper {
	_init_.Initialize()

	j := jsiiProxy_AwsClusterAutoScalerPolicyHelper{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-cluster-autoscaler-aws.AwsClusterAutoScalerPolicyHelper",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAwsClusterAutoScalerPolicyHelper_Override(a AwsClusterAutoScalerPolicyHelper) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-cluster-autoscaler-aws.AwsClusterAutoScalerPolicyHelper",
		nil, // no parameters
		a,
	)
}

// Experimental.
func AwsClusterAutoScalerPolicyHelper_AddPolicy(role interface{}) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-cluster-autoscaler-aws.AwsClusterAutoScalerPolicyHelper",
		"addPolicy",
		[]interface{}{role},
		&returns,
	)

	return returns
}

// Experimental.
type ClusterAutoScaler interface {
	constructs.Construct
	// Extra commands for controller.
	// Experimental.
	Command() *[]*string
	// service account for aws-load-balancer-controller.
	// Experimental.
	CreateServiceAccount() *bool
	// image for deployment.
	// Experimental.
	Image() *string
	// Namespace.
	// Experimental.
	Namespace() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Service Account Name.
	// Experimental.
	ServiceAccountName() *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClusterAutoScaler
type jsiiProxy_ClusterAutoScaler struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ClusterAutoScaler) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAutoScaler) CreateServiceAccount() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAutoScaler) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAutoScaler) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAutoScaler) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAutoScaler) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}


// Experimental.
func NewClusterAutoScaler(scope constructs.Construct, name *string, opts *ClusterAutoScalerOptions) ClusterAutoScaler {
	_init_.Initialize()

	j := jsiiProxy_ClusterAutoScaler{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-cluster-autoscaler-aws.ClusterAutoScaler",
		[]interface{}{scope, name, opts},
		&j,
	)

	return &j
}

// Experimental.
func NewClusterAutoScaler_Override(c ClusterAutoScaler, scope constructs.Construct, name *string, opts *ClusterAutoScalerOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-cluster-autoscaler-aws.ClusterAutoScaler",
		[]interface{}{scope, name, opts},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ClusterAutoScaler_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-cluster-autoscaler-aws.ClusterAutoScaler",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterAutoScaler) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type ClusterAutoScalerOptions struct {
	// Extra commands for controller.
	// Experimental.
	Command *[]*string `json:"command" yaml:"command"`
	// service account for aws-load-balancer-controller.
	// Experimental.
	CreateServiceAccount *bool `json:"createServiceAccount" yaml:"createServiceAccount"`
	// image for deployment.
	// Experimental.
	Image *string `json:"image" yaml:"image"`
	// Namespace.
	// Experimental.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// Service Account Name.
	// Experimental.
	ServiceAccountName *string `json:"serviceAccountName" yaml:"serviceAccountName"`
}

