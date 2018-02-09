# Network Tests


|            |  TCP  |  TLS  |
|-----------:|:-----:|:-----:|
|   Client   | S + P | TODO  |
|Echo Server |   Y   | TODO  |
|Client + Tor|  N/A  | TODO  |


### Echo Server
Echo server using wait groups and goroutines

### Client
##### Parallel Client
(-) Order and atomicity not maintained

(+) Fast

##### Serial Client
(+) Order and Atomicity maintained

(-) Slow

##### Client over Tor
TODO: Determine how to package cells to appropriately 
pass through a Tor Router. (Host my own for testing)  


## TODO:
1. TLS Client Serial(S) and Parallel(P)

2. TLS Server sync waitgroups & goroutines

3. Host Tor router and find example client

4. Pack cells to contact TLS / TCP Echo Server through Tor
