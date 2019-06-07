

const GovernedProxy = artifacts.require('GovernedProxy');
const GovernedProxyTest = artifacts.require('GovernedProxyTest');

contract("GovernedProxy", async accounts => {
    let first;
    let second;
    let proxy;
    let proxy_abi;
    const weeks = 60*60*24*7;

    before(async () => {
        first = await GovernedProxyTest.new();
        second = await GovernedProxyTest.new();
        proxy = await GovernedProxy.new(first.address);
        proxy_abi = await GovernedProxyTest.at(proxy.address);
    });

    it('should proxy', async () => {
        const res = await proxy_abi.getAddress.call({ from: accounts[0] });
        assert.equal(first.address.valueOf(), res.valueOf());
    });

    it('should refuse proposal - same impl', async () => {
        try {
            await proxy.proposeUpgrade.call(
                    first.address, 2 * weeks,
                    { from: accounts[0], value: web3.utils.toWei('10000', 'ether') });
            assert.fail("It must fail");
        } catch (e) {
            assert.equal(e.message, "Returned error: VM Exception while processing transaction: revert Already active!");
        }
    });

    // TODO: full test suite
});
