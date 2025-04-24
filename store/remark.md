npm install -g solc
solcjs --bin Store.sol
solcjs --abi Store.sol

abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=store.go