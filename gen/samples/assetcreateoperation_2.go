//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeAssetCreate

package samples

func init() {

	sampleDataAssetCreateOperation[2] = `{
  "common_options": {
    "blacklist_authorities": [
      "1.2.35824"
    ],
    "blacklist_markets": [],
    "core_exchange_rate": {
      "base": {
        "amount": 2000000,
        "asset_id": "1.3.0"
      },
      "quote": {
        "amount": 1,
        "asset_id": "1.3.1"
      }
    },
    "description": "The opposite of a tip",
    "extensions": [],
    "flags": 2,
    "issuer_permissions": 79,
    "market_fee_percent": 0,
    "max_market_fee": "1000000000000000",
    "max_supply": "1000000000000000",
    "whitelist_authorities": [],
    "whitelist_markets": []
  },
  "extensions": [],
  "fee": {
    "amount": 500000000,
    "asset_id": "1.3.0"
  },
  "is_prediction_market": false,
  "issuer": "1.2.35824",
  "precision": 0,
  "symbol": "SHMACK"
}`

}

//end of file
