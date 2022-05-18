PROJECT REFERENCE REPORT
-------------------------

This file is aimed at serving the Project Reference Report Requirement for Project 3 of CMSC 621 Advanced Operating System, Prof. Kalpakis, 2022 Spring. The report, which complements the code in this project, belongs to Vineeth Nallan Chakravarthula (XM91974).

Most of the code is completely original. The parts of the code that are reused from a different source are as follows:

The hash function:
    Link: https://www.csee.umbc.edu/~kalpakis/Courses/621-sp22/project/GoTokens-1.pdf
    Code used in file: servercode/servercode.go
    Line numbers in destination: 19-24
    Lines of code used: 6
    Reason: The specific hash function to be used was supplied in the question statement and was used as it is in order to ensure that the hash calculations are correct.

Lines of code resued: 6
Total lines of code: 552 + 173 = 725

Reused code ratio: 6/725 = 0.008


IMPLEMENTATION
-------------------------

The bulk of the project code is in the `servercode/servercode.go` file. The execution begins at the `clientcode/clientcode.go` file. 

**Create RPC Call**
Upon encountering a Create flag for a specified token in the command line, the client program reads the `yaml_final.yml` file and fetches the corresponding token's data. Once the data is fetched, the writer of the token is spawned by the client, which then becomes the recipient of a subsequent Create RPC call.

The writer, upon receiving the Create RPC Call, creates the token if it does not already exist. It then takes the responsibility of spawning the reader servers and issuing Create RPC calls to them so that the token may be replicated. The RPC calls are modified to be flexible enough to determine where the source of the RPC call is, which in turn allows the servers to determine whether to fetch data from the YAML file or not.


**Write RPC Call**
This works more or less on the same ideas behind the Create RPC call. Upon encountering the Write flag, the client program reads the `yaml_final.yml` file and fetches the corresponding token's data. Once the data is fetched, the writer of the token is sent the Write RPC call.

The writer, upon receiving the Write RPC call compares the timestamp on the token to the incoming call's timestamp. If the incoming call has a more recent timestamp, then the write operation is performed. The token's state information is updated, including the new timestamp. Now, it takes the responsibility of informing the readers to update the token's state information as well by sending the relevant readers the Write RPC call. As with Create, the Write RPC call is flexible enough to determine where the source of the call is, thereby allowing a distinction between the writer and the reader in the same function.


**Read RPC Call**
Here, the client program reads the `yaml_final.yml` file and fetches the corresponding token's data. This time, it focuses on getting the list of readers and chooses a reader randomly to send the Read RPC call to. Once the corresponding reader receives the RPC call, it checks whether the token exists or not. After this check is passed, the Read-Impose Write-Majority comes into play, where all the nodes with the token are contacted by dispatching RPC calls as goroutines. This is where the **RIWMTest RPC Call** comes into play. 

When a server receives this call, it checks for the token for which the information is requested, and sends back the timestamp and the final value. The reader that invokes this RPC call then begins to collect the results of the goroutines. Through the process of collection, the timestamps are compared and the most recent value is calculated. However, the read waits only until a majority of the calls are executed and returned, and returns the updated information to the client while also writing back to the other nodes.


**Drop RPC Call**
Like with the Create and Write RPC calls, the client program fetches the token's information from the `yaml_final.yml` file and gets the writer. The writer then receives a Drop RPC call that has to be executed. If the Drop RPC call is more recent than the last performed operation on the token, then the token is dropped from the server. 

The writer then dispatches Drop RPC calls to all the readers in order to ensure that the replicated token is deleted. It does this by dispatching the Drop RPC calls as goroutines, thereby ensuring that the deletion is not halted, and that the client would not have to wait for too long.


**Fail-Silent Behaviour**
To satisfy the requirement of emulating fail-silent behaviour, the server contains hardcoded values to check for the conditions for a specific token in a specific server. This is done so that pinpointing the behaviour will be simple. In this project, the fail-silent behaviour is emulated for the token with the ID: 1020 in the server running on the port: 65000. The server running on this port will fail to respond to queries on the token 1020 after a time duration of 10 seconds since the token's creation.


OTHER NOTES
-------------------------

The code contains a plethora of other functions that act as helper functions. The following is a comprehensive list of functions:

* `Hash`: calculate the hash value in a given range.
* `yaml_data_retriever`: retrive the token data for a specific token from the YAML file.
* `is_exists`: check if a token already exists or not.
* `print_all_tokens`: prints token IDs of all tokens except a specified token.
* `print_current_token`: prints complete information of a specified token.
* `fail_silent_check`: function to determine whether the fail-silent emulation conditions are satisfied.
* `get_port_list`: get the list of reader ports.
* `get_finalvals`: dispatch a read-impose write-majority RPC call and collect the final values and corresponding timestamps.