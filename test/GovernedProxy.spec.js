

const GovernedProxy = artifacts.require('GovernedProxy');
const SporkRegistryV1 = artifacts.require('SporkRegistryV1');
const MockContract = artifacts.require('MockContract');

contract("GovernedProxy", async accounts => {
    let first;
    let second;
    let proxy;
    let proxy_abi;
    let registry;
    const weeks = 60*60*24*7;

    before(async () => {
        registry = await SporkRegistryV1.deployed();
        first = await MockContract.new(registry.address);
        proxy = await GovernedProxy.new(first.address, registry.address);
        second = await MockContract.new(proxy.address);
        proxy_abi = await MockContract.at(proxy.address);
    });

    it('should proxy', async () => {
        const res = await proxy_abi.getAddress({ from: accounts[0] });
        assert.equal(first.address.valueOf(), res.valueOf());
    });

    it('should accept proposal', function(){ this.skip(); });
    it('should accept upgrade', function(){ this.skip(); });

    it('should refuse proposal - same impl', async () => {
        try {
            await proxy.proposeUpgrade(
                    first.address, 2 * weeks,
                    { from: accounts[0], value: web3.utils.toWei('10000', 'ether') });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Already active!/);
        }
    });

    it('should refuse proposal - wrong proxy', async () => {
        try {
            await proxy.proposeUpgrade(
                    registry.address, 2 * weeks,
                    { from: accounts[0], value: web3.utils.toWei('10000', 'ether') });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Wrong proxy!/);
        }
    });

    it('should refuse migrate()', async () => {
        try {
            await proxy.migrate(second.address, { from: accounts[0] });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Good try/);
        }
    });

    it('should refuse destroy()', async () => {
        try {
            await proxy.destroy(second.address, { from: accounts[0] });
            assert.fail("It must fail");
        } catch (e) {
            assert.match(e.message, /Good try/);
        }
    });

    it('should refuse proposal - wrong fee', function(){ this.skip(); });
    it('should refuse upgrade - Already active!', function(){ this.skip(); });
    it('should refuse upgrade - Not accepted!', function(){ this.skip(); });
    it('should refuse upgrade - Not registered!', function(){ this.skip(); });
    it('should refuse upgrade AFTER upgrade - Not registered!', function(){ this.skip(); });
});
