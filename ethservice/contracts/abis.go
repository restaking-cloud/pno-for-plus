package contracts

const (
	MULTICALL3_CONTRACT_ABI = `[
		{
		  "inputs": [
			{
			  "components": [
				{
				  "internalType": "address",
				  "name": "target",
				  "type": "address"
				},
				{
				  "internalType": "bytes",
				  "name": "callData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Call[]",
			  "name": "calls",
			  "type": "tuple[]"
			}
		  ],
		  "name": "aggregate",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "blockNumber",
			  "type": "uint256"
			},
			{
			  "internalType": "bytes[]",
			  "name": "returnData",
			  "type": "bytes[]"
			}
		  ],
		  "stateMutability": "payable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "components": [
				{
				  "internalType": "address",
				  "name": "target",
				  "type": "address"
				},
				{
				  "internalType": "bool",
				  "name": "allowFailure",
				  "type": "bool"
				},
				{
				  "internalType": "bytes",
				  "name": "callData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Call3[]",
			  "name": "calls",
			  "type": "tuple[]"
			}
		  ],
		  "name": "aggregate3",
		  "outputs": [
			{
			  "components": [
				{
				  "internalType": "bool",
				  "name": "success",
				  "type": "bool"
				},
				{
				  "internalType": "bytes",
				  "name": "returnData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Result[]",
			  "name": "returnData",
			  "type": "tuple[]"
			}
		  ],
		  "stateMutability": "payable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "components": [
				{
				  "internalType": "address",
				  "name": "target",
				  "type": "address"
				},
				{
				  "internalType": "bool",
				  "name": "allowFailure",
				  "type": "bool"
				},
				{
				  "internalType": "uint256",
				  "name": "value",
				  "type": "uint256"
				},
				{
				  "internalType": "bytes",
				  "name": "callData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Call3Value[]",
			  "name": "calls",
			  "type": "tuple[]"
			}
		  ],
		  "name": "aggregate3Value",
		  "outputs": [
			{
			  "components": [
				{
				  "internalType": "bool",
				  "name": "success",
				  "type": "bool"
				},
				{
				  "internalType": "bytes",
				  "name": "returnData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Result[]",
			  "name": "returnData",
			  "type": "tuple[]"
			}
		  ],
		  "stateMutability": "payable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "components": [
				{
				  "internalType": "address",
				  "name": "target",
				  "type": "address"
				},
				{
				  "internalType": "bytes",
				  "name": "callData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Call[]",
			  "name": "calls",
			  "type": "tuple[]"
			}
		  ],
		  "name": "blockAndAggregate",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "blockNumber",
			  "type": "uint256"
			},
			{
			  "internalType": "bytes32",
			  "name": "blockHash",
			  "type": "bytes32"
			},
			{
			  "components": [
				{
				  "internalType": "bool",
				  "name": "success",
				  "type": "bool"
				},
				{
				  "internalType": "bytes",
				  "name": "returnData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Result[]",
			  "name": "returnData",
			  "type": "tuple[]"
			}
		  ],
		  "stateMutability": "payable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getBasefee",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "basefee",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "blockNumber",
			  "type": "uint256"
			}
		  ],
		  "name": "getBlockHash",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "blockHash",
			  "type": "bytes32"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getBlockNumber",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "blockNumber",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getChainId",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "chainid",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getCurrentBlockCoinbase",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "coinbase",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getCurrentBlockDifficulty",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "difficulty",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getCurrentBlockGasLimit",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "gaslimit",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getCurrentBlockTimestamp",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "timestamp",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "addr",
			  "type": "address"
			}
		  ],
		  "name": "getEthBalance",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "balance",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getLastBlockHash",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "blockHash",
			  "type": "bytes32"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bool",
			  "name": "requireSuccess",
			  "type": "bool"
			},
			{
			  "components": [
				{
				  "internalType": "address",
				  "name": "target",
				  "type": "address"
				},
				{
				  "internalType": "bytes",
				  "name": "callData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Call[]",
			  "name": "calls",
			  "type": "tuple[]"
			}
		  ],
		  "name": "tryAggregate",
		  "outputs": [
			{
			  "components": [
				{
				  "internalType": "bool",
				  "name": "success",
				  "type": "bool"
				},
				{
				  "internalType": "bytes",
				  "name": "returnData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Result[]",
			  "name": "returnData",
			  "type": "tuple[]"
			}
		  ],
		  "stateMutability": "payable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bool",
			  "name": "requireSuccess",
			  "type": "bool"
			},
			{
			  "components": [
				{
				  "internalType": "address",
				  "name": "target",
				  "type": "address"
				},
				{
				  "internalType": "bytes",
				  "name": "callData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Call[]",
			  "name": "calls",
			  "type": "tuple[]"
			}
		  ],
		  "name": "tryBlockAndAggregate",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "blockNumber",
			  "type": "uint256"
			},
			{
			  "internalType": "bytes32",
			  "name": "blockHash",
			  "type": "bytes32"
			},
			{
			  "components": [
				{
				  "internalType": "bool",
				  "name": "success",
				  "type": "bool"
				},
				{
				  "internalType": "bytes",
				  "name": "returnData",
				  "type": "bytes"
				}
			  ],
			  "internalType": "struct Multicall3.Result[]",
			  "name": "returnData",
			  "type": "tuple[]"
			}
		  ],
		  "stateMutability": "payable",
		  "type": "function"
		}
	  ]`
	K2_LENDING_CONTRACT_ABI       = `[{"type":"constructor","inputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"DEFAULT_ADMIN_ROLE","inputs":[],"outputs":[{"name":"","type":"bytes32","internalType":"bytes32"}],"stateMutability":"view"},{"type":"function","name":"allowance","inputs":[{"name":"owner","type":"address","internalType":"address"},{"name":"spender","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"approve","inputs":[{"name":"spender","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"assumedLiquidity","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"balanceOf","inputs":[{"name":"account","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"banNodeOperator","inputs":[{"name":"bannedNO","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"batchNodeOperatorDeposit","inputs":[{"name":"_blsPubkeys","type":"bytes[]","internalType":"bytes[]"},{"name":"_payoutRecipients","type":"address[]","internalType":"address[]"},{"name":"_blsSignatures","type":"bytes[]","internalType":"bytes[]"},{"name":"_ecdsaSignatures","type":"tuple[]","internalType":"struct IProposerRegistry.SignatureECDSA[]","components":[{"name":"v","type":"uint8","internalType":"uint8"},{"name":"r","type":"bytes32","internalType":"bytes32"},{"name":"s","type":"bytes32","internalType":"bytes32"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"blsPublicKeyToKicked","inputs":[{"name":"","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"blsPublicKeyToNodeOperator","inputs":[{"name":"","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"borrow","inputs":[{"name":"debtPositionType","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"},{"name":"designatedVerifier","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","internalType":"uint256"},{"name":"rstConfigParams","type":"tuple","internalType":"struct IK2Lending.RSTConfigParams","components":[{"name":"mintRST","type":"bool","internalType":"bool"},{"name":"initialSupply","type":"uint256","internalType":"uint256"},{"name":"recipientOfRemainingSupply","type":"address","internalType":"address"},{"name":"percentageContributionToIncentives","type":"uint256","internalType":"uint256"},{"name":"symbol","type":"string","internalType":"string"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"borrowDuration","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"borrowFor","inputs":[{"name":"debtor","type":"address","internalType":"address"},{"name":"debtPositionType","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"},{"name":"designatedVerifier","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","internalType":"uint256"},{"name":"rstConfigParams","type":"tuple","internalType":"struct IK2Lending.RSTConfigParams","components":[{"name":"mintRST","type":"bool","internalType":"bool"},{"name":"initialSupply","type":"uint256","internalType":"uint256"},{"name":"recipientOfRemainingSupply","type":"address","internalType":"address"},{"name":"percentageContributionToIncentives","type":"uint256","internalType":"uint256"},{"name":"symbol","type":"string","internalType":"string"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"borrowedLiquidity","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"claimKETH","inputs":[{"name":"lender","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"claimableKETH","inputs":[{"name":"lender","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"claimableKETHForNodeOperator","inputs":[{"name":"nodeOperator","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"daoAddress","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"debtPositions","inputs":[{"name":"","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"debtor","type":"address","internalType":"address"},{"name":"hook","type":"address","internalType":"address"},{"name":"designatedVerifier","type":"address","internalType":"address"},{"name":"designatedVerifierURI","type":"string","internalType":"string"},{"name":"principalAmount","type":"uint256","internalType":"uint256"},{"name":"interestPerSec_RAY","type":"uint256","internalType":"uint256"},{"name":"endTimestamp","type":"uint256","internalType":"uint256"},{"name":"slashAmount","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"debtors","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"decimals","inputs":[],"outputs":[{"name":"","type":"uint8","internalType":"uint8"}],"stateMutability":"view"},{"type":"function","name":"decreaseAllowance","inputs":[{"name":"spender","type":"address","internalType":"address"},{"name":"subtractedValue","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"delegatedClaim","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"deposit","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"depositFor","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"recipient","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"getBorrowedLiquidity","inputs":[],"outputs":[{"name":"res","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getDebtor","inputs":[{"name":"debtor","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"tuple","internalType":"struct IK2Lending.DebtPosition","components":[{"name":"debtor","type":"address","internalType":"address"},{"name":"hook","type":"address","internalType":"address"},{"name":"designatedVerifier","type":"address","internalType":"address"},{"name":"designatedVerifierURI","type":"string","internalType":"string"},{"name":"principalAmount","type":"uint256","internalType":"uint256"},{"name":"interestPerSec_RAY","type":"uint256","internalType":"uint256"},{"name":"endTimestamp","type":"uint256","internalType":"uint256"},{"name":"slashAmount","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","internalType":"uint256"}]}],"stateMutability":"view"},{"type":"function","name":"getOutstandingInterest","inputs":[{"name":"debtor","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getRemainingDuration","inputs":[{"name":"debtor","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getRoleAdmin","inputs":[{"name":"role","type":"bytes32","internalType":"bytes32"}],"outputs":[{"name":"","type":"bytes32","internalType":"bytes32"}],"stateMutability":"view"},{"type":"function","name":"getTotalBorrowableAmount","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getTotalBorrowableAmountWithMaxBorrowRatio","inputs":[{"name":"debtPositionType","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"grantRole","inputs":[{"name":"role","type":"bytes32","internalType":"bytes32"},{"name":"account","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"hasRole","inputs":[{"name":"role","type":"bytes32","internalType":"bytes32"},{"name":"account","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"increaseAllowance","inputs":[{"name":"spender","type":"address","internalType":"address"},{"name":"addedValue","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"increaseDebt","inputs":[{"name":"debtPositionType","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"},{"name":"designatedVerifier","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","internalType":"uint256"},{"name":"resetDuration","type":"bool","internalType":"bool"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"increaseDebtFor","inputs":[{"name":"debtor","type":"address","internalType":"address"},{"name":"debtPositionType","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"},{"name":"designatedVerifier","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","internalType":"uint256"},{"name":"resetDuration","type":"bool","internalType":"bool"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"initialize","inputs":[{"name":"_configurator","type":"address","internalType":"address"},{"name":"_daoAddress","type":"address","internalType":"address"},{"name":"_k2Incentives","type":"address","internalType":"address"},{"name":"_name","type":"string","internalType":"string"},{"name":"_symbol","type":"string","internalType":"string"},{"name":"_proposerRegistry","type":"address","internalType":"address"},{"name":"_nodeOperatorInclusionList","type":"address","internalType":"address"},{"name":"_maxSlashableRatio_RAY","type":"uint256","internalType":"uint256"},{"name":"_terminatePenalty_RAY","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"interestRateModelByType","inputs":[{"name":"","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"isNodeOperatorBanned","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"k2Incentives","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"lenders","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"cumulativeKethPerShareLU_RAY","type":"uint256","internalType":"uint256"},{"name":"kethEarned","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"liquidate","inputs":[{"name":"debtor","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"maxSlashableRatio_RAY","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"name","inputs":[],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"nodeOperatorClaim","inputs":[{"name":"_recipientOverride","type":"address","internalType":"address"},{"name":"_nodeOperator","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"nodeOperatorInclusionList","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"nodeOperatorKick","inputs":[{"name":"_reporter","type":"address","internalType":"address"},{"name":"_blsPubkey","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"nodeOperatorToBlsPublicKeyCount","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"nodeOperatorToLendPosition","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"cumulativeKethPerShareLU_RAY","type":"uint256","internalType":"uint256"},{"name":"kethEarned","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"nodeOperatorToPayoutRecipient","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"nodeOperatorWithdraw","inputs":[{"name":"_recipientOverride","type":"address","internalType":"address"},{"name":"_nodeOperator","type":"address","internalType":"address"},{"name":"_blsPubkey","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"proposerRegistry","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"renounceRole","inputs":[{"name":"role","type":"bytes32","internalType":"bytes32"},{"name":"account","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"revokeRole","inputs":[{"name":"role","type":"bytes32","internalType":"bytes32"},{"name":"account","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"rstModule","inputs":[],"outputs":[{"name":"","type":"address","internalType":"contract RSTModule"}],"stateMutability":"view"},{"type":"function","name":"setBorrowDuration","inputs":[{"name":"_duration","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setConfigurator","inputs":[{"name":"newConfigurator","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setDelegatedRecipient","inputs":[{"name":"_recipient","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setDesignatedVerifierURI","inputs":[{"name":"designatedVerifierURI","type":"string","internalType":"string"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setHookAsDebtorForSBP","inputs":[{"name":"_hook","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setIndexUpdateBeacon","inputs":[{"name":"_beacon","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setInterestRateModel","inputs":[{"name":"_type","type":"uint8","internalType":"enum IK2Lending.DebtPositionType"},{"name":"_newInterestRateModel","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setKETHAddress","inputs":[{"name":"_kETH","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setMinParams","inputs":[{"name":"newMinDepositLimit","type":"uint256","internalType":"uint256"},{"name":"newMinTransferLimit","type":"uint256","internalType":"uint256"},{"name":"newMinLockUpPeriod","type":"uint256","internalType":"uint256"},{"name":"_maxSlashableRatio_RAY","type":"uint256","internalType":"uint256"},{"name":"_terminatePenalty_RAY","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setNodeOperatorInclusionList","inputs":[{"name":"_nodeOperatorInclusionList","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setRSTModule","inputs":[{"name":"_rstModule","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"slash","inputs":[{"name":"slashType","type":"uint8","internalType":"enum ReporterRegistry.SlashType"},{"name":"debtor","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"recipient","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"supportsInterface","inputs":[{"name":"interfaceId","type":"bytes4","internalType":"bytes4"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"symbol","inputs":[],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"terminate","inputs":[{"name":"debtor","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"terminatePenalty_RAY","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"topUpAndTerminate","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"topUpSlashAmount","inputs":[{"name":"debtor","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"totalSupply","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"transfer","inputs":[{"name":"to","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"transferFrom","inputs":[{"name":"from","type":"address","internalType":"address"},{"name":"to","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"updateInterest","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"updateLenderPosition","inputs":[{"name":"lender","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"updateNodeOperatorPosition","inputs":[{"name":"nodeOperator","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"withdraw","inputs":[{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"claim","type":"bool","internalType":"bool"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"event","name":"Approval","inputs":[{"name":"owner","type":"address","indexed":true,"internalType":"address"},{"name":"spender","type":"address","indexed":true,"internalType":"address"},{"name":"value","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"Borrowed","inputs":[{"name":"borrower","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"designatedVerifier","type":"address","indexed":true,"internalType":"address"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"interestPaid","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"DelegatedClaimRecipientSet","inputs":[{"name":"claimRecipient","type":"address","indexed":true,"internalType":"address"},{"name":"delegateClaimer","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"DesignatedVerifierURIUpdated","inputs":[{"name":"borrower","type":"address","indexed":true,"internalType":"address"},{"name":"designatedVerifierURI","type":"string","indexed":false,"internalType":"string"}],"anonymous":false},{"type":"event","name":"IncreasedDebt","inputs":[{"name":"borrower","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"designatedVerifier","type":"address","indexed":true,"internalType":"address"},{"name":"maxSlashableAmountPerLiveness","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"maxSlashableAmountPerCorruption","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"resetDuration","type":"bool","indexed":false,"internalType":"bool"},{"name":"interestPaid","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"Initialized","inputs":[{"name":"version","type":"uint8","indexed":false,"internalType":"uint8"}],"anonymous":false},{"type":"event","name":"KETHClaimed","inputs":[{"name":"depositor","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"KETHDeposited","inputs":[{"name":"depositor","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"KETHWithdrawn","inputs":[{"name":"depositor","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"Liquidated","inputs":[{"name":"borrower","type":"address","indexed":true,"internalType":"address"},{"name":"liquidator","type":"address","indexed":true,"internalType":"address"},{"name":"liquidationAmount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"NodeOperatorClaimed","inputs":[{"name":"nodeOperator","type":"address","indexed":true,"internalType":"address"},{"name":"payoutRecipient","type":"address","indexed":true,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"NodeOperatorDeposited","inputs":[{"name":"operator","type":"address","indexed":true,"internalType":"address"},{"name":"blsPublicKey","type":"bytes","indexed":false,"internalType":"bytes"},{"name":"payoutRecipient","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"NodeOperatorWithdrawn","inputs":[{"name":"operator","type":"address","indexed":true,"internalType":"address"},{"name":"blsPublicKey","type":"bytes","indexed":false,"internalType":"bytes"},{"name":"kicked","type":"bool","indexed":false,"internalType":"bool"}],"anonymous":false},{"type":"event","name":"RepaidSlashAmount","inputs":[{"name":"borrower","type":"address","indexed":true,"internalType":"address"},{"name":"topupAmount","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"RoleAdminChanged","inputs":[{"name":"role","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"previousAdminRole","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"newAdminRole","type":"bytes32","indexed":true,"internalType":"bytes32"}],"anonymous":false},{"type":"event","name":"RoleGranted","inputs":[{"name":"role","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"account","type":"address","indexed":true,"internalType":"address"},{"name":"sender","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"RoleRevoked","inputs":[{"name":"role","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"account","type":"address","indexed":true,"internalType":"address"},{"name":"sender","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"SBPHookUpdated","inputs":[{"name":"borrower","type":"address","indexed":true,"internalType":"address"},{"name":"hook","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"Slashed","inputs":[{"name":"debtor","type":"address","indexed":false,"internalType":"address"},{"name":"amount","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"recipient","type":"address","indexed":false,"internalType":"address"}],"anonymous":false},{"type":"event","name":"Terminated","inputs":[{"name":"debtor","type":"address","indexed":false,"internalType":"address"}],"anonymous":false},{"type":"event","name":"Transfer","inputs":[{"name":"from","type":"address","indexed":true,"internalType":"address"},{"name":"to","type":"address","indexed":true,"internalType":"address"},{"name":"value","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"error","name":"ComeBackLater","inputs":[]},{"type":"error","name":"ExceedMaxBorrowRatio","inputs":[]},{"type":"error","name":"ExceedMaxSlashableAmountPerCorruption","inputs":[]},{"type":"error","name":"ExceedMaxSlashableAmountPerLiveness","inputs":[]},{"type":"error","name":"HasDebt","inputs":[]},{"type":"error","name":"InvalidDebtPositionType","inputs":[]},{"type":"error","name":"InvalidLength","inputs":[]},{"type":"error","name":"MinDepositLimit","inputs":[]},{"type":"error","name":"MustPaySlashedAmount","inputs":[]},{"type":"error","name":"NoDebt","inputs":[]},{"type":"error","name":"NoElements","inputs":[]},{"type":"error","name":"NoSlashedAmount","inputs":[]},{"type":"error","name":"NoTerminationWithRST","inputs":[]},{"type":"error","name":"NodeOperatorAlreadyRegistered","inputs":[]},{"type":"error","name":"NodeOperatorBLSKeyNotPermitted","inputs":[]},{"type":"error","name":"NodeOperatorInvalid","inputs":[]},{"type":"error","name":"NodeOperatorInvalidRepresentative","inputs":[]},{"type":"error","name":"NodeOperatorInvalidStatus","inputs":[]},{"type":"error","name":"NodeOperatorKicked","inputs":[]},{"type":"error","name":"NodeOperatorNotRegistered","inputs":[]},{"type":"error","name":"NotAbleToLiquidate","inputs":[]},{"type":"error","name":"NotAllowedToDecreaseInterestRate","inputs":[]},{"type":"error","name":"NotEnoughAssumedLiquidity","inputs":[]},{"type":"error","name":"NotEnoughOutstandingInterestToSlash","inputs":[]},{"type":"error","name":"TooSmall","inputs":[]},{"type":"error","name":"Unauthorized","inputs":[]},{"type":"error","name":"ZeroAddress","inputs":[]}]`
	K2_NODE_OPERATOR_CONTRACT_ABI = `[{"type":"constructor","inputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"E_BALANCE_REPORT_TYPEHASH","inputs":[],"outputs":[{"name":"","type":"bytes32","internalType":"bytes32"}],"stateMutability":"view"},{"type":"function","name":"MAX_NATIVE_DELEGATION_PER_NODE_OPERATOR","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"MAX_OPEN_NATIVE_DELEGATION_CAPACITY","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"batchNodeOperatorKick","inputs":[{"name":"_blsPubkeys","type":"bytes[]","internalType":"bytes[]"},{"name":"_effectiveBalances","type":"uint256[]","internalType":"uint256[]"},{"name":"_designatedVerifierSignatures","type":"tuple[]","internalType":"struct EIP712Verifier.SignatureECDSA[]","components":[{"name":"v","type":"uint8","internalType":"uint8"},{"name":"r","type":"bytes32","internalType":"bytes32"},{"name":"s","type":"bytes32","internalType":"bytes32"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"delegatedClaim","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"eip712Domain","inputs":[],"outputs":[{"name":"fields","type":"bytes1","internalType":"bytes1"},{"name":"name","type":"string","internalType":"string"},{"name":"version","type":"string","internalType":"string"},{"name":"chainId","type":"uint256","internalType":"uint256"},{"name":"verifyingContract","type":"address","internalType":"address"},{"name":"salt","type":"bytes32","internalType":"bytes32"},{"name":"extensions","type":"uint256[]","internalType":"uint256[]"}],"stateMutability":"view"},{"type":"function","name":"getDomainSeparator","inputs":[],"outputs":[{"name":"","type":"bytes32","internalType":"bytes32"}],"stateMutability":"view"},{"type":"function","name":"incentivesHook","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"init","inputs":[{"name":"_initialOwner","type":"address","internalType":"address"},{"name":"_maxNumOfKeysThatCanRegister","type":"uint256","internalType":"uint256"},{"name":"_maxNumOfKeysThatCanRegisterPerMember","type":"uint256","internalType":"uint256"},{"name":"_proposerRegistry","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"isNewBLSKeyPermitted","inputs":[{"name":"_blsPubKey","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isPartOfInclusionList","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isValidEffectiveBalanceReport","inputs":[{"name":"_blsPubkey","type":"bytes","internalType":"bytes"},{"name":"_effectiveBalance","type":"uint256","internalType":"uint256"},{"name":"_designatedVerifierSignature","type":"tuple","internalType":"struct EIP712Verifier.SignatureECDSA","components":[{"name":"v","type":"uint8","internalType":"uint8"},{"name":"r","type":"bytes32","internalType":"bytes32"},{"name":"s","type":"bytes32","internalType":"bytes32"}]}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"lending","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"manageDesignatedVerifier","inputs":[{"name":"_verifier","type":"address","internalType":"address"},{"name":"_isEnabled","type":"bool","internalType":"bool"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"nodeOperatorClaim","inputs":[{"name":"_recipient","type":"address","internalType":"address"},{"name":"_blsPubKeys","type":"bytes[]","internalType":"bytes[]"},{"name":"_effectiveBalances","type":"uint256[]","internalType":"uint256[]"},{"name":"_designatedVerifierSignatures","type":"tuple[]","internalType":"struct EIP712Verifier.SignatureECDSA[]","components":[{"name":"v","type":"uint8","internalType":"uint8"},{"name":"r","type":"bytes32","internalType":"bytes32"},{"name":"s","type":"bytes32","internalType":"bytes32"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"nodeOperatorKick","inputs":[{"name":"_blsPubkey","type":"bytes","internalType":"bytes"},{"name":"_effectiveBalance","type":"uint256","internalType":"uint256"},{"name":"_designatedVerifierSignature","type":"tuple","internalType":"struct EIP712Verifier.SignatureECDSA","components":[{"name":"v","type":"uint8","internalType":"uint8"},{"name":"r","type":"bytes32","internalType":"bytes32"},{"name":"s","type":"bytes32","internalType":"bytes32"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"nodeOperatorWithdraw","inputs":[{"name":"_recipient","type":"address","internalType":"address"},{"name":"_blsPubKey","type":"bytes","internalType":"bytes"},{"name":"_effectiveBalance","type":"uint256","internalType":"uint256"},{"name":"_designatedVerifierSignature","type":"tuple","internalType":"struct EIP712Verifier.SignatureECDSA","components":[{"name":"v","type":"uint8","internalType":"uint8"},{"name":"r","type":"bytes32","internalType":"bytes32"},{"name":"s","type":"bytes32","internalType":"bytes32"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"onBLSKeyRegisteredToK2","inputs":[{"name":"_blsPubKey","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"onBLSKeyUnregisteredToK2","inputs":[{"name":"_blsPubKey","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"onSBPUpdated","inputs":[{"name":"debtor","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"onWithdraw","inputs":[{"name":"user","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"view"},{"type":"function","name":"owner","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"pnoRegistry","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"proposerRegistry","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"proxiableUUID","inputs":[],"outputs":[{"name":"","type":"bytes32","internalType":"bytes32"}],"stateMutability":"view"},{"type":"function","name":"renounceOwnership","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"reportExitFromProposerRegistry","inputs":[{"name":"_blsPubKey","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"reportTypedHash","inputs":[{"name":"_blsPubkey","type":"bytes","internalType":"bytes"},{"name":"_effectiveBalance","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"typedHash","type":"bytes32","internalType":"bytes32"}],"stateMutability":"view"},{"type":"function","name":"reporterRegistry","inputs":[],"outputs":[{"name":"","type":"address","internalType":"contract ReporterRegistry"}],"stateMutability":"view"},{"type":"function","name":"setDelegatedRecipient","inputs":[{"name":"_recipient","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setIncentivesHook","inputs":[{"name":"_hook","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setLending","inputs":[{"name":"_k2Lending","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setPNORegistry","inputs":[{"name":"_pnoRegistry","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setReporterRegistry","inputs":[{"name":"_reporterRegistry","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"totalNumberOfRegisteredKeysForInclusionListMember","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"totalOpenNativeDelegationCapacityConsumed","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"transferOwnership","inputs":[{"name":"newOwner","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"updateInclusionList","inputs":[{"name":"_targets","type":"address[]","internalType":"address[]"},{"name":"_enabled","type":"bool[]","internalType":"bool[]"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"updateMaxNativeDelegationPerNodeOperator","inputs":[{"name":"_maxNumOfKeysThatCanRegister","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"updateOpenNativeDelegationCapacity","inputs":[{"name":"_maxOpenNativeDelegationCapacity","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"upgradeTo","inputs":[{"name":"newImplementation","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"upgradeToAndCall","inputs":[{"name":"newImplementation","type":"address","internalType":"address"},{"name":"data","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"payable"},{"type":"event","name":"AdminChanged","inputs":[{"name":"previousAdmin","type":"address","indexed":false,"internalType":"address"},{"name":"newAdmin","type":"address","indexed":false,"internalType":"address"}],"anonymous":false},{"type":"event","name":"BeaconUpgraded","inputs":[{"name":"beacon","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"DelegatedClaimRecipientSet","inputs":[{"name":"claimRecipient","type":"address","indexed":true,"internalType":"address"},{"name":"delegateClaimer","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"EIP712DomainChanged","inputs":[],"anonymous":false},{"type":"event","name":"Initialized","inputs":[{"name":"version","type":"uint8","indexed":false,"internalType":"uint8"}],"anonymous":false},{"type":"event","name":"MaxNativeDelegationPerNodeOperatorUpdated","inputs":[{"name":"newValue","type":"uint256","indexed":true,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"MaxNumOfOpenNativeDelegationsUpdated","inputs":[{"name":"newValue","type":"uint256","indexed":true,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"OwnershipTransferred","inputs":[{"name":"previousOwner","type":"address","indexed":true,"internalType":"address"},{"name":"newOwner","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"Upgraded","inputs":[{"name":"implementation","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"isPartOfInclusionListUpdated","inputs":[{"name":"target","type":"address","indexed":true,"internalType":"address"},{"name":"enabled","type":"bool","indexed":true,"internalType":"bool"}],"anonymous":false}]`
	PNO_REGISTRY_CONTRACT_ABI     = `[{"type":"constructor","inputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"LIQUID_DELEGATION_TEMPORAL_RULE","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"PNO_NATIVE_DELEGATION_THRESHOLD","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"banNodeOperator","inputs":[{"name":"_sbpIndexId","type":"uint256","internalType":"uint256"},{"name":"_pno","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"createKNetwork","inputs":[{"name":"defaultRewardToken","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"currentSBPIndexForKNetwork","inputs":[{"name":"kNetworkDebtor","type":"address","internalType":"address"}],"outputs":[{"name":"sbpIndexId","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"init","inputs":[{"name":"_nodeOperatorModule","type":"address","internalType":"address"},{"name":"_initialOwner","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"isInKNetworkInclusionList","inputs":[{"name":"pno","type":"address","internalType":"address"},{"name":"sbpIndexId","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"isInInclusionList","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isKNetworkActive","inputs":[{"name":"sbpIndexId","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"isReady","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isKNetworkGateKeepingEnabled","inputs":[{"name":"sbpIndexId","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"isRestricted","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isPreferredNodeOperator","inputs":[{"name":"_pno","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isPreferredNodeOperatorForKNetwork","inputs":[{"name":"pno","type":"address","internalType":"address"},{"name":"sbpIndexId","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"isActive","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isProtocolPNOManagementEnabled","inputs":[],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"kNetworkFundSplitterFactory","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"lastLiquidDelegationTimestamp","inputs":[{"name":"user","type":"address","internalType":"address"}],"outputs":[{"name":"timestamp","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"lienContract","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"manageKNetworkInclusionList","inputs":[{"name":"_PNOs","type":"address[]","internalType":"address[]"},{"name":"_enabled","type":"bool","internalType":"bool"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"managePNO","inputs":[{"name":"_pno","type":"address","internalType":"address"},{"name":"_enabled","type":"bool","internalType":"bool"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"nodeOperatorModule","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"onDelegatedRecipientSet","inputs":[{"name":"","type":"address","internalType":"address"},{"name":"currentRecipient","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"view"},{"type":"function","name":"onNodeOperatorWithdraw","inputs":[{"name":"user","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"view"},{"type":"function","name":"onSBPUpdated","inputs":[{"name":"debtor","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"onWithdraw","inputs":[{"name":"user","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"view"},{"type":"function","name":"optIntoKNetwork","inputs":[{"name":"_sbpIndexId","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"optOutOfKNetwork","inputs":[{"name":"_sbpIndexId","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"owner","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"pnoFeeRecipient","inputs":[{"name":"pno","type":"address","internalType":"address"}],"outputs":[{"name":"feeRecipient","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"proxiableUUID","inputs":[],"outputs":[{"name":"","type":"bytes32","internalType":"bytes32"}],"stateMutability":"view"},{"type":"function","name":"reDelegateToPreferredNodeOperator","inputs":[{"name":"_pno","type":"address","internalType":"address"},{"name":"_amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"reDelegateToPreferredNodeOperatorFromNativeDelegation","inputs":[{"name":"_pno","type":"address","internalType":"address"},{"name":"_amountToRedelegateInETH","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"removePNO","inputs":[{"name":"_pno","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"removeRedelegation","inputs":[{"name":"_pno","type":"address","internalType":"address"},{"name":"_amount","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"removeRedelegationFromNativeDelegation","inputs":[{"name":"_pno","type":"address","internalType":"address"},{"name":"_amountToRedelegateInETH","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"renounceOwnership","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setIsProtocolPNOManagementEnabled","inputs":[{"name":"_isEnabled","type":"bool","internalType":"bool"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setKNetworkFundSplitterFactory","inputs":[{"name":"_factory","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setLienContract","inputs":[{"name":"_lienContract","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setNativeDelegationThreshold","inputs":[{"name":"_threshold","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setPNOFeeRecipient","inputs":[{"name":"feeRecipient","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setTemporalRule","inputs":[{"name":"_rule","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"togglePNOGatekeepingForKNetwork","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"totalNumberOfPNOsPartOfKNetwork","inputs":[{"name":"sbpIndexId","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"total","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"totalReDelegationsInETH","inputs":[{"name":"user","type":"address","internalType":"address"}],"outputs":[{"name":"total","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"totalReDelegationsInETHFromK2LP","inputs":[{"name":"user","type":"address","internalType":"address"}],"outputs":[{"name":"totalFromK2LpOnly","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"totalReDelegationsInETHReceivedByPNO","inputs":[{"name":"pno","type":"address","internalType":"address"}],"outputs":[{"name":"total","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"totalReDelegationsInETHToPNO","inputs":[{"name":"_user","type":"address","internalType":"address"},{"name":"_pno","type":"address","internalType":"address"},{"name":"_timestamp","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"totalReDelegationsInETHToPNO","inputs":[{"name":"_user","type":"address","internalType":"address"},{"name":"_pno","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"transferOwnership","inputs":[{"name":"newOwner","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"upgradeTo","inputs":[{"name":"newImplementation","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"upgradeToAndCall","inputs":[{"name":"newImplementation","type":"address","internalType":"address"},{"name":"data","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"payable"},{"type":"event","name":"AdminChanged","inputs":[{"name":"previousAdmin","type":"address","indexed":false,"internalType":"address"},{"name":"newAdmin","type":"address","indexed":false,"internalType":"address"}],"anonymous":false},{"type":"event","name":"BeaconUpgraded","inputs":[{"name":"beacon","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"Initialized","inputs":[{"name":"version","type":"uint8","indexed":false,"internalType":"uint8"}],"anonymous":false},{"type":"event","name":"OwnershipTransferred","inputs":[{"name":"previousOwner","type":"address","indexed":true,"internalType":"address"},{"name":"newOwner","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"PNOForSBPManaged","inputs":[{"name":"preferredNodeOperator","type":"address","indexed":true,"internalType":"address"},{"name":"sbpIndex","type":"uint256","indexed":true,"internalType":"uint256"},{"name":"enabled","type":"bool","indexed":false,"internalType":"bool"}],"anonymous":false},{"type":"event","name":"PNOManaged","inputs":[{"name":"preferredNodeOperator","type":"address","indexed":true,"internalType":"address"},{"name":"enabled","type":"bool","indexed":false,"internalType":"bool"}],"anonymous":false},{"type":"event","name":"PNONativeDelegationThresholdUpdated","inputs":[],"anonymous":false},{"type":"event","name":"PNORedelegationReceived","inputs":[{"name":"preferredNodeOperator","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"PNORedelegationRemoved","inputs":[{"name":"preferredNodeOperator","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"Upgraded","inputs":[{"name":"implementation","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"kNetworkCreated","inputs":[{"name":"kNetworkId","type":"uint256","indexed":true,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"kNetworkGatekeepingToggled","inputs":[{"name":"kNetworkId","type":"uint256","indexed":true,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"kNetworkInclusionListUpdated","inputs":[{"name":"pnos","type":"address[]","indexed":false,"internalType":"address[]"},{"name":"enabled","type":"bool","indexed":false,"internalType":"bool"}],"anonymous":false}]`
)