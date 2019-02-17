var util = require('util');
var path = require('path');
var hfc = require('fabric-client');

// some other settings the application might need to know
hfc.addConfigFile(path.join(__dirname, '../config.json'));
