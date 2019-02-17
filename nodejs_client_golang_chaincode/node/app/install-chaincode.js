'use strict';
var path = require('path');
var fs = require('fs');
var util = require('util');
var helper = require('./helper.js');
var logger = helper.getLogger('install-chaincode');

var tx_id = null;
var ORGS = helper.getORGS();

/**
 * Install the chaincode to the target peers
 * @param {*Target peers} peers 
 * @param {*Target chaincode name} chaincodeName 
 * @param {*Target chaincode path} chaincodePath 
 * @param {*Target chaincode version} chaincodeVersion 
 * @param {*Target channel name} channel_name 
 * @param {*Target org, including the target peers} org_name 
 */
var installChaincode = function (peers, chaincodeName, chaincodePath,
	chaincodeVersion, channel_name, org_name) {
	logger.info('\n\n============ Install chaincode on organizations ============\n');
	helper.setupChaincodeDeploy();

	var client = null;
	var channel = null;

	org_name =  helper.checkOrg(org_name);
	peers =  helper.checkPeers(peers,org_name);

	return helper.getClientForOrg(org_name).then(_client => {
		client = _client;
		channel = client.newChannel(channel_name);
		var targets = [];
		helper.setTargetPeers(client, channel, targets, org_name,peers);
		// send proposal to endorser
		var request = {
			targets: targets,
			chaincodePath: chaincodePath,
			chaincodeId: chaincodeName,
			chaincodeVersion: chaincodeVersion,
			chaincodeType: "golang"		
		};

		return client.installChaincode(request);
	}, (err) => {
		throw new Error('Failed to create user \'admin\'. ' + err);
	}).then((results) => {
		var proposalResponses = results[0];
		var proposal = results[1];
		var all_good = true;
		var errors = [];
		var isExist = 0;
		for (var i in proposalResponses) {
			let one_good = false;
			if (proposalResponses && proposalResponses[i].response && proposalResponses[i].response.status === 200) {
				one_good = true;
				logger.info('install proposal was good');
			} else {
				if(proposalResponses[i].details.indexOf("exists") != -1){
					logger.info("Chaincode is exists. Continue...");
					isExist++;
				}
				else {
					logger.error('install proposal was bad');
					errors.push(proposalResponses[i]);
				}
			}
			all_good = all_good & one_good;
		}
		if(isExist == proposalResponses.length) return { success: true };
		if (all_good) {
			logger.info(util.format('Successfully sent install Proposal and received ProposalResponse: Status - %s', proposalResponses[0].response.status));
			return { success: true };
		} else {
			throw new Error(util.format('Failed to send install Proposal or receive valid response: %s', errors));
		}
	},
		(err) => {
			logger.error('Failed to send install proposal due to error: ' + err.stack ? err.stack : err);
			throw new Error('Failed to send install proposal due to error: ' + err.stack ? err.stack : err);
		}).catch(err => {
			logger.error('Failed to send install proposal due to error: ' + err.stack ? err.stack : err);
			throw new Error('Failed to send install proposal due to error: ' + err.stack ? err.stack : err);
		});
};

exports.installChaincode = installChaincode;
