package simulation

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"mods.irisnet.org/modules/token/types"
	v1 "mods.irisnet.org/modules/token/types/v1"
)

func TestHandleIssueTokenDeliveryError(t *testing.T) {
	unexpectedErr := errors.New("unexpected delivery failure")
	tests := []struct {
		name        string
		err         error
		wantComment string
		wantErr     error
	}{
		{
			name:        "disabled issuance is an expected no-op",
			err:         fmt.Errorf("deliver tx: %w", types.ErrIssueTokenDisabled),
			wantComment: "token issuance is disabled",
		},
		{
			name:        "unexpected delivery error is preserved",
			err:         unexpectedErr,
			wantComment: "unable to deliver tx",
			wantErr:     unexpectedErr,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opMsg, futureOps, err := handleIssueTokenDeliveryError(v1.TypeMsgIssueToken, tc.err)

			require.False(t, opMsg.OK)
			require.Equal(t, types.ModuleName, opMsg.Route)
			require.Equal(t, v1.TypeMsgIssueToken, opMsg.Name)
			require.Equal(t, tc.wantComment, opMsg.Comment)
			require.Nil(t, futureOps)
			if tc.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tc.wantErr)
			}
		})
	}
}
