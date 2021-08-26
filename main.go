package main

import (
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	list := []interface{}{cdktf.Fn_Reverse(jsii.Strings("world", "hello")), cdktf.Fn_Sum(jsii.Numbers(1, 2, 3))}

	// The code that defines your stack goes here
	cdktf.NewTerraformOutput(scope, jsii.String("test"), &cdktf.TerraformOutputConfig{
		Value: cdktf.Fn_Join("", list),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "test-tf-functions")

	app.Synth()
}
