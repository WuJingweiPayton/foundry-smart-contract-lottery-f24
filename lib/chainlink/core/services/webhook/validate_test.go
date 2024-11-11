package webhook_test

import (
	"testing"

	"github.com/manyminds/api2go/jsonapi"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/bridges"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/webhook"
	webhookmocks "github.com/smartcontractkit/chainlink/v2/core/services/webhook/mocks"
)

func TestValidatedWebJobSpec(t *testing.T) {
	t.Parallel()
	var tt = []struct {
		name      string
		toml      string
		mock      func(t *testing.T, eim *webhookmocks.ExternalInitiatorManager)
		assertion func(t *testing.T, spec job.Job, err error)
	}{
		{
			name: "valid spec",
			toml: `
			type            = "webhook"
			schemaVersion   = 1
            externalJobID           = "0eec7e1d-d0d2-476c-a1a8-72dfb6633f46"
			observationSource   = """
				ds          [type=http method=GET url="https://chain.link/ETH-USD"];
				ds_parse    [type=jsonparse path="data,price"];
				ds -> ds_parse;
			"""
			`,
			assertion: func(t *testing.T, s job.Job, err error) {
				require.NoError(t, err)
				require.NotNil(t, s.WebhookSpec)
				b, err := jsonapi.Marshal(s.WebhookSpec)
				require.NoError(t, err)
				var r job.WebhookSpec
				err = jsonapi.Unmarshal(b, &r)
				require.NoError(t, err)
				require.Equal(t, "0eec7e1d-d0d2-476c-a1a8-72dfb6633f46", s.ExternalJobID.String())
			},
		},
		{
			name: "invalid job name",
			toml: `
			type            = "webhookjob"
			schemaVersion   = 1
            externalJobID           = "0eec7e1d-d0d2-476c-a1a8-72dfb6633f46"
			observationSource   = """
			    ds          [type=http method=GET url="https://chain.link/ETH-USD"];
			    ds_parse    [type=jsonparse path="data,price"];
			    ds_multiply [type=multiply times=100];
			    ds -> ds_parse -> ds_multiply;
			"""
			`,
			assertion: func(t *testing.T, s job.Job, err error) {
				require.Error(t, err)
				assert.Equal(t, "unsupported type webhookjob", err.Error())
			},
		},
		{
			name: "missing jobID is fine (it will be autogenerated later)",
			toml: `
            type            = "webhook"
            schemaVersion   = 1
            observationSource   = """
                ds          [type=http method=GET url="https://chain.link/ETH-USD"];
                ds_parse    [type=jsonparse path="data,price"];
                ds_multiply [type=multiply times=100];
                ds -> ds_parse -> ds_multiply;
            """
            `,
			assertion: func(t *testing.T, s job.Job, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "with multiple external initiators and externalJobID",
			toml: `
            type            = "webhook"
            schemaVersion   = 1
			externalJobID   = "0eec7e1d-d0d2-476c-a1a8-72dfb6633f46"
			externalInitiators = [
				{ name = "foo", spec = '{"foo": 42}' },
				{ name = "bar", spec = '{"bar": 42}' }
			]
            observationSource   = """
                ds          [type=http method=GET url="https://chain.link/ETH-USD"];
                ds_parse    [type=jsonparse path="data,price"];
                ds -> ds_parse;
            """
            `,
			mock: func(t *testing.T, eim *webhookmocks.ExternalInitiatorManager) {
				eim.On("FindExternalInitiatorByName", mock.Anything, "foo").Return(bridges.ExternalInitiator{ID: 42}, nil).Once()
				eim.On("FindExternalInitiatorByName", mock.Anything, "bar").Return(bridges.ExternalInitiator{ID: 43}, nil).Once()
			},
			assertion: func(t *testing.T, s job.Job, err error) {
				require.NoError(t, err)
				assert.Len(t, s.WebhookSpec.ExternalInitiatorWebhookSpecs, 2)
				assert.Equal(t, int64(42), s.WebhookSpec.ExternalInitiatorWebhookSpecs[0].ExternalInitiatorID)
				assert.Equal(t, `{"foo": 42}`, s.WebhookSpec.ExternalInitiatorWebhookSpecs[0].Spec.Raw)

				assert.Equal(t, int64(43), s.WebhookSpec.ExternalInitiatorWebhookSpecs[1].ExternalInitiatorID)
				assert.Equal(t, `{"bar": 42}`, s.WebhookSpec.ExternalInitiatorWebhookSpecs[1].Spec.Raw)

				require.NotNil(t, s.WebhookSpec)
				b, err := jsonapi.Marshal(s.WebhookSpec)
				require.NoError(t, err)
				require.Equal(t, "0eec7e1d-d0d2-476c-a1a8-72dfb6633f46", s.ExternalJobID.String())
				var r job.WebhookSpec
				err = jsonapi.Unmarshal(b, &r)
				require.NoError(t, err)
			},
		},
		{
			name: "with external initiators that do not exist",
			toml: `
            type            = "webhook"
            schemaVersion   = 1
			externalInitiators = [
				{ name = "foo", spec = '{"foo": 42}' },
				{ name = "bar", spec = '{"bar": 42}' },
				{ name = "baz", spec = '{"baz": 42}' }
			]
            observationSource   = """
                ds          [type=http method=GET url="https://chain.link/ETH-USD"];
                ds_parse    [type=jsonparse path="data,price"];
                ds -> ds_parse;
            """
            `,
			mock: func(t *testing.T, eim *webhookmocks.ExternalInitiatorManager) {
				eim.On("FindExternalInitiatorByName", mock.Anything, "foo").Return(bridges.ExternalInitiator{ID: 42}, nil).Once()
				eim.On("FindExternalInitiatorByName", mock.Anything, "bar").Return(bridges.ExternalInitiator{}, errors.New("something exploded")).Once()
				eim.On("FindExternalInitiatorByName", mock.Anything, "baz").Return(bridges.ExternalInitiator{}, errors.New("something exploded")).Once()
			},
			assertion: func(t *testing.T, s job.Job, err error) {
				require.EqualError(t, err, "unable to find external initiator named bar: something exploded; unable to find external initiator named baz: something exploded")
			},
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := testutils.Context(t)
			eim := new(webhookmocks.ExternalInitiatorManager)
			if tc.mock != nil {
				tc.mock(t, eim)
			}
			s, err := webhook.ValidatedWebhookSpec(ctx, tc.toml, eim)
			tc.assertion(t, s, err)
		})
	}
}
