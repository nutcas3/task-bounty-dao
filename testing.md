Based on your taskbounty module implementation, here are all the REST API query commands you can use to interact with your task bounty system:

Task Query Commands (REST API)
bash
CopyInsert
# List all tasks
curl http://localhost:1317/taskbounty/tasks

# Show a specific task by ID
curl http://localhost:1317/taskbounty/tasks/1

# List tasks by status (OPEN, CLAIMED, PROOF_SUBMITTED, COMPLETED)
curl http://localhost:1317/taskbounty/tasks/status/OPEN
curl http://localhost:1317/taskbounty/tasks/status/CLAIMED
curl http://localhost:1317/taskbounty/tasks/status/PROOF_SUBMITTED
curl http://localhost:1317/taskbounty/tasks/status/COMPLETED

# List tasks by creator address
curl http://localhost:1317/taskbounty/tasks/creator/cosmos1...

# List tasks by claimant address
curl http://localhost:1317/taskbounty/tasks/claimant/cosmos1...
With Pagination
bash
CopyInsert
# List all tasks with pagination
curl "http://localhost:1317/taskbounty/tasks?page=1&limit=10"

# List tasks by status with pagination
curl "http://localhost:1317/taskbounty/tasks/status/OPEN?page=1&limit=10"
With JSON Output Formatting
bash
CopyInsert
# Pretty-print JSON output
curl http://localhost:1317/taskbounty/tasks | jq
Using the Know Freedom Testnet
If you're connecting to the Know Freedom testnet (as mentioned in your project requirements):

bash
CopyInsert
# Set the node endpoint
export NODE="https://node0.testnet.knowfreedom.io:26657"

# List all tasks
curl "https://node0.testnet.knowfreedom.io:1317/taskbounty/tasks"
These commands will allow you to query all aspects of your tokenized task bounty system through the REST API. The server is already running on port 1317, so you can start using these commands right away.

Feedback submitted
