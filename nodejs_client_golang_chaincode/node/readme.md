## NODE.js CLIENT SDK - Example of client Node.js App

Here is a sample Node.JS client application that utilizes the Hyperledger Fabric NODE JS SDK to 

* Connect to the Oracle Blockchain Platform network using a set of config files
* Connect to a channel
* Installs chaincode written in the "go" programming language
* Instantiate chaincode on a set of specific peers on a specific channel
* Invoke chaincode

It demonstrates how you could utilize both the **__fabric-client__** & **__fabric-ca-client__** Node.js SDK APIs.

The "config.json" file located in the parent directory mirrors your existing Oracle Blockchain Platform environment. Namely it describes

* An ordering service
* An organization
* One peer (1 peers per Org)

It also describes where the security certificates with which to connect with your environment are located.

### Step 1: Install prerequisites

* **Node.js** v 6x

### Step 2: Initialize the sample application

We need to use the "npm" package manager to initialize the application. To do this run the following command in your terminal from the parent directory: `npm install`

### Step 3: Modify configuration files

In the parent directory "testAPI.sh", change the `CHANNEL_NAME` to the channel you wish to utilize to run the sample. The default channel is provided as 'default'.

### Step 4: Run the sample application

To run the application, execute the following shell command: `./test-APIs.sh` or `bash test-APIs.sh`

"All Done"
