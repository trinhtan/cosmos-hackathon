package sunchain

import (
	"encoding/hex"

	"github.com/trinhtan/cosmos-hackathon/x/sunchain/types"
)

// encodeRequestParams returns byte array of encoded request coin price
func encodeRequestParams(coinName string, multiplier uint64) string {
	encoder := types.NewEncoder()
	encoder.EncodeString(coinName)
	encoder.EncodeU64(multiplier)
	return hex.EncodeToString(encoder.GetEncodedData())
}
