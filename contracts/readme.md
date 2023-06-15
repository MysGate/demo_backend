## 1 
## npm install -g solc@xxx
npm install solc

## 2
solcjs --abi .\CrossController_flattened.sol

## 3
abigen --abi=.\CrossController_flattened_sol_CrossController.abi --pkg=cross --out=CrossController.go

## 4
solcjs --abi .\Bridge_flattened.sol

## 5
abigen --abi=.\Bridge_flattened_sol_Bridge.abi --pkg=bridge --out=Bridge.go

## 6 rename pkg to contracts