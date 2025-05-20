### Week 1
    ✅ Set up simple transport layer - works only on system with only one function 
    ✅ Echo messages using RPC - created message service
### Week 2
    ✅ Expand transport layer to work on multiple functions
    ✅ Figure out how to create set of client nodes - added nodeset service
    ✅ Add daemon for message service
    ✅ Set up a db service to save the endpoints for different services 
### Week 3
    ✅ Fixed all bugs and connected chat to all services (message, nodeset and db)
    [In Progress] Figure out how to broadcast messages
    [ ] Rewrite transport layer to allow multicast/broadcast messaging
    [In Progress] Rewrite client nodeset to access them easily
        TODO:
            -  create a Node struct
            -  make Nodeset from []string into []Node
    ✅ Add quit command to chat
    [ ] Refactor RPC
    [ ] Add context
    [ ] Test program's functionality   
    <!-- [ ] Add vector clocks for causality
    [ ] Test vector clocks (simulate real-life node connection)  -->


#### Learning Goals
    - Implement RPC from scratch. Learn more about: 
        - client stub
        - server stub
        - RPC mechanism
    - Broadcast message to client node cluster. Learn more about:
        - transport protocol
    - Add context to the application. Learn more about:
        - Contexts