package schema

import (
	"github.com/hashicorp/hcl-lang/schema"
)

// TODO: this should be a key value pair map
var EnvSchema = &schema.BodySchema{
	Attributes: map[string]*schema.AttributeSchema{},
}
