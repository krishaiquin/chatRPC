## Legend
    ✅ Done 
    🚧 In Progress
    🔁 Deferred
    🛑 Stop, no longer needed

## Weekly Tasks/Progress

### Week 1
    ✅ Set up simple transport layer - works only on system with only one function 
    ✅ Echo messages using RPC - created message service

### Week 2
    ✅ Expand transport layer to work on multiple functions
    ✅ Figure out how to create set of client nodes - added nodeset service
    ✅ Add daemon for message service
    ✅ Set up a db service to save the endpoints for different services 

### Week 3
    ✅ Fix all bugs and connected chat to all services (message, nodeset and db)
    ✅ Add quit command to chat
    🔁 Figure out how to broadcast messages 
    ✅ Re-design Nodeset - figure out how nodes would use it and how it should work
           
### Week 4
    🚧 Rewrite Nodeset Architecture 
    🛑 Rewrite transport layer to allow nodes to listen for "Update Nodeset" requests for Nodeset service

### Week 5
    ✅ Rewrite Nodeset Architecture
    ✅ Test Nodeset Architecture - pt. 1
    ✅ Figure out what's wrong when nodeset sends a request to client nodes
### Week 6
    ✅ Implement the potential solution describe in Notes / Design Insights Week 5 section
    ✅ Learn more about goroutines, select statements, sync.WaitGroup, channels
    ✅ Node discovery now works!
### Week 7
    🔁 Refactor Code
    🔁 Add better error handling
    ✅ Learn and add context (graceful shutdown on keyboard interrupt)
    ✅ Add chat messaging
    🚧 Fix chat messaging  
### Week 8 
    ✅ Fix chat messaging -- MVP now done
    ✅ Work on improvements like add user names     
    🚧 Add TUI
    🔁 Add better error handling
    🔁 Refactor Code

### 💭 Notes / Design Insights
    - Week 4:
        - Rewriting Node Discovery Implementation Ideas:
            -  create a Node struct. This contains NodeId and its address
            -  create a Cluster struct. This contains "this Node" and other Nodes in the cluster (aka nodeset). This will help to know which Node sent which message.
            -  make nodeset field(Cluster struct) from []string into []Node.    
                Note: Drew a diagram and it doesn't look like there problem with nodes joining the cluster BUT I'm hitting a roadblock with figuring out how to let other nodes in the cluster know aboout the newly added nodes.
                    - This means that nodes will have to listen for "Update Nodeset" requests from Nodeset service. Need to rewrite transport layer to be able to allow this.
            - Node and Cluster creation is done in Nodeset service. A newly-joined node will send a request to Nodeset to add it  to the cluster. Nodeset service will return a Cluster
        - Problem: Node and Cluster structs are going to be pointers. When it gets to client-side, the addresses don't mean anything.
            - IDEAS for solution:
                - Create Node/Cluster locally in nodes
                - NodeId creation and cluster management is done in nodeset
                - Nodes requests Nodeset for NodeId creation
                - Nodes requests Nodeset to Add node to the cluster. Return nil (for now)
                - Nodes requests Nodeset for the list of nodes in the cluster. Return []Nodes
    - Week 5
        - Current Node Discovery Implementation:
            - Client nodes have their own local copy of cluster (list of nodes)
            - Nodeset service has its own copy of cluster but this cluster is the source of truth
            - When a node enters a cluster - this node sends a request to nodeset service to add it to the cluster
            - Nodeset, then, sends a request to all the nodes in the cluster to update their copy of cluster.
        - Problem: Since client nodes both sends and receive requests, they don't know if the received packet is a request or a response
        - Potential solution: Add request# and response# - if request# == response#, send that response to the appropriate request.
    - Week 6
        - The suggested solution mentioned in Week 5 worked. 
        - More in-depth explanation:
            - Each packet from a node has a sequence number and a flag whether it is a request or a result.
            - The receiver node then unpacks the packet and process accordingly
            - If the packet is a request, the transport class will dispatch it to the appropriate function to process it
            - If the packet is a result, it will get sent to the application and remove the request from the list of pending requests with the associated seq#
        - In addition to that, I've put Listen() in goroutine so that services and nodes are able to send requests and receive results at the same time. 
        - In doing this, I am learning how to manage goroutines effectively, including working with channels, sync package, select statements.
        - I think this is a good point to clean up the code - refactor, add more error handling and finally introducing context in the project --> main goal for next week
    - Week 7
        - Spent most of week 7 remembering and understanding the code base, coming back from long hiatus
        - Learnt context package -- specifically, withCancel()
        - Started on chat messaging
            - sending and receiving message between nodes are partially working -- sender crashes after sending message
    - Week 8
        - Found the bug on chat messaging -- clientStub of Send() was expecting a return value
        - Upon reviewing the code, it looks like the error handling and the code itself are sufficient for what I intend for this project to be. For that reason, I've decided to defer error handling and code refactoring indefinitely. 
        - Improved nodeset and nodesetManager. Before, when a node enters the cluster, nodeset will send the full cluster list to each node including the new node. Now, when a node enters the cluster, the nodeset gives the new node its nodeId and the cluster list. Nodeset then sends a request to the rest of cluster to add that new Node to their list.
        - Used tview to create TUI of the chat. Added few elements, no functionalities yet
<!-- #### Learning Goals
    - Implement RPC from scratch. Learn more about: 
        - client stub
        - server stub
        - RPC mechanism
    - Broadcast message to client node cluster. Learn more about:
        - transport protocol
    - Add context to the application. Learn more about:
        - Contexts
    - Add vector clocks. Learn more about:
        - causal events -->