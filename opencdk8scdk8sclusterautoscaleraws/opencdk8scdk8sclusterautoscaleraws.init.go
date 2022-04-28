package opencdk8scdk8sclusterautoscaleraws

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-cluster-autoscaler-aws.AwsClusterAutoScalerPolicyHelper",
		reflect.TypeOf((*AwsClusterAutoScalerPolicyHelper)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AwsClusterAutoScalerPolicyHelper{}
		},
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-cluster-autoscaler-aws.ClusterAutoScaler",
		reflect.TypeOf((*ClusterAutoScaler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "createServiceAccount", GoGetter: "CreateServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountName", GoGetter: "ServiceAccountName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterAutoScaler{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-cluster-autoscaler-aws.ClusterAutoScalerOptions",
		reflect.TypeOf((*ClusterAutoScalerOptions)(nil)).Elem(),
	)
}
