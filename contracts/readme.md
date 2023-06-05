## 1 
## npm install -g solc@xxx
npm install solc

## 2
solcjs --abi .\CrossController_flattened.sol

## 3
abigen --abi=.\CrossController_flattened_sol_CrossController.abi --pkg=cross --out=CrossController.go