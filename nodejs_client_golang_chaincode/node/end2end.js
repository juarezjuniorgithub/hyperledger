var log4js = require('log4js');
var logger = log4js.getLogger('end2end');
logger.setLevel('DEBUG');
require('./config.js');//Load the config info

var install = require('./app/install-chaincode.js');
var instantiate = require('./app/instantiate-chaincode.js');
var invoke = require('./app/invoke-transaction.js');
var query = require('./app/query.js');

//Convert string to a array for chaincode function call
function array(val) {
    return val.split(',');
}

//Get the error message
function getErrorMessage(field) {
    var response = {
        success: false,
        message: field + ' field is missing or Invalid in the request'
    };
    return response;
}

//Abstract commandline arguments  
var args = process.argv.splice(2);

//Get the call method 
method_choose = args[0];
//Declear the var
var params_input_json = JSON.parse(args[1]);


function param_install_check(param_json){
    peers = param_json.peers;
    chaincodeName = param_json.chaincodeId;
    chaincodePath = param_json.chaincodePath;
    chaincodeVersion = param_json.chaincodeVersion;
    channelName = param_json.channelName;
    orgName = param_json.org;

    logger.info('peers : ' + peers); // target peers list
    logger.info('chaincodeName : ' + chaincodeName);
    logger.info('chaincodePath  : ' + chaincodePath);
    logger.info('chaincodeVersion  : ' + chaincodeVersion);
    logger.info('channelName  : ' + channelName);
    logger.info('orgName  : ' + orgName);

    if (!peers || peers.length == 0) {
        console.error(getErrorMessage('\'peers\''));
        return false;
    }
    if (!chaincodeName) {
        console.error(getErrorMessage('\'chaincodeName\''));
        return false;
    }
    if (!chaincodePath) {
        console.error(getErrorMessage('\'chaincodePath\''));
        return false;
    }
    if (!chaincodeVersion) {
        console.error(getErrorMessage('\'chaincodeVersion\''));
        return false;
    }
    if (!channelName) {
        console.error(getErrorMessage('\'channelName\''));
        return false;
    }
    return true;
}

function param_instantiate_check(param_json){
    peers = param_json.peers;
    chaincodeName = param_json.chaincodeId;
    chaincodeVersion = param_json.chaincodeVersion;
    channelName = param_json.channelName;
    chaincodePath = param_json.chaincodePath;
    cfcn = param_json.cfcn;
    cargs = param_json.cargs;
    orgName = param_json.org;

    logger.debug('peers  : ' + peers);
    logger.debug('channelName  : ' + channelName);
    logger.debug('chaincodeName : ' + chaincodeName);
    logger.debug('chaincodeVersion  : ' + chaincodeVersion);
    logger.debug('chaincodePath  : ' + chaincodePath);
    logger.debug('fcn  : ' + cfcn);
    logger.debug('args  : ' + cargs);
    logger.info('orgName  : ' + orgName);

    if (!peers || peers.length == 0) {
        console.error(getErrorMessage('\'peers\''));
        return false;
    }
    if (!chaincodeName) {
        console.error(getErrorMessage('\'chaincodeName\''));
        return false;
    }
    if (!chaincodeVersion) {
        console.error(getErrorMessage('\'chaincodeVersion\''));
        return false;
    }
    if (!channelName) {
        console.error(getErrorMessage('\'channelName\''));
        return false;
    }
    if (!chaincodePath) {
        console.error(getErrorMessage('\'chaincodePath\''));
        return false;
    }
    if (!cargs) {
        console.error(getErrorMessage('\'args\''));
        return false;
    }
    if (!cfcn) {
        console.error(getErrorMessage('\'fcn\''));
        return false;
    }
    if (!orgName) {
        console.error(getErrorMessage('\'orgName\''));
        return false;
    }
    return true;
}

