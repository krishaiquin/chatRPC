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
    [ ] Implement the potential solution describe in Notes / Design Insights Week 5 section
<!-- [ ] Refactor RPC -->
<!-- [ ] Add context -->
<!-- [ ] Test program's functionality    -->
<!-- [ ] Add vector clocks for causality -->
<!-- [ ] Test vector clocks (simulate real-life node connection)  -->

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
            - When a node enters a cluster - this node sends a request to nodeset service to it to the cluster
            - Nodeset, then, sends a request to all the nodes in the cluster to update their copy of cluster.
        - Problem: Since client nodes both sends and receive requests, they don't know if the received packet is a request or a response
        - Potential solution: Add request# and response# - if request# == response#, send that response to the appropriate request.
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