const CronosToken = artifacts.require("CronosToken");

module.exports = function (deployer) {
    deployer.deploy(
        CronosToken,
        "GasLess Token",
        "GLT",
        "1000000000000000000000000"
    );
};
