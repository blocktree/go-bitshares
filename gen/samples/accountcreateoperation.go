//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeAccountCreate

package samples

import (
	"github.com/denkhaus/bitshares/gen/data"
	"github.com/denkhaus/bitshares/types"
)

func init() {
	data.OpSampleMap[types.OperationTypeAccountCreate] =
		sampleDataAccountCreateOperation
}

var sampleDataAccountCreateOperation = `{
  "active": {
    "account_auths": [],
    "address_auths": [],
    "key_auths": [
      [
        "BTS6K35Bajw29N4fjP4XADHtJ7bEj2xHJ8CoY2P2s1igXTB5oMBhR",
        1
      ]
    ],
    "weight_threshold": 1
  },
  "extensions": {
    "active_special_authority": [
      1,
      {
        "asset": "1.3.333",
        "num_top_holders": 15
      }
    ],
    "buyback_options": {
      "asset_to_buy": "1.3.127",
      "asset_to_buy_issuer": "1.2.31",
      "markets": [
        "1.3.20"
      ]
    },
    "null_ext": {},
    "owner_special_authority": [
      1,
      {
        "asset": "1.3.127",
        "num_top_holders": 10
      }
    ]
  },
  "fee": {
    "amount": 19029492,
    "asset_id": "1.3.0"
  },
  "name": "daniel.pan",
  "options": {
    "extensions": [],
    "memo_key": "BTS6K35Bajw29N4fjP4XADHtJ7bEj2xHJ8CoY2P2s1igXTB5oMBhR",
    "num_committee": 0,
    "num_witness": 0,
    "votes": [],
    "voting_account": "1.2.0"
  },
  "owner": {
    "account_auths": [],
    "address_auths": [],
    "key_auths": [
      [
        "BTS8YMTPBafG4cC5qje9UGoRBxU8XdMjJUHUwHNKgM6AiZA79yUiP",
        1
      ]
    ],
    "weight_threshold": 1
  },
  "referrer": "1.2.332",
  "referrer_percent": 100,
  "registrar": "1.2.332"
}`

//end of file