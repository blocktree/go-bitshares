//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeAssetUpdate

package samples

func init() {

	sampleDataAssetUpdateOperation[2] = `{
  "asset_to_update": "1.3.856",
  "extensions": [],
  "fee": {
    "amount": 29368345,
    "asset_id": "1.3.0"
  },
  "issuer": "1.2.96397",
  "new_options": {
    "blacklist_authorities": [
      "1.2.96397"
    ],
    "blacklist_markets": [],
    "core_exchange_rate": {
      "base": {
        "amount": 100000,
        "asset_id": "1.3.0"
      },
      "quote": {
        "amount": 200000000,
        "asset_id": "1.3.856"
      }
    },
    "description": "{\"main\":\"OpenLedger NuShares backed asset - https://openledger.info/\",\"market\":\"OPEN.BTC\"}",
    "extensions": [],
    "flags": 131,
    "issuer_permissions": 79,
    "market_fee_percent": 20,
    "max_market_fee": "1000000000000000",
    "max_supply": "1000000000000000",
    "whitelist_authorities": [],
    "whitelist_markets": []
  }
}`

}

//end of file
