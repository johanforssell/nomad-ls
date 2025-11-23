package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/zclconf/go-cty/cty"
)

// TODO: update docs and constraints
var VariableSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{
		"type": {
			Description: lang.Markdown("The type of HCL variable: `string`, `number`, `bool`."),
			Constraint:  &schema.LiteralType{Type: cty.String},
		},
		"default": {
			Description: lang.Markdown("The default value used when no value for this variable is provided."),
			Constraint:  &schema.LiteralType{Type: cty.String},
		},
	},
}
