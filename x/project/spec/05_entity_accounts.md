# Entity Accounts

The project module has a number of entity/project accounts created per
entity/project:

- `InitiatingNodePayFees`: when an oracle service is delivered, \[by default\] 10% of
  the oracle fee payment is sent to this address.
- `KaijuPayFees`: when an oracle service is delivered, \[by default\] 80% of the
  oracle fee payment is sent to this address.
- `KaijuFees`: at project payout, any amount in the `KaijuPayFees` account is moved to this
  account address. In turn, any amount now in KaijuFees is sent to pay the network fee to kaiju (the account associated with a DID which is
  specified in the genesis file).
- `<projectDid>`: this is the project's funding account, from which tokens
  are paid to the oracle service.
- Additionally, one entity/project account for each agent is also created. Any
  payments intended for the particular agent is sent to their corresponding
  entity account. At project payout, the agent's tokens can be withdrawn.

The fee defaults can be configured from the Genesis file.

Example accounts for a project with DID `did:kaiju:U7GK8p8rVhJMKhBVRCJJ8c` and one
agent with DID `did:kaiju:RYLHkfNpbA8Losy68jt4yF`:

```json
{
  "InitiatingNodePayFees": "kaiju1xvjy68xrrtxnypwev9r8tmjys9wk0zkkspzjmq",
  "KaijuPayFees": "kaiju1udgxtf6yd09mwnnd0ljpmeq4vnyhxdg03uvne3",
  "KaijuFees": "kaiju1ff9we62w6eyes7wscjup3p40vy4uz0sa7j0ajc",
  "did:kaiju:U7GK8p8rVhJMKhBVRCJJ8c": "kaiju1rmkak6t606wczsps9ytpga3z4nre4z3nwc04p8",
  "did:kaiju:RYLHkfNpbA8Losy68jt4yF": "kaiju18nmp3w2xwz0rzkh8sdwkyz4fzjegemtx9vw3ky"
}
```
