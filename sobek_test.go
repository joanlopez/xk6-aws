package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/stretchr/testify/require"

	"go.k6.io/k6/js/modulestest"
)

func Test_fromSobekObject(t *testing.T) {
	rt := modulestest.NewRuntime(t)

	t.Run("NestedStructs(eventbridge.CreateConnectionInput)", func(t *testing.T) {
		root := rt.VU.Runtime().NewObject()

		createConnectionApiKeyAuthRequestParameters := rt.VU.Runtime().NewObject()
		createConnectionApiKeyAuthRequestParameters.Set(pascalToSnake("ApiKeyName"), "randomApiKeyName")
		createConnectionApiKeyAuthRequestParameters.Set(pascalToSnake("ApiKeyValue"), "randomApiKeyValue")

		connectionBodyParameter := rt.VU.Runtime().NewObject()
		connectionBodyParameter.Set(pascalToSnake("IsValueSecret"), true)
		connectionBodyParameter.Set(pascalToSnake("Key"), "randomKey")
		connectionBodyParameter.Set(pascalToSnake("Value"), "randomValue")

		bodyParameters := rt.VU.Runtime().NewArray(connectionBodyParameter)

		invocationHttpParameters := rt.VU.Runtime().NewObject()
		invocationHttpParameters.Set(pascalToSnake("BodyParameters"), bodyParameters)

		authParameters := rt.VU.Runtime().NewObject()
		authParameters.Set(pascalToSnake("ApiKeyAuthParameters"), createConnectionApiKeyAuthRequestParameters)
		authParameters.Set(pascalToSnake("InvocationHttpParameters"), invocationHttpParameters)

		root.Set(pascalToSnake("AuthParameters"), authParameters)
		root.Set(pascalToSnake("AuthorizationType"), "BASIC")

		in := &eventbridge.CreateConnectionInput{}
		require.NoError(t, fromSobekObject(rt.VU.Runtime(), root, in))

		require.Equal(t, "randomApiKeyName", *in.AuthParameters.ApiKeyAuthParameters.ApiKeyName)
		require.Equal(t, "randomKey", *in.AuthParameters.InvocationHttpParameters.BodyParameters[0].Key)
		require.Equal(t, types.ConnectionAuthorizationTypeBasic, in.AuthorizationType)
	})
}
