const { expect } = require("chai");

describe("LiuToken Contract", function () {
    let Token, token, owner, addr1, addr2;

    beforeEach(async function () {
        Token = await ethers.getContractFactory("LiuToken");
        [owner, addr1, addr2] = await ethers.getSigners();

        token = await Token.deploy("LiuToken", "LIU", 18, 1000000); // 初始供应量为 1,000,000
        await token.deployed();
    });

    it("Should assign the total supply of tokens to the owner", async function () {
        const ownerBalance = await token.balanceOf(owner.address);
        expect(await token.totalSupply()).to.equal(ownerBalance);
    });

    it("Should transfer tokens between accounts", async function () {
        await token.transfer(addr1.address, 50);
        const addr1Balance = await token.balanceOf(addr1.address);
        expect(addr1Balance).to.equal(50);
    });

    it("Should fail if sender doesn't have enough tokens", async function () {
        const initialOwnerBalance = await token.balanceOf(owner.address);

        await expect(
            token.connect(addr1).transfer(owner.address, 1)
        ).to.be.revertedWith("Insufficient balance");

        expect(await token.balanceOf(owner.address)).to.equal(initialOwnerBalance);
    });

    it("Should approve tokens for delegated transfer", async function () {
        await token.approve(addr1.address, 100);
        expect(await token.allowance(owner.address, addr1.address)).to.equal(100);
    });

    it("Should transfer tokens via transferFrom", async function () {
        await token.approve(addr1.address, 100);
        await token.connect(addr1).transferFrom(owner.address, addr2.address, 100);

        expect(await token.balanceOf(addr2.address)).to.equal(100);
        expect(await token.allowance(owner.address, addr1.address)).to.equal(0);
    });

    it("Should mint new tokens only by the owner", async function () {
        await token.mint(addr1.address, 100);
        const addr1Balance = await token.balanceOf(addr1.address);
        expect(addr1Balance).to.equal(100);

        await expect(token.connect(addr1).mint(addr2.address, 100))
            .to.be.revertedWith("Not the contract owner");
    });

    it("Should burn tokens correctly", async function () {
        const initialSupply = await token.totalSupply();

        await token.burn(50000);
        const finalSupply = await token.totalSupply();

        expect(finalSupply).to.equal(initialSupply.sub(50000));
    });
});
