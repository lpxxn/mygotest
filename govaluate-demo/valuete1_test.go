package govaluate_demo

import (
	"testing"

	"github.com/Knetic/govaluate"
	"github.com/stretchr/testify/assert"
)

func TestBool1(t *testing.T) {
	expression, err := govaluate.NewEvaluableExpression("foo > 5")
	assert.Nil(t, err)

	param := make(map[string]interface{})
	param["foo"] = 3

	result, err := expression.Evaluate(param)
	assert.Nil(t, err)
	assert.False(t, result.(bool))

	param["foo"] = 6
	result, err = expression.Evaluate(param)
	assert.Nil(t, err)
	assert.True(t, result.(bool))

	expression, err = govaluate.NewEvaluableExpression("10 > 0")
	result, err = expression.Evaluate(nil)
	t.Log(result)
}