function param_invoke_check(param_json){
    peers = param_json.peers;
    chaincodeName = param_json.chaincodeId;
    channelName = param_json.channelName;
    cfcn = param_json.cfcn;
    cargs = param_json.cargs;
    orgName = param_json.org;

    logger.debug('peers  : ' + peers);
    logger.debug('channelName  : ' + channelName);
    logger.debug('chaincodeName : ' + chaincodeName);
    logger.debug('fcn  : ' + cfcn);
    logger.debug('args  : ' + cargs);
    logger.info('orgName  : ' + orgName);

    if (!peers || peers.length == 0) {
        console.error(getErrorMessage('\'peers\''));
        return false;
    }
    if (!chaincodeName) {
        console.error(getErrorMessage('\'chaincodeName\''));
        return false;
    }
    if (!channelName) {
        console.error(getErrorMessage('\'channelName\''));
        return false;
    }
    if (!cfcn) {
        console.error(getErrorMessage('\'fcn\''));
        return false;
    }
    if (!cargs) {
        console.error(getErrorMessage('\'args\''));
        return false;
    }
    if (!orgName) {
        console.error(getErrorMessage('\'orgName\''));
        return false;
    }
    return true;
}

function param_query_check(param_json){
    channelName = param_json.channelName;
    chaincodeName = param_json.chaincodeId;
    cargs = param_json.cargs;
    cfcn = param_json.cfcn;
    peers = param_json.peers;
    orgName = param_json.org;

    logger.debug('peers  : ' + peers);
    logger.debug('channelName : ' + channelName);
    logger.debug('chaincodeName : ' + chaincodeName);
    logger.debug('fcn : ' + cfcn);
    logger.debug('args : ' + cargs);
    logger.info('orgName  : ' + orgName);

    if (!peers || peers.length == 0) {
        console.error(getErrorMessage('\'peers\''));
        return false;
    }
    if (!chaincodeName) {
        console.error(getErrorMessage('\'chaincodeName\''));
        return false;
    }
    if (!channelName) {
        console.error(getErrorMessage('\'channelName\''));
        return false;
    }
    if (!cfcn) {
        console.error(getErrorMessage('\'fcn\''));
        return false;
    }
    if (!cargs) {
        console.error(getErrorMessage('\'args\''));
        return false;
    }
    if (!orgName) {
        console.error(getErrorMessage('\'orgName\''));
        return false;
    }
    return true;
}

switch (method_choose) {
    case "install":
        //Read the param
        logger.info('==================== INSTALL CHAINCODE ==================');
        if(!param_install_check(params_input_json)){
            console.log("Params are invalid, Please check and run the test again")
            return ;
        }
        
        //Call the function
        return install.installChaincode(array(params_input_json.peers), params_input_json.chaincodeId, params_input_json.chaincodePath, 
            params_input_json.chaincodeVersion, params_input_json.channelName, params_input_json.org).then(result => {
            console.log(result);
        }, err => {
            console.error(err);
        }).catch(err => { console.error(err); });
        break;

    case "instantiate":
        logger.debug('==================== INSTANTIATE CHAINCODE ==================');
        if(!param_instantiate_check(params_input_json)){
            console.log("Params are invalid, Please check and run the test again")
            return ;            
        }
        
        //Call the function
        instantiate.instantiateChaincode(array(params_input_json.peers), params_input_json.channelName, params_input_json.chaincodeId,
            params_input_json.chaincodeVersion, params_input_json.cfcn, params_input_json.chaincodePath, array(params_input_json.cargs), params_input_json.org)
            .then((message) => {
                console.log(message);
                process.exit();
            }, err => {
                console.error(err);
                process.exit();
            }).catch(err => { console.error(err); });
        break;

    case "invoke":
        logger.debug('==================== INVOKE ON CHAINCODE ==================');
        if(!param_invoke_check(params_input_json)){
            console.log("Params are invalid, Please check and run the test again")
            return ;            
        }

        //Call the function
        invoke.invokeChaincode(array(params_input_json.peers), params_input_json.channelName, params_input_json.chaincodeId, 
            params_input_json.cfcn, array(params_input_json.cargs), params_input_json.org)
            .then((message) => {
                console.log(message);
                process.exit();
            }, err => {
                console.error(err);
                process.exit();
            }).catch(err => { console.error(err); });
        break;

    case "query":
        logger.debug('==================== QUERY BY CHAINCODE ==================');
        if(!param_query_check(params_input_json)){
            console.log("Params are invalid, Please check and run the test again")
            return ;            
        }

        //Call the function
        query.queryChaincode(array(params_input_json.peers), params_input_json.channelName, params_input_json.chaincodeId, 
        array(params_input_json.cargs), params_input_json.cfcn, params_input_json.org)
            .then((message) => {
                console.log(message);
            }, err => {
                console.error(err);
            }).catch(err => { console.error(err); });
        break;
    default:
        console.log("No target method support");
}

return true;




