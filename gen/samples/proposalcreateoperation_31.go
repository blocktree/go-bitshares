//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeProposalCreate

package samples

func init() {

	sampleDataProposalCreateOperation[31] = `{
  "expiration_time": "2016-05-10T13:33:11",
  "extensions": [],
  "fee": {
    "amount": 2287446,
    "asset_id": "1.3.0"
  },
  "fee_paying_account": "1.2.31223",
  "proposed_ops": [
    {
      "op": [
        6,
        {
          "account": "1.2.110000",
          "active": {
            "account_auths": [
              [
                "1.2.29528",
                1
              ],
              [
                "1.2.31223",
                1
              ],
              [
                "1.2.109870",
                1
              ]
            ],
            "address_auths": [],
            "key_auths": [],
            "weight_threshold": 2
          },
          "extensions": {},
          "fee": {
            "amount": 24708,
            "asset_id": "1.3.0"
          },
          "new_options": {
            "extensions": [],
            "memo_key": "BTS8ZEgugZRT4wAUAJTiXnp1A8xi8goNpsFnK1F8suqxWLpKH49iA",
            "num_committee": 0,
            "num_witness": 0,
            "votes": [],
            "voting_account": "1.2.5"
          },
          "owner": {
            "account_auths": [
              [
                "1.2.29528",
                1
              ],
              [
                "1.2.31223",
                1
              ],
              [
                "1.2.109870",
                1
              ]
            ],
            "address_auths": [],
            "key_auths": [],
            "weight_threshold": 2
          }
        }
      ]
    }
  ]
}`

}

//end of file
