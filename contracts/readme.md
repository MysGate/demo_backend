## 1
npm install solcjs

## 2
solcjs --abi .\CrossController_flattened.sol

## 3
abigen --abi=.\CrossController_flattened_sol_CrossController.abi --pkg=cross --out=CrossController.go