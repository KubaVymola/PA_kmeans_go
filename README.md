# Usage

Use either:

- $ go run . -- \<number of points\> \<k\> \<Number of parallel threads\>

or

- 1. $ go build k_means.go vector.go
  2. $ ./k_means -- \<number of points\> \<k\> \<Number of parallel threads\>

# Remarks

- There is a data race condition when setting the variable "change" in the function "calculateNewOwners". I suspect, that this should not cause a problem, since the variable is treated as write-only during the extent of the data race and the value can only be set from false to true, never in the other way.

- In any case, there are instructions on how to get rid of the data race condition in the code. However, it very much slows the parallelism.

- The program takes about 750 ms to converge for 10,000 points and k = 20. When running on 8-core CPU with the data race present, it can take as little as 250 ms to converge and I have not observed any misbehaviour of the program.

- Example output gif is present in the ./output directory

# Author

Jakub Výmola (VYM0038), VŠB-FEI, 2021-2022
Course: Parallel aglorithms