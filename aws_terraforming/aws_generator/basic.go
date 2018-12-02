package aws_generator

import "waze/terraform/terraform_utils"

type Generator interface {
	Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error)
	PostGenerateHook(resources []terraform_utils.TerraformResource) ([]terraform_utils.TerraformResource, error)
}

type BasicGenerator struct{}

func (BasicGenerator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	panic("implement me")
}

func (BasicGenerator) PostGenerateHook(resources []terraform_utils.TerraformResource) ([]terraform_utils.TerraformResource, error) {
	return resources, nil
}